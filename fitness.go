package main

import "fmt"

type BotRevenue struct {
	BotNumber        int
	Revenue          float64
	TotalBuysCount   int
	SuccessBuysCount int
	FailedBuysCount  int
}

func Fitness(botConfig BotConfig, botNumber int, botRevenue chan BotRevenue, fitnessDatasets *[]Dataset) {
	totalRevenue := 0.0
	totalBuysCount := 0
	totalSuccessBuysCount := 0

	for _, dataset := range *fitnessDatasets {
		datasetRevenue, buyCount, successBuysCount := doBuysAndSells(dataset, botConfig)
		totalRevenue += datasetRevenue
		totalBuysCount += buyCount
		totalSuccessBuysCount += successBuysCount
	}

	botRevenue <- BotRevenue{
		BotNumber:        botNumber,
		Revenue:          totalRevenue,
		TotalBuysCount:   totalBuysCount,
		SuccessBuysCount: totalSuccessBuysCount,
		FailedBuysCount:  totalBuysCount - totalSuccessBuysCount,
	}
	//return totalRevenue
}

func doBuysAndSells(dataset Dataset, botConfig BotConfig) (float64, int, int) {
	dataSource := NewDataSource()
	coinBotFactory := NewCoinBotFactory(&dataSource)
	exchangeManager := NewExchangeManager(botConfig)
	candleMarketStat := NewCandleMarketStat(botConfig, &dataSource)
	positiveApproach := NewPositiveApproach(botConfig, &exchangeManager, &candleMarketStat)
	trailing := NewTrailingSymbol(botConfig.TrailingLowPercentage)

	for candleNum, candle := range dataset.AltCoinCandles {
		btcDataset := *dataset.BtcCandles
		hasSecondPercentageBuySignal := false

		candleHandler(
			candle,
			btcDataset[candleNum],
			botConfig,
			dataSource,
			coinBotFactory,
			exchangeManager,
			candleMarketStat,
			positiveApproach,
			&trailing,
			&hasSecondPercentageBuySignal,
		)
	}

	rev := exchangeManager.GetTotalRevenue()
	buyCount := exchangeManager.GetBuysCount()
	success := exchangeManager.GetSuccessBuysCount()
	commission := float64(buyCount) * COMMISSION
	exchangeManager.Close()

	datasetRevenue := rev - commission
	failed := buyCount - success

	fmt.Println(fmt.Sprintf("%s: DatasetRevenue: %f, TotalBuys: %d, Success: %d, Failed: %d", dataset.AltCoinName, datasetRevenue, buyCount, success, failed))

	return datasetRevenue, buyCount, success
}

func candleHandler(
	candle Candle,
	btcCandle Candle,
	botConfig BotConfig,
	dataSource DataSource,
	coinBotFactory CoinBotFactory,
	exchangeManager ExchangeManager,
	candleMarketStat CandleMarketStat,
	positiveApproach PositiveApproach,
	trailing *Trailing,
	hasSecondPercentageBuySignal *bool,
) {
	dataSource.AddCandleFor(candle.Symbol, candle)
	dataSource.AddCandleFor(btcCandle.Symbol, btcCandle)
	bot := coinBotFactory.FactoryCoinBot(candle.Symbol, botConfig)

	/* !Important to update trailing each candle update */
	isUpdated := trailing.Update(candle)
	if isUpdated {
		fmt.Println(fmt.Sprintf("TrailingUpdate: COIN: %s, EXCHANGE_RATE: %f, TIME: %s", candle.Symbol, candle.ClosePrice, candle.CloseTime))
	}

	updateBuys(candle, exchangeManager, candleMarketStat, trailing, hasSecondPercentageBuySignal)
	//positiveApproach.UpdateBuys(candle)

	if !*hasSecondPercentageBuySignal && candleMarketStat.HasCoinGoodDoubleTrend(candle) /*&& candleMarketStat.HasBtcBuyPercentage() */ && bot.HasBuySignal() {

		//if positiveApproach.HasSignal(candle) {
		if SIMULTANEOUS_BUYS_COUNT > exchangeManager.CountUnsoldBuys(candle.Symbol) {
			// Do buy

			fmt.Println(fmt.Sprintf("COIN: %s, BUY: %s, EXCHANGE_RATE: %f, Volume: %f", candle.Symbol, candle.CloseTime, candle.GetCurrentPrice(), candle.Volume))
			exchangeManager.Buy(candle.Symbol, candle.GetCurrentPrice(), candle.CloseTime)
			*hasSecondPercentageBuySignal = true
			trailing.Start(candle)
			//bot.ResetHasReached()
		}
		//}
	}
}

func updateBuys(
	candle Candle,
	exchangeManager ExchangeManager,
	candleMarketStat CandleMarketStat,
	trailing *Trailing,
	hasSecondPercentageBuySignal *bool,
) {
	if trailing.CanSellByStop(candle) {
		trailingUnsoldBuys := exchangeManager.UpdateAllExitSymbols(candle.Symbol, candle.ClosePrice, candle.CloseTime)
		if len(trailingUnsoldBuys) > 0 {
			*hasSecondPercentageBuySignal = false
		}
	}

	//unsoldBuys := exchangeManager.UpdateBuys(candle.Symbol, candle.ClosePrice, candle.CloseTime)
	//if len(unsoldBuys) > 0 {
	//	*hasSecondPercentageBuySignal = false
	//}
	//
	//if candleMarketStat.HasBtcSellPercentage() {
	//	btcUnsoldBuys := exchangeManager.UpdateAllExitSymbols(candle.Symbol, candle.ClosePrice, candle.CloseTime)
	//	if len(btcUnsoldBuys) > 0 {
	//		*hasSecondPercentageBuySignal = false
	//	}
	//}
}

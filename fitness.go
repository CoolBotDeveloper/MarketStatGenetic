package main

import (
	"fmt"
	"math"
)

type BotRevenue struct {
	BotNumber        int
	Revenue          float64
	TotalBuysCount   int
	SuccessBuysCount int
	FailedBuysCount  int

	PlusRevenue  float64
	MinusRevenue float64
}

func Fitness(botConfig BotConfig, botNumber int, botRevenue chan BotRevenue, fitnessDatasets *[]Dataset) {
	totalRevenue := 0.0
	totalBuysCount := 0
	totalSuccessBuysCount := 0
	totalPlusRevenue := 0.0
	totalMinusRevenue := 0.0

	for _, dataset := range *fitnessDatasets {
		datasetRevenue, buyCount, successBuysCount, plusRevenue, minusRevenue := doBuysAndSells(dataset, botConfig)
		totalRevenue += datasetRevenue
		totalBuysCount += buyCount
		totalSuccessBuysCount += successBuysCount
		totalPlusRevenue += plusRevenue
		totalMinusRevenue += minusRevenue
	}

	botRevenue <- BotRevenue{
		BotNumber:        botNumber,
		Revenue:          totalRevenue,
		TotalBuysCount:   totalBuysCount,
		SuccessBuysCount: totalSuccessBuysCount,
		FailedBuysCount:  totalBuysCount - totalSuccessBuysCount,
		PlusRevenue:      totalPlusRevenue,
		MinusRevenue:     totalMinusRevenue,
	}
	//return totalRevenue
}

func doBuysAndSells(dataset Dataset, botConfig BotConfig) (float64, int, int, float64, float64) {
	dataSource := NewDataSource()
	coinBotFactory := NewCoinBotFactory(&dataSource)
	exchangeManager := NewExchangeManager(botConfig)
	candleMarketStat := NewCandleMarketStat(botConfig, &dataSource)
	positiveApproach := NewPositiveApproach(botConfig, &exchangeManager, &candleMarketStat)
	trailing := NewTrailingSymbol(botConfig)

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

	failed := buyCount - success

	plusRevenue := 0.0
	minusRevenue := 0.0

	if buyCount > 0 {
		prevPlusRevenue := exchangeManager.GetPlusRevenue() - float64(success)*COMMISSION
		if prevPlusRevenue > 0.0 {
			plusRevenue = prevPlusRevenue
		} else {
			minusRevenue = prevPlusRevenue
		}

		minusRevenue = math.Abs(exchangeManager.GetMinusRevenue() - float64(failed)*COMMISSION - minusRevenue)
	}

	exchangeManager.Close()

	datasetRevenue := rev - commission
	//failed := buyCount - success

	fmt.Println(fmt.Sprintf("%s: DatasetRevenue: %f, TotalBuys: %d, Success: %d, Failed: %d", dataset.AltCoinName, datasetRevenue, buyCount, success, failed))

	return datasetRevenue, buyCount, success, plusRevenue, minusRevenue
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
	trailing.Update(candle)
	//if isUpdated {
	//	fmt.Println(fmt.Sprintf("TrailingUpdate: COIN: %s, EXCHANGE_RATE: %f, TIME: %s", candle.Symbol, candle.ClosePrice, candle.CloseTime))
	//}

	updateBuys(candle, exchangeManager, candleMarketStat, trailing, hasSecondPercentageBuySignal)
	//positiveApproach.UpdateBuys(candle)

	if !*hasSecondPercentageBuySignal && candleMarketStat.HasCoinGoodDoubleTrend(candle) && candleMarketStat.HasBtcBuyPercentage() && bot.HasBuySignal() {

		//if positiveApproach.HasSignal(candle) {
		if SIMULTANEOUS_BUYS_COUNT > exchangeManager.CountUnsoldBuys(candle.Symbol) && exchangeManager.CanBuyInGivenPeriod(candle.Symbol) {
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
		if trailingStopPrice, ok := trailing.GetStopPrice(candle); ok {
			trailingUnsoldBuys := exchangeManager.UpdateAllExitSymbols(candle.Symbol, trailingStopPrice, candle.CloseTime)
			if len(trailingUnsoldBuys) > 0 {
				*hasSecondPercentageBuySignal = false
			}
			trailing.Finish(candle)
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

package main

import (
	"fmt"
)

type BotRevenue struct {
	BotNumber int
	Revenue   float64
}

func Fitness(botConfig BotConfig, botNumber int, botRevenue chan BotRevenue, fitnessDatasets *[]Dataset) {
	totalRevenue := 0.0

	for _, dataset := range *fitnessDatasets {
		datasetRevenue := doBuysAndSells(dataset, botConfig)
		totalRevenue += datasetRevenue
		fmt.Println(fmt.Sprintf("%s: DatasetRevenue: %f", dataset.AltCoinName, datasetRevenue))
	}

	botRevenue <- BotRevenue{
		BotNumber: botNumber,
		Revenue:   totalRevenue,
	}
	//return totalRevenue
}

func doBuysAndSells(dataset Dataset, botConfig BotConfig) float64 {
	dataSource := NewDataSource()
	coinBotFactory := NewCoinBotFactory(&dataSource)
	exchangeManager := NewExchangeManager(botConfig)
	candleMarketStat := NewCandleMarketStat(botConfig, &dataSource)
	positiveApproach := NewPositiveApproach(botConfig, &exchangeManager, &candleMarketStat)

	for candleNum, candle := range dataset.AltCoinCandles {
		btcDataset := *dataset.BtcCandles

		candleHandler(
			candle,
			btcDataset[candleNum],
			botConfig,
			dataSource,
			coinBotFactory,
			exchangeManager,
			candleMarketStat,
			positiveApproach,
		)
	}

	rev := exchangeManager.GetTotalRevenue()
	commission := float64(exchangeManager.GetBuysCount()) * COMMISSION
	exchangeManager.Close()

	return rev - commission
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
) {
	dataSource.AddCandleFor(candle.Symbol, candle)
	dataSource.AddCandleFor(btcCandle.Symbol, btcCandle)
	bot := coinBotFactory.FactoryCoinBot(candle.Symbol, botConfig)

	updateBuys(candle, exchangeManager, candleMarketStat)
	positiveApproach.UpdateBuys(candle)

	if candleMarketStat.HasCoinGoodDoubleTrend(candle) &&
		candleMarketStat.HasBtcBuyPercentage() &&
		bot.HasBuySignal() {

		if positiveApproach.HasSignal(candle) {
			if 1 > exchangeManager.CountUnsoldBuys(candle.Symbol) {
				// Do buy
				exchangeManager.Buy(candle.Symbol, candle.ClosePrice)
			}
		}
	}
}

func updateBuys(candle Candle, exchangeManager ExchangeManager, candleMarketStat CandleMarketStat) {
	exchangeManager.UpdateBuys(candle.Symbol, candle.ClosePrice, candle.CloseTime)

	if candleMarketStat.HasBtcSellPercentage() {
		exchangeManager.UpdateAllExitSymbols(candle.Symbol, candle.ClosePrice, candle.CloseTime)
	}
}

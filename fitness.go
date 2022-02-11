package main

import "fmt"

var dataSource DataSource
var coinBotFactory CoinBotFactory
var exchangeManager ExchangeManager
var candleMarketStat CandleMarketStat

func Fitness(botConfig BotConfig) float64 {
	totalRevenue := 0.0
	datasets := ImportDatasets()

	for _, dataset := range datasets {
		fmt.Println(dataset.AltCoinName)
		totalRevenue += doBuysAndSells(dataset, botConfig)
	}

	return totalRevenue
}

func doBuysAndSells(dataset Dataset, botConfig BotConfig) float64 {
	dataSource = NewDataSource()
	coinBotFactory = NewCoinBotFactory(&dataSource)
	exchangeManager = NewExchangeManager(botConfig)
	candleMarketStat = NewCandleMarketStat(botConfig, &dataSource)

	for candleNum, candle := range dataset.AltCoinCandles {
		candleHandler(candle, dataset.BtcCandles[candleNum], botConfig)
	}

	rev := exchangeManager.GetTotalRevenue()
	commission := float64(exchangeManager.GetBuysCount()) * COMMISSION
	exchangeManager.Close()

	return rev - commission
}

func candleHandler(candle Candle, btcCandle Candle, botConfig BotConfig) {
	dataSource.AddCandleFor(candle.Symbol, candle)
	dataSource.AddCandleFor(btcCandle.Symbol, btcCandle)

	bot := coinBotFactory.FactoryCoinBot(candle.Symbol, botConfig)

	hasBuySignal := bot.HasBuySignal()
	if candleMarketStat.HasCoinGoodDoubleTrend(candle) &&
		candleMarketStat.HasBtcBuyPercentage() &&
		hasBuySignal &&
		1 > exchangeManager.CountUnsoldBuys(candle.Symbol) {

		// Do buy
		exchangeManager.Buy(candle.Symbol, candle.ClosePrice)
	}
}

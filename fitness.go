package main

import "fmt"

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
	dataSource := NewDataSource()
	coinBotFactory := NewCoinBotFactory(&dataSource)
	exchangeManager := NewExchangeManager(botConfig)
	candleMarketStat := NewCandleMarketStat(botConfig, &dataSource)

	for candleNum, candle := range dataset.AltCoinCandles {
		candleHandler(
			candle,
			dataset.BtcCandles[candleNum],
			botConfig,
			dataSource,
			coinBotFactory,
			exchangeManager,
			candleMarketStat,
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
) {
	dataSource.AddCandleFor(candle.Symbol, candle)
	dataSource.AddCandleFor(btcCandle.Symbol, btcCandle)
	bot := coinBotFactory.FactoryCoinBot(candle.Symbol, botConfig)

	updateBuys(candle, exchangeManager, candleMarketStat)

	if candleMarketStat.HasCoinGoodDoubleTrend(candle) &&
		candleMarketStat.HasBtcBuyPercentage() &&
		1 > exchangeManager.CountUnsoldBuys(candle.Symbol) &&
		bot.HasBuySignal() {

		// Do buy
		exchangeManager.Buy(candle.Symbol, candle.ClosePrice)
	}
}

func updateBuys(candle Candle, exchangeManager ExchangeManager, candleMarketStat CandleMarketStat) {
	exchangeManager.UpdateBuys(candle.Symbol, candle.ClosePrice)

	if candleMarketStat.HasBtcSellPercentage() {
		exchangeManager.UpdateAllExitSymbols(candle.Symbol, candle.ClosePrice)
	}
}

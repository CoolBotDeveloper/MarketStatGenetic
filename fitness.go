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
}

func doBuysAndSells(dataset Dataset, botConfig BotConfig) (float64, int, int, float64, float64) {
	dataSource := NewDataSource()
	coinBotFactory := NewCoinBotFactory(&dataSource)
	exchangeManager := NewExchangeManager(botConfig)

	for _, candle := range dataset.AltCoinCandles {
		candleHandler(
			candle,
			botConfig,
			dataSource,
			coinBotFactory,
			exchangeManager,
		)
	}

	rev := exchangeManager.GetTotalRevenue()
	buyCount := exchangeManager.GetBuysCount()
	successCount := exchangeManager.GetSuccessBuysCount()
	commission := float64(buyCount) * COMMISSION

	failedCount := buyCount - successCount

	plusRevenue := 0.0
	minusRevenue := 0.0

	if buyCount > 0 {
		prevPlusRevenue := exchangeManager.GetPlusRevenue() - float64(successCount)*COMMISSION
		if prevPlusRevenue >= 0.0 {
			plusRevenue = prevPlusRevenue
		} else {
			minusRevenue = prevPlusRevenue
		}

		minusRevenue = math.Abs(exchangeManager.GetMinusRevenue()) +
			math.Abs(float64(failedCount)*COMMISSION) +
			math.Abs(minusRevenue)
	}

	exchangeManager.Close()
	datasetRevenue := rev - commission

	fmt.Println(fmt.Sprintf("%s: DatasetRevenue: %f, TotalBuys: %d, Success: %d, Failed: %d", dataset.AltCoinName, datasetRevenue, buyCount, successCount, failedCount))

	return datasetRevenue, buyCount, successCount, plusRevenue, minusRevenue
}

func candleHandler(
	candle Candle,
	botConfig BotConfig,
	dataSource DataSource,
	coinBotFactory CoinBotFactory,
	exchangeManager ExchangeManager,
) {
	dataSource.AddCandleFor(candle.Symbol, candle)
	bot := coinBotFactory.FactoryCoinBot(candle.Symbol, botConfig)

	updateBuys(candle, exchangeManager)

	if //candleMarketStat.HasCoinGoodDoubleTrend(candle) &&
	//candleMarketStat.HasAltCoinMarketPercentage(candle) &&
	//candleMarketStat.HasCoinGoodSingleTrend(candle) &&
	//candleMarketStat.HasNotCoinMaxPercentage(candle) &&
	bot.HasBuySignal() {

		if SIMULTANEOUS_BUYS_COUNT > exchangeManager.CountUnsoldBuys(candle.Symbol) {
			// Do buy
			fmt.Println(fmt.Sprintf("COIN: %s, BUY: %s, EXCHANGE_RATE: %f, Volume: %f", candle.Symbol, candle.CloseTime, candle.GetCurrentPrice(), candle.Volume))
			exchangeManager.Buy(candle.Symbol, candle.GetCurrentPrice(), candle.CloseTime)
		}
	}
}

func updateBuys(
	candle Candle,
	exchangeManager ExchangeManager,
) {
	exchangeManager.UpdateNormalBuys(candle.Symbol, candle.ClosePrice, candle.CloseTime)
}

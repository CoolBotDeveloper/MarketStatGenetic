package main

type CandleMarketStat struct {
	config     BotConfig
	dataSource *DataSource
}

func NewCandleMarketStat(config BotConfig, dataSource *DataSource) CandleMarketStat {
	return CandleMarketStat{config: config, dataSource: dataSource}
}

func (marketStat *CandleMarketStat) HasBtcBuyPercentage() bool {
	if lastCandle, ok := marketStat.dataSource.GetLastCandleFor(BITCOIN_SYMBOL); ok {
		percentage, hasPercentage := marketStat.GetSymbolPercentageForPeriod(lastCandle, BITCOIN_SYMBOL, marketStat.config.BtcMinBuyPeriodMinutes)
		if hasPercentage {
			return percentage >= marketStat.config.BtcMinBuyPercentage
		}
	}

	return false
}

func (marketStat *CandleMarketStat) HasBtcSellPercentage() bool {
	if lastCandle, ok := marketStat.dataSource.GetLastCandleFor(BITCOIN_SYMBOL); ok {
		percentage, hasPercentage := marketStat.GetSymbolPercentageForPeriod(lastCandle, BITCOIN_SYMBOL, marketStat.config.BtcSellPeriodMinutes)
		if hasPercentage {
			return percentage <= marketStat.config.BtcSellPercentage
		}
	}

	return false
}

func (marketStat *CandleMarketStat) GetSymbolPercentageForPeriod(candle Candle, symbol string, periodMinutes int) (float64, bool) {
	candles := marketStat.dataSource.GetCandlesFor(symbol)
	count := len(candles)

	if count <= (periodMinutes + 1) {
		return 0.0, false
	}

	closePriceCandles := GetClosePrice(candles, periodMinutes)
	growth := marketStat.calcGrowth(closePriceCandles[0], candle.GetCurrentPrice())

	return growth, true
}

func (marketStat *CandleMarketStat) HasCoinGoodSingleTrend(candle Candle) bool {
	secondGrowthPercentage, hasSecondPercentage := marketStat.GetSymbolPercentageForPeriod(
		candle,
		candle.Symbol,
		marketStat.config.AltCoinMinBuySecondPeriodMinutes,
	)

	if hasSecondPercentage {
		return secondGrowthPercentage >= marketStat.config.AltCoinMinBuySecondPercentage
	}

	return false
}

func (marketStat *CandleMarketStat) HasCoinGoodDoubleTrend(candle Candle) bool {
	trendPercentage, hasTrendPercentage := marketStat.GetSymbolPercentageForPeriod(candle, candle.Symbol, marketStat.config.AltCoinMinBuyFirstPeriodMinutes)
	directionPercentage, hasDirectionPercentage := marketStat.GetSymbolPercentageForPeriod(candle, candle.Symbol, marketStat.config.AltCoinMinBuySecondPeriodMinutes)

	if hasTrendPercentage && hasDirectionPercentage {
		return trendPercentage >= marketStat.config.AltCoinMinBuyFirstPercentage &&
			directionPercentage >= marketStat.config.AltCoinMinBuySecondPercentage
	}

	return false
}

func (marketStat *CandleMarketStat) calcGrowth(startPrice, endPrice float64) float64 {
	if startPrice == 0 || endPrice == 0 {
		return 0.0
	}

	return ((endPrice * 100) / startPrice) - 100
}

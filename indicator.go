package main

import (
	"github.com/markcheno/go-talib"
	taindic "github.com/xyths/go-indicators"
)

type BuyTechnicalIndicator interface {
	HasBuySignal(candles []Candle) bool
}

// Super trend indicator
type SuperTrendIndicator struct {
	config BotConfig
}

func NewSuperTrendIndicator(config BotConfig) SuperTrendIndicator {
	return SuperTrendIndicator{config: config}
}

func (indicator *SuperTrendIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	properCount := indicator.config.AltCoinSuperTrendCandles + 250
	if count < properCount {
		return false
	}

	trendCandles := properCount
	_, trend := taindic.SuperTrend(
		(*indicator).config.AltCoinSuperMultiplier,
		(*indicator).config.AltCoinSuperTrendCandles,
		GetHighPrice(candles, trendCandles),
		GetLowPrice(candles, trendCandles),
		GetClosePrice(candles, trendCandles),
	)
	lastTrendIdx := len(trend) - 1

	return trend[lastTrendIdx] && indicator.hadRedTrendBefore(trend)
}

func (indicator *SuperTrendIndicator) hadRedTrendBefore(trend []bool) bool {
	if len(trend) < 2 {
		return false
	}

	prevIdx := len(trend) - 2
	return trend[prevIdx] == false
}

// Second SuperTrend indicator
//type SuperTrendIndicatorSecond struct {
//	config BotConfig
//}
//
//func NewSuperTrendIndicatorSecond(config BotConfig) SuperTrendIndicatorSecond {
//	return SuperTrendIndicatorSecond{config: config}
//}
//
//func (indicator *SuperTrendIndicatorSecond) HasBuySignal(candles []Candle) bool {
//	count := len(candles)
//	properCount := indicator.config.SuperTrendCandlesSecond + indicator.config.SuperTrendProperCount
//	if count < properCount {
//		return false
//	}
//
//	trendCandles := properCount
//	_, trend := taindic.SuperTrend(
//		(*indicator).config.SuperTrendMultiplierSecond,
//		(*indicator).config.SuperTrendCandlesSecond,
//		GetHighPrice(candles, trendCandles),
//		GetLowPrice(candles, trendCandles),
//		GetClosePrice(candles, trendCandles),
//	)
//	lastTrendIdx := len(trend) - 1
//
//	return trend[lastTrendIdx]
//}

// Bitcoin Super trend indicator
type BitcoinSuperTrendIndicator struct {
	config BotConfig
}

func NewBitcoinSuperTrendIndicator(config BotConfig) BitcoinSuperTrendIndicator {
	return BitcoinSuperTrendIndicator{config: config}
}

func (indicator *BitcoinSuperTrendIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	properCount := indicator.config.BtcSuperTrendCandles + 250
	if count < properCount {
		return false
	}

	trendCandles := properCount
	_, trend := taindic.SuperTrend(
		(*indicator).config.BtcSuperTrendMultiplier,
		(*indicator).config.BtcSuperTrendCandles,
		GetHighPrice(candles, trendCandles),
		GetLowPrice(candles, trendCandles),
		GetClosePrice(candles, trendCandles),
	)
	lastTrendIdx := len(trend) - 1

	return trend[lastTrendIdx]
}

// Average volume indicator
type AverageVolumeIndicator struct {
	config BotConfig
}

func NewAverageVolumeIndicator(config BotConfig) AverageVolumeIndicator {
	return AverageVolumeIndicator{config: config}
}

func (indicator *AverageVolumeIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	if count < indicator.config.AverageVolumeCandles {
		return false
	}

	volumes := GetVolumes(candles, indicator.config.AverageVolumeCandles)
	avgVolume := GetAvg(volumes)

	return avgVolume >= indicator.config.AverageVolumeMinimal
}

// Adx indicator
type AdxIndicator struct {
	config BotConfig
}

func NewAdxIndicator(config BotConfig) AdxIndicator {
	return AdxIndicator{config: config}
}

func (indicator *AdxIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	maxCandlesCount := indicator.config.AdxDiLen * 3
	if count < maxCandlesCount {
		return false
	}

	adx := talib.Adx(
		GetHighPrice(candles, maxCandlesCount),
		GetLowPrice(candles, maxCandlesCount),
		GetClosePrice(candles, maxCandlesCount),
		indicator.config.AdxDiLen,
	)

	bottomThreshold := indicator.config.AdxBottomThreshold
	topThreshold := indicator.config.AdxTopThreshold
	adxValue := adx[len(adx)-1]

	return bottomThreshold < adxValue && adxValue < topThreshold
}

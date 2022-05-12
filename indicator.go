package main

import (
	"github.com/markcheno/go-talib"
	taindic "github.com/xyths/go-indicators"
	"math"
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

	return trend[lastTrendIdx] /*&& indicator.hadRedTrendBefore(trend)*/
}

func (indicator *SuperTrendIndicator) hadRedTrendBefore(trend []bool) bool {
	if len(trend) < 2 {
		return false
	}

	prevIdx := len(trend) - 2
	return trend[prevIdx] == false
}

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

	//inTimePeriod := 4
	//volumes := GetVolumes(candles, indicator.config.AverageVolumeCandles)
	volumes := GetSignedVolumes(candles, indicator.config.AverageVolumeCandles)
	//if len(volumes) >= (inTimePeriod * 4) {
	//	volumes = talib.Sma(volumes, inTimePeriod)
	//}

	avgVolume := GetTotal(volumes)
	//avgVolume := GetAvg(volumes)

	return avgVolume >= indicator.config.AverageVolumeMinimal
}

// Whole day total volume indicator
type WholeDayTotalVolumeIndicator struct {
	config BotConfig
}

func NewWholeDayTotalVolumeIndicator(config BotConfig) WholeDayTotalVolumeIndicator {
	return WholeDayTotalVolumeIndicator{config: config}
}

func (indicator *WholeDayTotalVolumeIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	if count < indicator.config.WholeDayTotalVolumeCandles {
		return false
	}

	volumes := GetSignedVolumes(candles, indicator.config.WholeDayTotalVolumeCandles)

	return GetTotal(volumes) >= indicator.config.WholeDayTotalVolumeMinVolume
}

// Half volume indicator
type HalfVolumeIndicator struct {
	config BotConfig
}

func NewHalfVolumeIndicator(config BotConfig) HalfVolumeIndicator {
	return HalfVolumeIndicator{config: config}
}

func (indicator *HalfVolumeIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	//needCount := indicator.config.HalfVolumeFirstCandles + indicator.config.HalfVolumeSecondCandles
	needCount := indicator.config.HalfVolumeFirstCandles + indicator.config.HalfVolumeFirstCandles
	if count < needCount {
		return false
	}

	//onlyGreen := false // true - только для зеленых, false - все

	volumes := GetVolumes(candles, needCount)
	//if onlyGreen {
	//	volumes = GetSignedVolumes(candles, needCount)
	//}

	volumesFirst := volumes[:indicator.config.HalfVolumeFirstCandles]
	volumesSecond := volumes[indicator.config.HalfVolumeFirstCandles:]

	// -->
	//if onlyGreen {
	//	volumesFirst = GetOnlyPositiveValues(volumesFirst)
	//	volumesSecond = GetOnlyPositiveValues(volumesSecond)
	//}
	// <--

	halfVolumeFirst := GetTotal(volumesFirst)
	halfVolumeSecond := GetTotal(volumesSecond)
	growth := CalcGrowth(halfVolumeFirst, halfVolumeSecond)

	return growth >= indicator.config.HalfVolumeGrowthPercentage
}

// Median volume indicator
type MedianVolumeIndicator struct {
	config BotConfig
}

func NewMedianVolumeIndicator(config BotConfig) MedianVolumeIndicator {
	return MedianVolumeIndicator{config: config}
}

func (indicator *MedianVolumeIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	if count < indicator.config.AverageVolumeCandles {
		return false
	}

	volumes := GetVolumes(candles, indicator.config.AverageVolumeCandles)
	medianVolume := Median(volumes)

	return medianVolume >= indicator.config.AverageVolumeMinimal
}

// Candle body height indicator
type CandleBodyHeightIndicator struct {
	config BotConfig
}

func NewCandleBodyHeightIndicator(config BotConfig) CandleBodyHeightIndicator {
	return CandleBodyHeightIndicator{config: config}
}

func (indicator *CandleBodyHeightIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	if count < indicator.config.CandleBodyCandles {
		return false
	}

	diffs := GetOpenClosePriceDiffs(candles, indicator.config.CandleBodyCandles)
	medianDiff := Median(diffs)

	return indicator.config.CandleBodyHeightMinPrice < medianDiff && medianDiff < indicator.config.CandleBodyHeightMaxPrice
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

	return bottomThreshold < adxValue &&
		adxValue < topThreshold /*&&
		indicator.hasGrowth(adx) &&
		indicator.config.AdxMinGrowthPercentage <= indicator.calcGrowthPercentage(adx)*/
}

func (indicator *AdxIndicator) hasGrowth(adx []float64) bool {
	candlesCount := 3
	lastAdx := adx[len(adx)-1]              // current candle adx value
	prevAdx := adx[len(adx)-1-candlesCount] // N previous candle adx value

	return prevAdx < lastAdx
}

func (indicator *AdxIndicator) calcGrowthPercentage(adx []float64) float64 {
	candlesCount := 3
	lastAdx := adx[len(adx)-1]              // current candle adx value
	prevAdx := adx[len(adx)-1-candlesCount] // N previous candle adx value

	// growth angle
	angle := (lastAdx - prevAdx) / float64(candlesCount)

	// growth percentage
	return (angle * 100) / 90
}

// Price growth indicator
type PriceGrowthIndicator struct {
	config BotConfig
}

func NewPriceGrowthIndicator(config BotConfig) PriceGrowthIndicator {
	return PriceGrowthIndicator{config: config}
}

func (indicator *PriceGrowthIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	if count < indicator.config.BtcPriceGrowthCandles {
		return false
	}

	closeCandles := GetClosePrice(candles, indicator.config.BtcPriceGrowthCandles)
	growth := CalcGrowth(closeCandles[0], closeCandles[len(closeCandles)-1])

	return indicator.config.BtcPriceGrowthMinPercentage <= growth
}

// Price fall indicator
type PriceFallIndicator struct {
	config BotConfig
}

func NewPriceFallIndicator(config BotConfig) PriceFallIndicator {
	return PriceFallIndicator{config: config}
}

func (indicator *PriceFallIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	if count < indicator.config.PriceFallCandles+1 {
		return false
	}

	closeCandles := GetClosePrice(candles, indicator.config.PriceFallCandles+1)
	fall := CalcGrowth(closeCandles[0], closeCandles[len(closeCandles)-1])

	return indicator.config.PriceFallMinPercentage <= fall // -5 > -4 ~~~> true
}

// Flat line indicator
type FlatLineIndicator struct {
	config BotConfig
}

func NewFlatLineIndicator(config BotConfig) FlatLineIndicator {
	return FlatLineIndicator{config: config}
}

func (indicator *FlatLineIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	needCount := indicator.config.FlatLineCandles + indicator.config.FlatLineSkipCandles + 1
	if count < needCount {
		return false
	}

	skipEnd := len(candles) - indicator.config.FlatLineSkipCandles - 1
	candles = candles[:skipEnd]

	//smaPeriod := indicator.getSmaPeriod(candles)
	closeCandles := GetClosePrice(candles, indicator.config.FlatLineCandles+1)
	//closeCandles = talib.Sma(talib.Sma(closeCandles, smaPeriod), 4)
	//closeCandles = closeCandles[smaPeriod:]

	onLineCount := indicator.countOnLinePrices(closeCandles, indicator.config.FlatLineDispersionPercentage, indicator.config.FlatLineCandles)
	onLinePercentage := float64(onLineCount*100) / float64(len(closeCandles))

	return onLinePercentage >= indicator.config.FlatLineOnLinePricesPercentage
}

func (indicator *FlatLineIndicator) getSmaPeriod(candles []Candle) int {
	count := len(candles)
	if count < 4 {
		return count
	}

	return 4
}

func (indicator *FlatLineIndicator) countOnLinePrices(closeCandles []float64, heightPercentage float64, period int) int {
	onLineCount := 2

	lastIdx := len(closeCandles) - 1
	firstPrice := closeCandles[0]
	lastPrice := closeCandles[lastIdx]

	for index, currentPrice := range closeCandles {
		// if first or last candle, just skip, because they are already on line
		if index == 0 || index == lastIdx {
			continue
		}

		isPriceOnLine := indicator.isPriceOnLine(currentPrice, firstPrice, lastPrice, heightPercentage, period, index)
		if isPriceOnLine {
			onLineCount++
		}
	}

	return onLineCount
}

func (indicator *FlatLineIndicator) isPriceOnLine(currentPrice, firstPrice, lastPrice, heightPercentage float64, period, currentPriceIndex int) bool {
	height := indicator.calcHeight(firstPrice, heightPercentage)
	onLinePrice := indicator.calcOnLinePrice(firstPrice, lastPrice, period, currentPriceIndex)

	return (onLinePrice-height) <= currentPrice && currentPrice <= (onLinePrice+height)
}

func (indicator *FlatLineIndicator) calcHeight(firstPrice, heightPercentage float64) float64 {
	return (firstPrice * heightPercentage) / 100
}

func (indicator *FlatLineIndicator) calcOnLinePrice(firstPrice, lastPrice float64, period, currentPriceIndex int) float64 {
	// y = kx + b
	k := indicator.calcK(firstPrice, lastPrice, period)
	b := indicator.calcB(firstPrice, k, 0)
	x := float64(currentPriceIndex)

	return math.Abs(k*x + b)
}

func (indicator *FlatLineIndicator) calcK(firstPrice, lastPrice float64, period int) float64 {
	x1 := 0.0
	x2 := float64(period - 1)

	return math.Abs((firstPrice - lastPrice) / (x1 - x2))
}

func (indicator *FlatLineIndicator) calcB(firstPrice, k float64, x int) float64 {
	return firstPrice - (k * float64(x))
}

// Flat line search indicator
type FlatLineSearchIndicator struct {
	config BotConfig
}

func NewFlatLineSearchIndicator(config BotConfig) FlatLineSearchIndicator {
	return FlatLineSearchIndicator{config: config}
}

func (indicator *FlatLineSearchIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	needCount := indicator.config.FlatLineSearchRelativePeriodCandles
	//needCount := indicator.config.FlatLineSearchWindowCandles * indicator.config.FlatLineSearchWindowsCount
	if count < needCount {
		return false
	}

	relativePrice := GetClosePrice(candles, needCount)[0]
	windowMoveNeedCount := indicator.config.FlatLineSearchWindowCandles * indicator.config.FlatLineSearchWindowsCount
	normalizedPrices := indicator.normalizePrices(append(
		[]float64{relativePrice},
		GetClosePrice(candles, windowMoveNeedCount)...,
	))

	normalizedPrices = normalizedPrices[1:]
	for i := 0; i < indicator.config.FlatLineSearchWindowsCount; i++ {
		start := i * indicator.config.FlatLineSearchWindowsCount
		end := start + indicator.config.FlatLineSearchWindowCandles

		// If we have the line, it means that we already had the line before, and we are so late to make some buys
		windowPrices := normalizedPrices[start:end]
		if indicator.isLine(windowPrices) {
			return true
		}
	}

	return false
}

func (indicator *FlatLineSearchIndicator) isLine(windowPrices []float64) bool {
	onLineCount := indicator.countOnLinePrices(
		windowPrices,
		indicator.config.FlatLineSearchDispersionPercentage,
		indicator.config.FlatLineSearchWindowCandles,
	)

	onLinePercentage := float64(onLineCount*100) / float64(indicator.config.FlatLineSearchWindowCandles)

	return onLinePercentage >= indicator.config.FlatLineSearchOnLinePricesPercentage
}

func (indicator *FlatLineSearchIndicator) countOnLinePrices(closeCandles []float64, heightPercentage float64, period int) int {
	onLineCount := 2

	lastIdx := len(closeCandles) - 1
	firstPrice := closeCandles[0]
	lastPrice := closeCandles[lastIdx]

	for index, currentPrice := range closeCandles {
		// if first or last candle, just skip, because they are already on line
		if index == 0 || index == lastIdx {
			continue
		}

		isPriceOnLine := indicator.isPriceOnLine(currentPrice, firstPrice, lastPrice, heightPercentage, period, index)
		if isPriceOnLine {
			onLineCount++
		}
	}

	return onLineCount
}

func (indicator *FlatLineSearchIndicator) isPriceOnLine(currentPrice, firstPrice, lastPrice, heightPercentage float64, period, currentPriceIndex int) bool {
	height := indicator.calcHeight(firstPrice, heightPercentage)
	onLinePrice := indicator.calcOnLinePrice(firstPrice, lastPrice, period, currentPriceIndex)

	return (onLinePrice-height) <= currentPrice && currentPrice <= (onLinePrice+height)
}

func (indicator *FlatLineSearchIndicator) calcHeight(firstPrice, heightPercentage float64) float64 {
	return (firstPrice * heightPercentage) / 100
}

func (indicator *FlatLineSearchIndicator) calcOnLinePrice(firstPrice, lastPrice float64, period, currentPriceIndex int) float64 {
	// y = kx + b
	k := indicator.calcK(firstPrice, lastPrice, period)
	b := indicator.calcB(firstPrice, k, 0)
	x := float64(currentPriceIndex)

	return math.Abs(k*x + b)
}

func (indicator *FlatLineSearchIndicator) calcK(firstPrice, lastPrice float64, period int) float64 {
	x1 := 0.0
	x2 := float64(period - 1)

	return math.Abs((firstPrice - lastPrice) / (x1 - x2))
}

func (indicator *FlatLineSearchIndicator) calcB(firstPrice, k float64, x int) float64 {
	return firstPrice - (k * float64(x))
}

func (indicator *FlatLineSearchIndicator) normalizePrices(prices []float64) []float64 {
	min := Min(prices)
	max := Max(prices)

	normalized := []float64{}

	for _, price := range prices {
		norm := (price - min) / (max - min)
		normalized = append(normalized, math.Round(norm*100))
	}

	return prices
}

// Two line indicator
type TwoLineIndicator struct {
	config BotConfig
}

func NewTwoLineIndicator(config BotConfig) TwoLineIndicator {
	return TwoLineIndicator{config: config}
}

func (indicator TwoLineIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	needCount := indicator.config.TwoLineCandles + indicator.config.TwoLineSkipCandles + 1
	if count < needCount {
		return false
	}

	closePrices := GetClosePrice(candles, needCount)
	end := len(closePrices) - indicator.config.TwoLineSkipCandles

	closePrices = closePrices[:end]
	halfCandlesCount := int(math.Round(float64(indicator.config.TwoLineCandles) / 2.0))

	avgHalfFirst := GetAvg(closePrices[:halfCandlesCount])
	avgHalfSecond := GetAvg(closePrices[halfCandlesCount:])
	diffPercentage := CalcGrowth(avgHalfFirst, avgHalfSecond)

	return 0.0 < diffPercentage && diffPercentage <= indicator.config.TwoLineMaxDiffPercentage
}

// Triple growth indicator
type TripleGrowthIndicator struct {
	config BotConfig
}

func NewTripleGrowthIndicator(config BotConfig) TripleGrowthIndicator {
	return TripleGrowthIndicator{config: config}
}

func (indicator TripleGrowthIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	needCount := indicator.config.TripleGrowthCandles
	if count < needCount {
		return false
	}

	closePrices := GetClosePrice(candles, needCount)
	half := int(math.Round(float64(needCount) / 2))

	firstHalf := closePrices[:half]
	secondHalf := closePrices[half:]

	firstGrowth := CalcGrowth(firstHalf[0], firstHalf[len(firstHalf)-1])
	secondGrowth := CalcGrowth(secondHalf[0], secondHalf[len(secondHalf)-1])

	return 0.0 <= firstGrowth && 0.0 <= secondGrowth && firstGrowth < secondGrowth
}

// Past max price indicator
type PastMaxPriceIndicator struct {
	config BotConfig
}

func NewPastMaxPriceIndicator(config BotConfig) PastMaxPriceIndicator {
	return PastMaxPriceIndicator{config: config}
}

func (indicator PastMaxPriceIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	needCount := indicator.config.PastMaxPricePeriod
	if count < needCount {
		return false
	}

	closePrices := GetClosePrice(candles, needCount)
	maxPrices := closePrices[:len(closePrices)-2]

	max := Max(maxPrices)
	currentPrice := closePrices[len(closePrices)-1]

	return max < currentPrice
}

// Smooth growth indicator
type SmoothGrowthIndicator struct {
	config BotConfig
}

func NewSmoothGrowthIndicator(config BotConfig) SmoothGrowthIndicator {
	return SmoothGrowthIndicator{config: config}
}

func (indicator SmoothGrowthIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	needCount := indicator.config.SmoothGrowthCandles
	if count < needCount {
		return false
	}

	angle := indicator.calcGrowthAngle(GetClosePrice(candles, needCount))

	return indicator.config.SmoothGrowthAngle <= angle
}

func (indicator *SmoothGrowthIndicator) calcGrowthAngle(prices []float64) float64 {
	candlesCount := len(prices)
	firstPrice := prices[0]
	lastPrice := prices[candlesCount-1]
	radians := math.Atan((lastPrice - firstPrice) / float64(candlesCount))

	return radians * (180 / math.Pi)
}

// Each volume min value indicator
type EachVolumeMinValueIndicator struct {
	config BotConfig
}

func NewEachVolumeMinValueIndicator(config BotConfig) EachVolumeMinValueIndicator {
	return EachVolumeMinValueIndicator{config: config}
}

func (indicator EachVolumeMinValueIndicator) HasBuySignal(candles []Candle) bool {
	count := len(candles)
	needCount := indicator.config.EachVolumeMinValueCandles + indicator.config.EachVolumeMinValueSkipCandles
	if count < needCount {
		return false
	}

	volumes := GetVolumes(candles, needCount)
	end := len(volumes) - indicator.config.EachVolumeMinValueSkipCandles - 1
	volumes = volumes[:end]

	for _, volume := range volumes {
		if indicator.config.EachVolumeMinValueMinVolume > volume {
			return false
		}
	}

	return true
}

package main

import (
	"github.com/montanaflynn/stats"
	"math"
	"strconv"
	"time"
)

func GetHighPrice(candles []Candle, count int) []float64 {
	var prices []float64

	firstIdx := getKlineCandleListFirstIdx(&candles, count)
	lastIdx := getKlineCandleListLastIdx(&candles)

	for _, candle := range candles[firstIdx:lastIdx] {
		prices = append(prices, candle.HighPrice)
	}

	return prices
}

func GetLowPrice(candles []Candle, count int) []float64 {
	var prices []float64

	firstIdx := getKlineCandleListFirstIdx(&candles, count)
	lastIdx := getKlineCandleListLastIdx(&candles)

	for _, candle := range candles[firstIdx:lastIdx] {
		prices = append(prices, candle.LowPrice)
	}

	return prices
}

func GetClosePrice(candles []Candle, count int) []float64 {
	var prices []float64

	firstIdx := getKlineCandleListFirstIdx(&candles, count)
	lastIdx := getKlineCandleListLastIdx(&candles)

	for _, candle := range candles[firstIdx:lastIdx] {
		prices = append(prices, candle.ClosePrice)
	}

	return prices
}

func GetVolumes(candles []Candle, count int) []float64 {
	var volumes []float64

	firstIdx := getKlineCandleListFirstIdx(&candles, count)
	lastIdx := getKlineCandleListLastIdx(&candles)

	for _, candle := range candles[firstIdx:lastIdx] {
		volumes = append(volumes, candle.Volume)
	}

	return volumes
}

func GetOnlyPositiveValues(values []float64) []float64 {
	var newValues []float64

	for _, value := range values {
		if value > 0.0 {
			newValues = append(newValues, value)
		}
	}

	return newValues
}

func GetAvg(values []float64) float64 {
	return GetTotal(values) / float64(len(values))
}

func GetTotal(values []float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}

	return total
}

func GetOpenClosePriceDiffs(candles []Candle, count int) []float64 {
	var diffs []float64

	firstIdx := getKlineCandleListFirstIdx(&candles, count)
	lastIdx := getKlineCandleListLastIdx(&candles)

	for _, candle := range candles[firstIdx:lastIdx] {
		openCloseDiff := math.Abs(CalcGrowth(candle.OpenPrice, candle.ClosePrice))
		diffs = append(diffs, openCloseDiff)
	}

	return diffs
}

func CalcGrowth(startPrice, endPrice float64) float64 {
	if startPrice == 0 || endPrice == 0 {
		return 0.0
	}

	return ((endPrice * 100) / startPrice) - 100
}

func ConvertDateStringToTime(dateString string) time.Time {
	layout := "2006-01-02 15:04:05"
	parsedTime, _ := time.Parse(layout, dateString)
	return parsedTime
}

func GetCurrentMinusTime(candleTime time.Time, minutes int) time.Time {
	//candleTime := time.Now()
	candleTime = candleTime.Add(-time.Minute * time.Duration(minutes))

	return candleTime
}

func ConvertDataFrameToBotConfig(dataFrame map[interface{}]interface{}) BotConfig {
	return BotConfig{
		HighSellPercentage: convertToFloat64(dataFrame["HighSellPercentage"]),
		LowSellPercentage:  convertToFloat64(dataFrame["LowSellPercentage"]),

		AltCoinMinBuyFirstPeriodMinutes:  convertToInt(dataFrame["AltCoinMinBuyFirstPeriodMinutes"]),
		AltCoinMinBuyFirstPercentage:     convertToFloat64(dataFrame["AltCoinMinBuyFirstPercentage"]),
		AltCoinMinBuySecondPeriodMinutes: convertToInt(dataFrame["AltCoinMinBuySecondPeriodMinutes"]),
		AltCoinMinBuySecondPercentage:    convertToFloat64(dataFrame["AltCoinMinBuySecondPercentage"]),

		BtcMinBuyPeriodMinutes: convertToInt(dataFrame["BtcMinBuyPeriodMinutes"]),
		BtcMinBuyPercentage:    convertToFloat64(dataFrame["BtcMinBuyPercentage"]),
		BtcSellPeriodMinutes:   convertToInt(dataFrame["BtcSellPeriodMinutes"]),
		BtcSellPercentage:      convertToFloat64(dataFrame["BtcSellPercentage"]),

		UnsoldFirstSellDurationMinutes: convertToInt(dataFrame["UnsoldFirstSellDurationMinutes"]),
		UnsoldFirstSellPercentage:      convertToFloat64(dataFrame["UnsoldFirstSellPercentage"]),
		UnsoldFinalSellDurationMinutes: convertToInt(dataFrame["UnsoldFinalSellDurationMinutes"]),

		AltCoinSuperTrendCandles: convertToInt(dataFrame["AltCoinSuperTrendCandles"]),
		AltCoinSuperMultiplier:   convertToFloat64(dataFrame["AltCoinSuperMultiplier"]),

		BtcSuperTrendCandles:    convertToInt(dataFrame["BtcSuperTrendCandles"]),
		BtcSuperTrendMultiplier: convertToFloat64(dataFrame["BtcSuperTrendMultiplier"]),

		AverageVolumeCandles: convertToInt(dataFrame["AverageVolumeCandles"]),
		AverageVolumeMinimal: convertToFloat64(dataFrame["AverageVolumeMinimal"]),

		AdxDiLen:               convertToInt(dataFrame["AdxDiLen"]),
		AdxBottomThreshold:     convertToFloat64(dataFrame["AdxBottomThreshold"]),
		AdxTopThreshold:        convertToFloat64(dataFrame["AdxTopThreshold"]),
		AdxMinGrowthPercentage: convertToFloat64(dataFrame["AdxMinGrowthPercentage"]),

		RealBuyTopResetReachRevenue:   convertToFloat64(dataFrame["RealBuyTopResetReachRevenue"]),
		RealBuyBottomStopReachRevenue: convertToFloat64(dataFrame["RealBuyBottomStopReachRevenue"]),
		FakeBuyReachStopRevenue:       convertToFloat64(dataFrame["FakeBuyReachStopRevenue"]),

		CandleBodyCandles:        convertToInt(dataFrame["CandleBodyCandles"]),
		CandleBodyHeightMinPrice: convertToFloat64(dataFrame["CandleBodyHeightMinPrice"]),
		CandleBodyHeightMaxPrice: convertToFloat64(dataFrame["CandleBodyHeightMaxPrice"]),

		BtcPriceGrowthCandles:       convertToInt(dataFrame["BtcPriceGrowthCandles"]),
		BtcPriceGrowthMinPercentage: convertToFloat64(dataFrame["BtcPriceGrowthMinPercentage"]),
		BtcPriceGrowthMaxPercentage: convertToFloat64(dataFrame["BtcPriceGrowthMaxPercentage"]),

		PriceFallCandles:       convertToInt(dataFrame["PriceFallCandles"]),
		PriceFallMinPercentage: convertToFloat64(dataFrame["PriceFallMinPercentage"]),

		TrailingLowPercentage: convertToFloat64(dataFrame["TrailingLowPercentage"]),

		FlatLineCandles:                convertToInt(dataFrame["FlatLineCandles"]),
		FlatLineSkipCandles:            convertToInt(dataFrame["FlatLineSkipCandles"]),
		FlatLineDispersionPercentage:   convertToFloat64(dataFrame["FlatLineDispersionPercentage"]),
		FlatLineOnLinePricesPercentage: convertToFloat64(dataFrame["FlatLineOnLinePricesPercentage"]),

		TwoLineCandles:           convertToInt(dataFrame["TwoLineCandles"]),
		TwoLineMaxDiffPercentage: convertToFloat64(dataFrame["TwoLineMaxDiffPercentage"]),
		TwoLineSkipCandles:       convertToInt(dataFrame["TwoLineSkipCandles"]),

		TrailingTopPercentage:      convertToFloat64(dataFrame["TrailingTopPercentage"]),
		TrailingReducePercentage:   convertToFloat64(dataFrame["TrailingReducePercentage"]),
		TrailingIncreasePercentage: convertToFloat64(dataFrame["TrailingIncreasePercentage"]),

		StopBuyAfterSellPeriodMinutes: convertToInt(dataFrame["StopBuyAfterSellPeriodMinutes"]),

		AltCoinMarketCandles:       convertToInt(dataFrame["AltCoinMarketCandles"]),
		AltCoinMarketMinPercentage: convertToFloat64(dataFrame["AltCoinMarketMinPercentage"]),

		AltCoinMinBuyMaxSecondPercentage: convertToFloat64(dataFrame["AltCoinMinBuyMaxSecondPercentage"]),

		WholeDayTotalVolumeCandles:   convertToInt(dataFrame["WholeDayTotalVolumeCandles"]),
		WholeDayTotalVolumeMinVolume: convertToFloat64(dataFrame["WholeDayTotalVolumeMinVolume"]),

		HalfVolumeFirstCandles:     convertToInt(dataFrame["HalfVolumeFirstCandles"]),
		HalfVolumeSecondCandles:    convertToInt(dataFrame["HalfVolumeSecondCandles"]),
		HalfVolumeGrowthPercentage: convertToFloat64(dataFrame["HalfVolumeGrowthPercentage"]),

		TrailingActivationPercentage: convertToFloat64(dataFrame["TrailingActivationPercentage"]),

		FlatLineSearchWindowCandles:          convertToInt(dataFrame["FlatLineSearchWindowCandles"]),
		FlatLineSearchWindowsCount:           convertToInt(dataFrame["FlatLineSearchWindowsCount"]),
		FlatLineSearchDispersionPercentage:   convertToFloat64(dataFrame["FlatLineSearchDispersionPercentage"]),
		FlatLineSearchOnLinePricesPercentage: convertToFloat64(dataFrame["FlatLineSearchOnLinePricesPercentage"]),
		FlatLineSearchRelativePeriodCandles:  convertToInt(dataFrame["FlatLineSearchRelativePeriodCandles"]),

		TripleGrowthCandles:          convertToInt(dataFrame["TripleGrowthCandles"]),
		TripleGrowthSecondPercentage: convertToFloat64(dataFrame["TripleGrowthSecondPercentage"]),

		PastMaxPricePeriod: convertToInt(dataFrame["PastMaxPricePeriod"]),

		SmoothGrowthCandles: convertToInt(dataFrame["SmoothGrowthCandles"]),
		SmoothGrowthAngle:   convertToFloat64(dataFrame["SmoothGrowthAngle"]),

		EachVolumeMinValueCandles:     convertToInt(dataFrame["EachVolumeMinValueCandles"]),
		EachVolumeMinValueMinVolume:   convertToFloat64(dataFrame["EachVolumeMinValueMinVolume"]),
		EachVolumeMinValueSkipCandles: convertToInt(dataFrame["EachVolumeMinValueSkipCandles"]),

		TrailingFixationActivatePercentage:  convertToFloat64(dataFrame["TrailingFixationActivatePercentage"]),
		TrailingFixationPercentage:          convertToFloat64(dataFrame["TrailingFixationPercentage"]),
		TrailingSecondaryIncreasePercentage: convertToFloat64(dataFrame["TrailingSecondaryIncreasePercentage"]),

		AltCoinMaxCandles:    convertToInt(dataFrame["AltCoinMaxCandles"]),
		AltCoinMaxPercentage: convertToFloat64(dataFrame["AltCoinMaxPercentage"]),

		WaitAfterPeriod:     convertToInt(dataFrame["WaitAfterPeriod"]),
		WaitAfterMinRevenue: convertToFloat64(dataFrame["WaitAfterMinRevenue"]),

		MinQuoteVolume: convertToFloat64(dataFrame["MinQuoteVolume"]),

		TrailingIncreaseSpeedCoefficient: convertToFloat64(dataFrame["TrailingIncreaseSpeedCoefficient"]),
		TrailingReduceSpeedCoefficient:   convertToFloat64(dataFrame["TrailingReduceSpeedCoefficient"]),
	}
}

func CountInArray(needle float64, array *[]float64) int {
	count := 0
	searchArray := *array
	for _, element := range searchArray {
		if needle == element {
			count++
		}
	}
	return count
}

func Median(values []float64) float64 {
	median, _ := stats.Median(values)
	return median
}

func CalcSuccessPercentageByRevenue(botRevenue BotRevenue) float64 {
	plusRevenue := botRevenue.PlusRevenue
	minusRevenue := botRevenue.MinusRevenue

	totalRevenue := plusRevenue + minusRevenue
	if totalRevenue == 0.0 {
		return 0.0
	}

	percentage := (plusRevenue * 100) / totalRevenue
	if percentage < 0.0 {
		return 0.0
	}

	return percentage
}

func CalcSelection(revenue, successPercentage float64) float64 {
	if revenue == DEFAULT_REVENUE {
		return DEFAULT_REVENUE
	}

	return revenue * successPercentage
}

func GetSignedVolumes(candles []Candle, count int) []float64 {
	var volumes []float64

	firstIdx := getKlineCandleListFirstIdx(&candles, count)
	lastIdx := getKlineCandleListLastIdx(&candles)

	for _, candle := range candles[firstIdx:lastIdx] {
		sign := 1.0
		if (candle.ClosePrice - candle.OpenPrice) < 0 {
			sign = -1.0
		}

		volumes = append(volumes, sign*candle.Volume)
	}

	return volumes
}

func Min(values []float64) float64 {
	if 0 == len(values) {
		panic("No values for MIN function")
	}

	min := values[0]

	for _, value := range values {
		if value < min {
			min = value
		}
	}

	return min
}

func Max(values []float64) float64 {
	if 0 == len(values) {
		panic("No values for MAX function")
	}

	max := values[0]

	for _, value := range values {
		if value > max {
			max = value
		}
	}

	return max
}

func getKlineCandleListLastIdx(candles *[]Candle) int {
	return len(*candles) - 1
}

func getKlineCandleListFirstIdx(candles *[]Candle, candlesCount int) int {
	firstIdx := len(*candles) - candlesCount - 1
	if firstIdx < 0 {
		return 0
	}

	return firstIdx
}

func convertToInt(value interface{}) int {
	switch typeValue := value.(type) {
	case int64:
		return int(typeValue)
	}
	return 0
}

func convertToFloat64(value interface{}) float64 {
	switch typeValue := value.(type) {
	case float64:
		return float64(typeValue)
	}
	return math.NaN()
}

func convertStringToFloat64(typeValue string) float64 {
	value, _ := strconv.ParseFloat(typeValue, 64)
	return value
}

func convertStringToInt(typeValue string) int {
	value, _ := strconv.ParseInt(typeValue, 10, 64)
	return int(value)
}

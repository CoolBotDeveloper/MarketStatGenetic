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

func GetAvg(values []float64) float64 {
	total := 0.0
	for _, value := range values {
		total += value
	}

	return total / float64(len(values))
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
	}
}

func InArray(needle float64, array *[]float64) bool {
	searchArray := *array
	for _, element := range searchArray {
		if needle == element {
			return true
		}
	}
	return false
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

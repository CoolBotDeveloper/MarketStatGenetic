package main

import (
	"fmt"
	"os"
	"time"
)

func CalcAvgVolume(candles []Candle) float64 {
	total := 0.0
	for _, candle := range candles {
		total += candle.Volume
	}

	return total / float64(len(candles))
}

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

func BotConfigToCsv(botConfig BotConfig, result float64, generationNum int) {
	botsCsvFile, _ := os.Create(fmt.Sprintf("generation_%d.csv", generationNum))
	content := getCsvHeader() + getCsvBody(botConfig, result)
	_, err := botsCsvFile.WriteString(content)
	if err != nil {
		fmt.Println("Error on writing generation csv file.")
	}
	botsCsvFile.Close()
}

func getCsvHeader() string {
	header := "HighSellPercentage,"
	header += "LowSellPercentage,"

	header += "AltCoinMinBuyFirstPeriodMinutes,"
	header += "AltCoinMinBuyFirstPercentage,"
	header += "AltCoinMinBuySecondPeriodMinutes,"
	header += "AltCoinMinBuySecondPercentage,"

	header += "BtcMinBuyPeriodMinutes,"
	header += "BtcMinBuyPercentage,"
	header += "BtcSellPeriodMinutes,"
	header += "BtcSellPercentage,"

	header += "UnsoldFirstSellDurationMinutes,"
	header += "UnsoldFirstSellPercentage,"
	header += "UnsoldFinalSellDurationMinutes,"

	header += "AltCoinSuperTrendCandles,"
	header += "AltCoinSuperMultiplier,"

	header += "BtcSuperTrendCandles,"
	header += "BtcSuperTrendMultiplier,"

	header += "AverageVolumeCandles,"
	header += "AverageVolumeMinimal,"

	header += "AdxDiLen,"
	header += "AdxBottomThreshold,"
	header += "AdxTopThreshold,"
	header += "Result\n"

	return header
}

func getCsvBody(botConfig BotConfig, result float64) string {
	body := fmt.Sprintf("%f,", botConfig.HighSellPercentage)
	body += fmt.Sprintf("%f,", botConfig.LowSellPercentage)

	body += fmt.Sprintf("%d,", botConfig.AltCoinMinBuyFirstPeriodMinutes)
	body += fmt.Sprintf("%f,", botConfig.AltCoinMinBuyFirstPercentage)
	body += fmt.Sprintf("%d,", botConfig.AltCoinMinBuySecondPeriodMinutes)
	body += fmt.Sprintf("%f,", botConfig.AltCoinMinBuySecondPercentage)

	body += fmt.Sprintf("%d,", botConfig.BtcMinBuyPeriodMinutes)
	body += fmt.Sprintf("%f,", botConfig.BtcMinBuyPercentage)
	body += fmt.Sprintf("%d,", botConfig.BtcSellPeriodMinutes)
	body += fmt.Sprintf("%f,", botConfig.BtcSellPercentage)

	body += fmt.Sprintf("%d,", botConfig.UnsoldFirstSellDurationMinutes)
	body += fmt.Sprintf("%f,", botConfig.UnsoldFirstSellPercentage)
	body += fmt.Sprintf("%d,", botConfig.UnsoldFinalSellDurationMinutes)

	body += fmt.Sprintf("%d,", botConfig.AltCoinSuperTrendCandles)
	body += fmt.Sprintf("%f,", botConfig.AltCoinSuperMultiplier)

	body += fmt.Sprintf("%d,", botConfig.BtcSuperTrendCandles)
	body += fmt.Sprintf("%f,", botConfig.BtcSuperTrendMultiplier)

	body += fmt.Sprintf("%d,", botConfig.AverageVolumeCandles)
	body += fmt.Sprintf("%f,", botConfig.AverageVolumeMinimal)

	body += fmt.Sprintf("%d,", botConfig.AdxDiLen)
	body += fmt.Sprintf("%f,", botConfig.AdxBottomThreshold)
	body += fmt.Sprintf("%f,", botConfig.AdxTopThreshold)
	body += fmt.Sprintf("%f\n", result)

	return body
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

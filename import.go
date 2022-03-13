package main

import (
	"encoding/csv"
	"os"
)

func ImportFromCsv(fileName string) []BotConfig {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Can not load initial bots from file.")
	}

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()

	var bots []BotConfig
	for rowNumber, row := range rows {
		if rowNumber == 0 {
			continue
		}

		bot := BotConfig{
			HighSellPercentage: convertStringToFloat64(row[0]),
			LowSellPercentage:  convertStringToFloat64(row[1]),

			AltCoinMinBuyFirstPeriodMinutes:  convertStringToInt(row[2]),
			AltCoinMinBuyFirstPercentage:     convertStringToFloat64(row[3]),
			AltCoinMinBuySecondPeriodMinutes: convertStringToInt(row[4]),
			AltCoinMinBuySecondPercentage:    convertStringToFloat64(row[5]),

			BtcMinBuyPeriodMinutes: convertStringToInt(row[6]),
			BtcMinBuyPercentage:    convertStringToFloat64(row[7]),
			BtcSellPeriodMinutes:   convertStringToInt(row[8]),
			BtcSellPercentage:      convertStringToFloat64(row[9]),

			UnsoldFirstSellDurationMinutes: convertStringToInt(row[10]),
			UnsoldFirstSellPercentage:      convertStringToFloat64(row[11]),
			UnsoldFinalSellDurationMinutes: convertStringToInt(row[12]),

			AltCoinSuperTrendCandles: convertStringToInt(row[13]),
			AltCoinSuperMultiplier:   convertStringToFloat64(row[14]),

			BtcSuperTrendCandles:    convertStringToInt(row[15]),
			BtcSuperTrendMultiplier: convertStringToFloat64(row[16]),

			AverageVolumeCandles: convertStringToInt(row[17]),
			AverageVolumeMinimal: convertStringToFloat64(row[18]),

			AdxDiLen:           convertStringToInt(row[19]),
			AdxBottomThreshold: convertStringToFloat64(row[20]),
			AdxTopThreshold:    convertStringToFloat64(row[21]),

			RealBuyTopResetReachRevenue:   convertStringToFloat64(row[22]),
			RealBuyBottomStopReachRevenue: convertStringToFloat64(row[23]),
			FakeBuyReachStopRevenue:       convertStringToFloat64(row[24]),

			CandleBodyCandles:        convertStringToInt(row[25]),
			CandleBodyHeightMinPrice: convertStringToFloat64(row[26]),
			CandleBodyHeightMaxPrice: convertStringToFloat64(row[27]),

			TotalRevenue: convertStringToFloat64(row[28]),
		}

		bots = append(bots, bot)
	}

	return bots
}

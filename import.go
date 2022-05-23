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

			AdxDiLen:               convertStringToInt(row[19]),
			AdxBottomThreshold:     convertStringToFloat64(row[20]),
			AdxTopThreshold:        convertStringToFloat64(row[21]),
			AdxMinGrowthPercentage: convertStringToFloat64(row[22]),

			RealBuyTopResetReachRevenue:   convertStringToFloat64(row[23]),
			RealBuyBottomStopReachRevenue: convertStringToFloat64(row[24]),
			FakeBuyReachStopRevenue:       convertStringToFloat64(row[25]),

			CandleBodyCandles:        convertStringToInt(row[26]),
			CandleBodyHeightMinPrice: convertStringToFloat64(row[27]),
			CandleBodyHeightMaxPrice: convertStringToFloat64(row[28]),

			BtcPriceGrowthCandles:       convertStringToInt(row[29]),
			BtcPriceGrowthMinPercentage: convertStringToFloat64(row[30]),
			BtcPriceGrowthMaxPercentage: convertStringToFloat64(row[31]),

			PriceFallCandles:       convertStringToInt(row[32]),
			PriceFallMinPercentage: convertStringToFloat64(row[33]),

			TrailingLowPercentage: convertStringToFloat64(row[34]),

			FlatLineCandles:                convertStringToInt(row[35]),
			FlatLineSkipCandles:            convertStringToInt(row[36]),
			FlatLineDispersionPercentage:   convertStringToFloat64(row[37]),
			FlatLineOnLinePricesPercentage: convertStringToFloat64(row[38]),

			TwoLineCandles:           convertStringToInt(row[39]),
			TwoLineMaxDiffPercentage: convertStringToFloat64(row[40]),
			TwoLineSkipCandles:       convertStringToInt(row[41]),

			TrailingTopPercentage:      convertStringToFloat64(row[42]),
			TrailingReducePercentage:   convertStringToFloat64(row[43]),
			TrailingIncreasePercentage: convertStringToFloat64(row[44]),

			StopBuyAfterSellPeriodMinutes: convertStringToInt(row[45]),

			AltCoinMarketCandles:       convertStringToInt(row[46]),
			AltCoinMarketMinPercentage: convertStringToFloat64(row[47]),

			AltCoinMinBuyMaxSecondPercentage: convertStringToFloat64(row[48]),

			WholeDayTotalVolumeCandles:   convertStringToInt(row[49]),
			WholeDayTotalVolumeMinVolume: convertStringToFloat64(row[50]),

			HalfVolumeFirstCandles:     convertStringToInt(row[51]),
			HalfVolumeSecondCandles:    convertStringToInt(row[52]),
			HalfVolumeGrowthPercentage: convertStringToFloat64(row[53]),

			TrailingActivationPercentage: convertStringToFloat64(row[54]),

			FlatLineSearchWindowCandles:          convertStringToInt(row[55]),
			FlatLineSearchWindowsCount:           convertStringToInt(row[56]),
			FlatLineSearchDispersionPercentage:   convertStringToFloat64(row[57]),
			FlatLineSearchOnLinePricesPercentage: convertStringToFloat64(row[58]),
			FlatLineSearchRelativePeriodCandles:  convertStringToInt(row[59]),

			TripleGrowthCandles:          convertStringToInt(row[60]),
			TripleGrowthSecondPercentage: convertStringToFloat64(row[61]),

			PastMaxPricePeriod: convertStringToInt(row[62]),

			SmoothGrowthCandles: convertStringToInt(row[63]),
			SmoothGrowthAngle:   convertStringToFloat64(row[64]),

			EachVolumeMinValueCandles:     convertStringToInt(row[65]),
			EachVolumeMinValueMinVolume:   convertStringToFloat64(row[66]),
			EachVolumeMinValueSkipCandles: convertStringToInt(row[67]),

			TrailingFixationActivatePercentage:  convertStringToFloat64(row[68]),
			TrailingFixationPercentage:          convertStringToFloat64(row[69]),
			TrailingSecondaryIncreasePercentage: convertStringToFloat64(row[70]),

			AltCoinMaxCandles:    convertStringToInt(row[71]),
			AltCoinMaxPercentage: convertStringToFloat64(row[72]),

			TotalRevenue:      convertStringToFloat64(row[73]),
			SuccessPercentage: convertStringToFloat64(row[74]),
			PlusRevenue:       convertStringToFloat64(row[75]),
			MinusRevenue:      convertStringToFloat64(row[76]),
			Selection:         convertStringToFloat64(row[77]),
		}

		bots = append(bots, bot)
	}

	return bots
}

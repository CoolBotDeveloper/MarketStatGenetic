package main

import (
	"context"
	"github.com/rocketlaunchr/dataframe-go"
	"math/rand"
)

const BEST_BOTS_COUNT = 7
const BEST_BOTS_FROM_PREV_GEN = 3
const BOTS_COUNT = 25
const GENERATION_COUNT = 2000
const DEFAULT_REVENUE = -10000000

func InitBotsDataFrame() *dataframe.DataFrame {
	return dataframe.NewDataFrame(
		dataframe.NewSeriesFloat64("HighSellPercentage", nil),
		dataframe.NewSeriesFloat64("LowSellPercentage", nil),

		dataframe.NewSeriesInt64("AltCoinMinBuyFirstPeriodMinutes", nil),
		dataframe.NewSeriesFloat64("AltCoinMinBuyFirstPercentage", nil),
		dataframe.NewSeriesInt64("AltCoinMinBuySecondPeriodMinutes", nil),
		dataframe.NewSeriesFloat64("AltCoinMinBuySecondPercentage", nil),

		dataframe.NewSeriesInt64("BtcMinBuyPeriodMinutes", nil),
		dataframe.NewSeriesFloat64("BtcMinBuyPercentage", nil),
		dataframe.NewSeriesInt64("BtcSellPeriodMinutes", nil),
		dataframe.NewSeriesFloat64("BtcSellPercentage", nil),

		dataframe.NewSeriesInt64("UnsoldFirstSellDurationMinutes", nil),
		dataframe.NewSeriesFloat64("UnsoldFirstSellPercentage", nil),
		dataframe.NewSeriesInt64("UnsoldFinalSellDurationMinutes", nil),

		dataframe.NewSeriesInt64("AltCoinSuperTrendCandles", nil),
		dataframe.NewSeriesFloat64("AltCoinSuperMultiplier", nil),

		dataframe.NewSeriesInt64("BtcSuperTrendCandles", nil),
		dataframe.NewSeriesFloat64("BtcSuperTrendMultiplier", nil),

		dataframe.NewSeriesInt64("AverageVolumeCandles", nil),
		dataframe.NewSeriesFloat64("AverageVolumeMinimal", nil),

		dataframe.NewSeriesInt64("AdxDiLen", nil),
		dataframe.NewSeriesFloat64("AdxBottomThreshold", nil),
		dataframe.NewSeriesFloat64("AdxTopThreshold", nil),
		dataframe.NewSeriesFloat64("AdxMinGrowthPercentage", nil),

		dataframe.NewSeriesFloat64("RealBuyTopResetReachRevenue", nil),
		dataframe.NewSeriesFloat64("RealBuyBottomStopReachRevenue", nil),
		dataframe.NewSeriesFloat64("FakeBuyReachStopRevenue", nil),

		dataframe.NewSeriesInt64("CandleBodyCandles", nil),
		dataframe.NewSeriesFloat64("CandleBodyHeightMinPrice", nil),
		dataframe.NewSeriesFloat64("CandleBodyHeightMaxPrice", nil),

		dataframe.NewSeriesInt64("BtcPriceGrowthCandles", nil),
		dataframe.NewSeriesFloat64("BtcPriceGrowthMinPercentage", nil),
		dataframe.NewSeriesFloat64("BtcPriceGrowthMaxPercentage", nil),

		dataframe.NewSeriesInt64("PriceFallCandles", nil),
		dataframe.NewSeriesFloat64("PriceFallMinPercentage", nil),

		dataframe.NewSeriesFloat64("TrailingLowPercentage", nil),

		dataframe.NewSeriesInt64("FlatLineCandles", nil),
		dataframe.NewSeriesInt64("FlatLineSkipCandles", nil),
		dataframe.NewSeriesFloat64("FlatLineDispersionPercentage", nil),
		dataframe.NewSeriesFloat64("FlatLineOnLinePricesPercentage", nil),

		dataframe.NewSeriesInt64("TwoLineCandles", nil),
		dataframe.NewSeriesFloat64("TwoLineMaxDiffPercentage", nil),
		dataframe.NewSeriesInt64("TwoLineSkipCandles", nil),

		dataframe.NewSeriesFloat64("TrailingTopPercentage", nil),
		dataframe.NewSeriesFloat64("TrailingReducePercentage", nil),
		dataframe.NewSeriesFloat64("TrailingIncreasePercentage", nil),

		dataframe.NewSeriesInt64("StopBuyAfterSellPeriodMinutes", nil),

		dataframe.NewSeriesInt64("AltCoinMarketCandles", nil),
		dataframe.NewSeriesFloat64("AltCoinMarketMinPercentage", nil),

		dataframe.NewSeriesFloat64("AltCoinMinBuyMaxSecondPercentage", nil),

		dataframe.NewSeriesInt64("WholeDayTotalVolumeCandles", nil),
		dataframe.NewSeriesFloat64("WholeDayTotalVolumeMinVolume", nil),

		dataframe.NewSeriesInt64("HalfVolumeFirstCandles", nil),
		dataframe.NewSeriesInt64("HalfVolumeSecondCandles", nil),
		dataframe.NewSeriesFloat64("HalfVolumeGrowthPercentage", nil),

		dataframe.NewSeriesFloat64("TrailingActivationPercentage", nil),

		dataframe.NewSeriesInt64("FlatLineSearchWindowCandles", nil),
		dataframe.NewSeriesInt64("FlatLineSearchWindowsCount", nil),
		dataframe.NewSeriesFloat64("FlatLineSearchDispersionPercentage", nil),
		dataframe.NewSeriesFloat64("FlatLineSearchOnLinePricesPercentage", nil),
		dataframe.NewSeriesInt64("FlatLineSearchRelativePeriodCandles", nil),

		dataframe.NewSeriesInt64("TripleGrowthCandles", nil),
		dataframe.NewSeriesFloat64("TripleGrowthSecondPercentage", nil),

		dataframe.NewSeriesInt64("PastMaxPricePeriod", nil),

		dataframe.NewSeriesInt64("SmoothGrowthCandles", nil),
		dataframe.NewSeriesFloat64("SmoothGrowthAngle", nil),

		dataframe.NewSeriesInt64("EachVolumeMinValueCandles", nil),
		dataframe.NewSeriesFloat64("EachVolumeMinValueMinVolume", nil),
		dataframe.NewSeriesInt64("EachVolumeMinValueSkipCandles", nil),

		dataframe.NewSeriesFloat64("TrailingFixationActivatePercentage", nil),
		dataframe.NewSeriesFloat64("TrailingFixationPercentage", nil),
		dataframe.NewSeriesFloat64("TrailingSecondaryIncreasePercentage", nil),

		dataframe.NewSeriesInt64("AltCoinMaxCandles", nil),
		dataframe.NewSeriesFloat64("AltCoinMaxPercentage", nil),

		dataframe.NewSeriesFloat64("TotalRevenue", nil),
		dataframe.NewSeriesFloat64("SuccessPercentage", nil),
		dataframe.NewSeriesFloat64("PlusRevenue", nil),
		dataframe.NewSeriesFloat64("MinusRevenue", nil),

		dataframe.NewSeriesFloat64("Selection", nil),
	)
}

func GetInitialBots() *dataframe.DataFrame {
	initialBotsDataFrame := InitBotsDataFrame()
	for botNumber := 0; botNumber < BOTS_COUNT; botNumber++ {
		botConfig := InitBotConfig()
		initialBotsDataFrame.Append(nil, GetBotConfigMapInterface(botConfig))
	}
	return initialBotsDataFrame
}

func GetInitialBotsFromFile(fileName string) *dataframe.DataFrame {
	initialBotsDataFrame := InitBotsDataFrame()
	csvBotConfigs := ImportFromCsv(fileName)
	for _, botConfig := range csvBotConfigs {
		initialBotsDataFrame.Append(nil, GetBotConfigMapInterface(botConfig))
	}
	return initialBotsDataFrame
}

func SetBotTotalRevenue(
	bots *dataframe.DataFrame,
	botNumber int, revenue,
	successPercentage float64,
	plusRevenue float64,
	minusRevenue float64,
) {
	bots.UpdateRow(botNumber, nil, map[string]interface{}{
		"TotalRevenue":      revenue,
		"SuccessPercentage": successPercentage,
		"PlusRevenue":       plusRevenue,
		"MinusRevenue":      minusRevenue,
		"Selection":         CalcSelection(revenue, successPercentage),
	})
}

func SortBestBots(bots *dataframe.DataFrame) *dataframe.DataFrame {
	sks := []dataframe.SortKey{
		//{
		//	Key:  "TotalRevenue",
		//	Desc: true,
		//},
		//{
		//	Key:  "SuccessPercentage",
		//	Desc: true,
		//},
		{
			Key:  "Selection",
			Desc: true,
		},
	}
	ctx := context.Background()
	bots.Sort(ctx, sks)
	return bots
}

func SelectNBots(numberOfBots int, bots *dataframe.DataFrame) *dataframe.DataFrame {
	botsDataFrame := InitBotsDataFrame()
	iterator := bots.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
	alreadyHasRevenue := []float64{}

	for {
		botNumber, bot, _ := iterator()
		if botNumber == nil || numberOfBots < botsDataFrame.NRows()+1 {
			break
		}

		botRevenue := convertToFloat64(bot["TotalRevenue"])
		if 0 < CountInArray(botRevenue, &alreadyHasRevenue) {
			continue
		}

		if botRevenue != 0.0 && botRevenue != DEFAULT_REVENUE {
			alreadyHasRevenue = append(alreadyHasRevenue, botRevenue)
		}

		botsDataFrame.Append(nil, createBotDataFrameRow(bot))
	}
	return botsDataFrame
}

func CombineParentAndChildBots(
	parentBots *dataframe.DataFrame,
	childBots *dataframe.DataFrame,
) *dataframe.DataFrame {
	botsDataFrame := InitBotsDataFrame()
	queueBots := []*dataframe.DataFrame{
		parentBots,
		childBots,
	}

	for _, bots := range queueBots {
		iterator := bots.ValuesIterator(dataframe.ValuesOptions{0, 1, true})

		for {
			botNumber, bot, _ := iterator()
			if botNumber == nil {
				break
			}
			botsDataFrame.Append(nil, createBotDataFrameRow(bot))
		}
	}

	return botsDataFrame
}

func createBotDataFrameRow(bot map[interface{}]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"HighSellPercentage": bot["HighSellPercentage"],
		"LowSellPercentage":  bot["LowSellPercentage"],

		"AltCoinMinBuyFirstPeriodMinutes":  bot["AltCoinMinBuyFirstPeriodMinutes"],
		"AltCoinMinBuyFirstPercentage":     bot["AltCoinMinBuyFirstPercentage"],
		"AltCoinMinBuySecondPeriodMinutes": bot["AltCoinMinBuySecondPeriodMinutes"],
		"AltCoinMinBuySecondPercentage":    bot["AltCoinMinBuySecondPercentage"],

		"BtcMinBuyPeriodMinutes": bot["BtcMinBuyPeriodMinutes"],
		"BtcMinBuyPercentage":    bot["BtcMinBuyPercentage"],
		"BtcSellPeriodMinutes":   bot["BtcSellPeriodMinutes"],
		"BtcSellPercentage":      bot["BtcSellPercentage"],

		"UnsoldFirstSellDurationMinutes": bot["UnsoldFirstSellDurationMinutes"],
		"UnsoldFirstSellPercentage":      bot["UnsoldFirstSellPercentage"],
		"UnsoldFinalSellDurationMinutes": bot["UnsoldFinalSellDurationMinutes"],

		"AltCoinSuperTrendCandles": bot["AltCoinSuperTrendCandles"],
		"AltCoinSuperMultiplier":   bot["AltCoinSuperMultiplier"],

		"BtcSuperTrendCandles":    bot["BtcSuperTrendCandles"],
		"BtcSuperTrendMultiplier": bot["BtcSuperTrendMultiplier"],

		"AverageVolumeCandles": bot["AverageVolumeCandles"],
		"AverageVolumeMinimal": bot["AverageVolumeMinimal"],

		"AdxDiLen":               bot["AdxDiLen"],
		"AdxBottomThreshold":     bot["AdxBottomThreshold"],
		"AdxTopThreshold":        bot["AdxTopThreshold"],
		"AdxMinGrowthPercentage": bot["AdxMinGrowthPercentage"],

		"RealBuyTopResetReachRevenue":   bot["RealBuyTopResetReachRevenue"],
		"RealBuyBottomStopReachRevenue": bot["RealBuyBottomStopReachRevenue"],
		"FakeBuyReachStopRevenue":       bot["FakeBuyReachStopRevenue"],

		"CandleBodyCandles":        bot["CandleBodyCandles"],
		"CandleBodyHeightMinPrice": bot["CandleBodyHeightMinPrice"],
		"CandleBodyHeightMaxPrice": bot["CandleBodyHeightMaxPrice"],

		"BtcPriceGrowthCandles":       bot["BtcPriceGrowthCandles"],
		"BtcPriceGrowthMinPercentage": bot["BtcPriceGrowthMinPercentage"],
		"BtcPriceGrowthMaxPercentage": bot["BtcPriceGrowthMaxPercentage"],

		"PriceFallCandles":       bot["PriceFallCandles"],
		"PriceFallMinPercentage": bot["PriceFallMinPercentage"],

		"TrailingLowPercentage": bot["TrailingLowPercentage"],

		"FlatLineCandles":                bot["FlatLineCandles"],
		"FlatLineSkipCandles":            bot["FlatLineSkipCandles"],
		"FlatLineDispersionPercentage":   bot["FlatLineDispersionPercentage"],
		"FlatLineOnLinePricesPercentage": bot["FlatLineOnLinePricesPercentage"],

		"TwoLineCandles":           bot["TwoLineCandles"],
		"TwoLineMaxDiffPercentage": bot["TwoLineMaxDiffPercentage"],
		"TwoLineSkipCandles":       bot["TwoLineSkipCandles"],

		"TrailingTopPercentage":      bot["TrailingTopPercentage"],
		"TrailingReducePercentage":   bot["TrailingReducePercentage"],
		"TrailingIncreasePercentage": bot["TrailingIncreasePercentage"],

		"StopBuyAfterSellPeriodMinutes": bot["StopBuyAfterSellPeriodMinutes"],

		"AltCoinMarketCandles":       bot["AltCoinMarketCandles"],
		"AltCoinMarketMinPercentage": bot["AltCoinMarketMinPercentage"],

		"AltCoinMinBuyMaxSecondPercentage": bot["AltCoinMinBuyMaxSecondPercentage"],

		"WholeDayTotalVolumeCandles":   bot["WholeDayTotalVolumeCandles"],
		"WholeDayTotalVolumeMinVolume": bot["WholeDayTotalVolumeMinVolume"],

		"HalfVolumeFirstCandles":     bot["HalfVolumeFirstCandles"],
		"HalfVolumeSecondCandles":    bot["HalfVolumeSecondCandles"],
		"HalfVolumeGrowthPercentage": bot["HalfVolumeGrowthPercentage"],

		"TrailingActivationPercentage": bot["TrailingActivationPercentage"],

		"FlatLineSearchWindowCandles":          bot["FlatLineSearchWindowCandles"],
		"FlatLineSearchWindowsCount":           bot["FlatLineSearchWindowsCount"],
		"FlatLineSearchDispersionPercentage":   bot["FlatLineSearchDispersionPercentage"],
		"FlatLineSearchOnLinePricesPercentage": bot["FlatLineSearchOnLinePricesPercentage"],
		"FlatLineSearchRelativePeriodCandles":  bot["FlatLineSearchRelativePeriodCandles"],

		"TripleGrowthCandles":          bot["TripleGrowthCandles"],
		"TripleGrowthSecondPercentage": bot["TripleGrowthSecondPercentage"],

		"PastMaxPricePeriod": bot["PastMaxPricePeriod"],

		"SmoothGrowthCandles": bot["SmoothGrowthCandles"],
		"SmoothGrowthAngle":   bot["SmoothGrowthAngle"],

		"EachVolumeMinValueCandles":     bot["EachVolumeMinValueCandles"],
		"EachVolumeMinValueMinVolume":   bot["EachVolumeMinValueMinVolume"],
		"EachVolumeMinValueSkipCandles": bot["EachVolumeMinValueSkipCandles"],

		"TrailingFixationActivatePercentage":  bot["TrailingFixationActivatePercentage"],
		"TrailingFixationPercentage":          bot["TrailingFixationPercentage"],
		"TrailingSecondaryIncreasePercentage": bot["TrailingSecondaryIncreasePercentage"],

		"AltCoinMaxCandles":    bot["AltCoinMaxCandles"],
		"AltCoinMaxPercentage": bot["AltCoinMaxPercentage"],

		"TotalRevenue":      bot["TotalRevenue"],
		"SuccessPercentage": bot["SuccessPercentage"],
		"PlusRevenue":       bot["PlusRevenue"],
		"MinusRevenue":      bot["MinusRevenue"],

		"Selection": bot["Selection"],
	}
}

func MakeChildren(parentBots *dataframe.DataFrame) *dataframe.DataFrame {
	childrenBots := InitBotsDataFrame()
	maleIterator := parentBots.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
	for {
		maleBotNumber, maleBot, _ := maleIterator()
		if maleBotNumber == nil {
			break
		}

		femaleIterator := parentBots.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
		for {
			femaleBotNumber, femaleBot, _ := femaleIterator()
			if femaleBotNumber == nil {
				break
			}

			if *maleBotNumber == *femaleBotNumber {
				continue
			}

			child := makeChild(
				ConvertDataFrameToBotConfig(maleBot),
				ConvertDataFrameToBotConfig(femaleBot),
			)

			childrenBots.Append(nil, child)
		}
	}
	childrenBots = shuffleBots(childrenBots)
	return SelectNBots(BOTS_COUNT, childrenBots)
}

func makeChild(
	maleBotConfig BotConfig,
	femaleBotConfig BotConfig,
) map[string]interface{} {
	childBotConfig := BotConfig{
		HighSellPercentage: GetFloatFatherOrMomGen(maleBotConfig.HighSellPercentage, femaleBotConfig.HighSellPercentage),
		LowSellPercentage:  GetFloatFatherOrMomGen(maleBotConfig.LowSellPercentage, femaleBotConfig.LowSellPercentage),

		AltCoinMinBuyFirstPeriodMinutes:  GetIntFatherOrMomGen(maleBotConfig.AltCoinMinBuyFirstPeriodMinutes, femaleBotConfig.AltCoinMinBuyFirstPeriodMinutes),
		AltCoinMinBuyFirstPercentage:     GetFloatFatherOrMomGen(maleBotConfig.AltCoinMinBuyFirstPercentage, femaleBotConfig.AltCoinMinBuyFirstPercentage),
		AltCoinMinBuySecondPeriodMinutes: GetIntFatherOrMomGen(maleBotConfig.AltCoinMinBuySecondPeriodMinutes, femaleBotConfig.AltCoinMinBuySecondPeriodMinutes),
		AltCoinMinBuySecondPercentage:    GetFloatFatherOrMomGen(maleBotConfig.AltCoinMinBuySecondPercentage, femaleBotConfig.AltCoinMinBuySecondPercentage),

		BtcMinBuyPeriodMinutes: GetIntFatherOrMomGen(maleBotConfig.BtcMinBuyPeriodMinutes, femaleBotConfig.BtcMinBuyPeriodMinutes),
		BtcMinBuyPercentage:    GetFloatFatherOrMomGen(maleBotConfig.BtcMinBuyPercentage, femaleBotConfig.BtcMinBuyPercentage),
		BtcSellPeriodMinutes:   GetIntFatherOrMomGen(maleBotConfig.BtcSellPeriodMinutes, femaleBotConfig.BtcSellPeriodMinutes),
		BtcSellPercentage:      GetFloatFatherOrMomGen(maleBotConfig.BtcSellPercentage, femaleBotConfig.BtcSellPercentage),

		UnsoldFirstSellDurationMinutes: GetIntFatherOrMomGen(maleBotConfig.UnsoldFirstSellDurationMinutes, femaleBotConfig.UnsoldFirstSellDurationMinutes),
		UnsoldFirstSellPercentage:      GetFloatFatherOrMomGen(maleBotConfig.UnsoldFirstSellPercentage, femaleBotConfig.UnsoldFirstSellPercentage),
		UnsoldFinalSellDurationMinutes: GetIntFatherOrMomGen(maleBotConfig.UnsoldFinalSellDurationMinutes, femaleBotConfig.UnsoldFinalSellDurationMinutes),

		AltCoinSuperTrendCandles: GetIntFatherOrMomGen(maleBotConfig.AltCoinSuperTrendCandles, femaleBotConfig.AltCoinSuperTrendCandles),
		AltCoinSuperMultiplier:   GetFloatFatherOrMomGen(maleBotConfig.AltCoinSuperMultiplier, femaleBotConfig.AltCoinSuperMultiplier),

		BtcSuperTrendCandles:    GetIntFatherOrMomGen(maleBotConfig.BtcSuperTrendCandles, femaleBotConfig.BtcSuperTrendCandles),
		BtcSuperTrendMultiplier: GetFloatFatherOrMomGen(maleBotConfig.BtcSuperTrendMultiplier, femaleBotConfig.BtcSuperTrendMultiplier),

		AverageVolumeCandles: GetIntFatherOrMomGen(maleBotConfig.AverageVolumeCandles, femaleBotConfig.AverageVolumeCandles),
		AverageVolumeMinimal: GetFloatFatherOrMomGen(maleBotConfig.AverageVolumeMinimal, femaleBotConfig.AverageVolumeMinimal),

		AdxDiLen:               GetIntFatherOrMomGen(maleBotConfig.AdxDiLen, femaleBotConfig.AdxDiLen),
		AdxBottomThreshold:     GetFloatFatherOrMomGen(maleBotConfig.AdxBottomThreshold, femaleBotConfig.AdxBottomThreshold),
		AdxTopThreshold:        GetFloatFatherOrMomGen(maleBotConfig.AdxTopThreshold, femaleBotConfig.AdxTopThreshold),
		AdxMinGrowthPercentage: GetFloatFatherOrMomGen(maleBotConfig.AdxMinGrowthPercentage, femaleBotConfig.AdxMinGrowthPercentage),

		RealBuyTopResetReachRevenue:   GetFloatFatherOrMomGen(maleBotConfig.RealBuyTopResetReachRevenue, femaleBotConfig.RealBuyTopResetReachRevenue),
		RealBuyBottomStopReachRevenue: GetFloatFatherOrMomGen(maleBotConfig.RealBuyBottomStopReachRevenue, femaleBotConfig.RealBuyBottomStopReachRevenue),
		FakeBuyReachStopRevenue:       GetFloatFatherOrMomGen(maleBotConfig.FakeBuyReachStopRevenue, femaleBotConfig.FakeBuyReachStopRevenue),

		CandleBodyCandles:        GetIntFatherOrMomGen(maleBotConfig.CandleBodyCandles, femaleBotConfig.CandleBodyCandles),
		CandleBodyHeightMinPrice: GetFloatFatherOrMomGen(maleBotConfig.CandleBodyHeightMinPrice, femaleBotConfig.CandleBodyHeightMinPrice),
		CandleBodyHeightMaxPrice: GetFloatFatherOrMomGen(maleBotConfig.CandleBodyHeightMaxPrice, femaleBotConfig.CandleBodyHeightMaxPrice),

		BtcPriceGrowthCandles:       GetIntFatherOrMomGen(maleBotConfig.BtcPriceGrowthCandles, femaleBotConfig.BtcPriceGrowthCandles),
		BtcPriceGrowthMinPercentage: GetFloatFatherOrMomGen(maleBotConfig.BtcPriceGrowthMinPercentage, femaleBotConfig.BtcPriceGrowthMinPercentage),
		BtcPriceGrowthMaxPercentage: GetFloatFatherOrMomGen(maleBotConfig.BtcPriceGrowthMaxPercentage, femaleBotConfig.BtcPriceGrowthMaxPercentage),

		PriceFallCandles:       GetIntFatherOrMomGen(maleBotConfig.PriceFallCandles, femaleBotConfig.PriceFallCandles),
		PriceFallMinPercentage: GetFloatFatherOrMomGen(maleBotConfig.PriceFallMinPercentage, femaleBotConfig.PriceFallMinPercentage),

		TrailingLowPercentage: GetFloatFatherOrMomGen(maleBotConfig.TrailingLowPercentage, femaleBotConfig.TrailingLowPercentage),

		FlatLineCandles:                GetIntFatherOrMomGen(maleBotConfig.FlatLineCandles, femaleBotConfig.FlatLineCandles),
		FlatLineSkipCandles:            GetIntFatherOrMomGen(maleBotConfig.FlatLineSkipCandles, femaleBotConfig.FlatLineSkipCandles),
		FlatLineDispersionPercentage:   GetFloatFatherOrMomGen(maleBotConfig.FlatLineDispersionPercentage, femaleBotConfig.FlatLineDispersionPercentage),
		FlatLineOnLinePricesPercentage: GetFloatFatherOrMomGen(maleBotConfig.FlatLineOnLinePricesPercentage, femaleBotConfig.FlatLineOnLinePricesPercentage),

		TwoLineCandles:           GetIntFatherOrMomGen(maleBotConfig.TwoLineCandles, femaleBotConfig.TwoLineCandles),
		TwoLineMaxDiffPercentage: GetFloatFatherOrMomGen(maleBotConfig.TwoLineMaxDiffPercentage, femaleBotConfig.TwoLineMaxDiffPercentage),
		TwoLineSkipCandles:       GetIntFatherOrMomGen(maleBotConfig.TwoLineSkipCandles, femaleBotConfig.TwoLineSkipCandles),

		TrailingTopPercentage:      GetFloatFatherOrMomGen(maleBotConfig.TrailingTopPercentage, femaleBotConfig.TrailingTopPercentage),
		TrailingReducePercentage:   GetFloatFatherOrMomGen(maleBotConfig.TrailingReducePercentage, femaleBotConfig.TrailingReducePercentage),
		TrailingIncreasePercentage: GetFloatFatherOrMomGen(maleBotConfig.TrailingIncreasePercentage, femaleBotConfig.TrailingIncreasePercentage),

		StopBuyAfterSellPeriodMinutes: GetIntFatherOrMomGen(maleBotConfig.StopBuyAfterSellPeriodMinutes, femaleBotConfig.StopBuyAfterSellPeriodMinutes),

		AltCoinMarketCandles:       GetIntFatherOrMomGen(maleBotConfig.AltCoinMarketCandles, femaleBotConfig.AltCoinMarketCandles),
		AltCoinMarketMinPercentage: GetFloatFatherOrMomGen(maleBotConfig.AltCoinMarketMinPercentage, femaleBotConfig.AltCoinMarketMinPercentage),

		AltCoinMinBuyMaxSecondPercentage: GetFloatFatherOrMomGen(maleBotConfig.AltCoinMinBuyMaxSecondPercentage, femaleBotConfig.AltCoinMinBuyMaxSecondPercentage),

		WholeDayTotalVolumeCandles:   GetIntFatherOrMomGen(maleBotConfig.WholeDayTotalVolumeCandles, femaleBotConfig.WholeDayTotalVolumeCandles),
		WholeDayTotalVolumeMinVolume: GetFloatFatherOrMomGen(maleBotConfig.WholeDayTotalVolumeMinVolume, femaleBotConfig.WholeDayTotalVolumeMinVolume),

		HalfVolumeFirstCandles:     GetIntFatherOrMomGen(maleBotConfig.HalfVolumeFirstCandles, femaleBotConfig.HalfVolumeFirstCandles),
		HalfVolumeSecondCandles:    GetIntFatherOrMomGen(maleBotConfig.HalfVolumeSecondCandles, femaleBotConfig.HalfVolumeSecondCandles),
		HalfVolumeGrowthPercentage: GetFloatFatherOrMomGen(maleBotConfig.HalfVolumeGrowthPercentage, femaleBotConfig.HalfVolumeGrowthPercentage),

		TrailingActivationPercentage: GetFloatFatherOrMomGen(maleBotConfig.TrailingActivationPercentage, femaleBotConfig.TrailingActivationPercentage),

		FlatLineSearchWindowCandles:          GetIntFatherOrMomGen(maleBotConfig.FlatLineSearchWindowCandles, femaleBotConfig.FlatLineSearchWindowCandles),
		FlatLineSearchWindowsCount:           GetIntFatherOrMomGen(maleBotConfig.FlatLineSearchWindowsCount, femaleBotConfig.FlatLineSearchWindowsCount),
		FlatLineSearchDispersionPercentage:   GetFloatFatherOrMomGen(maleBotConfig.FlatLineSearchDispersionPercentage, femaleBotConfig.FlatLineSearchDispersionPercentage),
		FlatLineSearchOnLinePricesPercentage: GetFloatFatherOrMomGen(maleBotConfig.FlatLineSearchOnLinePricesPercentage, femaleBotConfig.FlatLineSearchOnLinePricesPercentage),
		FlatLineSearchRelativePeriodCandles:  GetIntFatherOrMomGen(maleBotConfig.FlatLineSearchRelativePeriodCandles, femaleBotConfig.FlatLineSearchRelativePeriodCandles),

		TripleGrowthCandles:          GetIntFatherOrMomGen(maleBotConfig.TripleGrowthCandles, femaleBotConfig.TripleGrowthCandles),
		TripleGrowthSecondPercentage: GetFloatFatherOrMomGen(maleBotConfig.TripleGrowthSecondPercentage, femaleBotConfig.TripleGrowthSecondPercentage),

		PastMaxPricePeriod: GetIntFatherOrMomGen(maleBotConfig.PastMaxPricePeriod, femaleBotConfig.PastMaxPricePeriod),

		SmoothGrowthCandles: GetIntFatherOrMomGen(maleBotConfig.SmoothGrowthCandles, femaleBotConfig.SmoothGrowthCandles),
		SmoothGrowthAngle:   GetFloatFatherOrMomGen(maleBotConfig.SmoothGrowthAngle, femaleBotConfig.SmoothGrowthAngle),

		EachVolumeMinValueCandles:     GetIntFatherOrMomGen(maleBotConfig.EachVolumeMinValueCandles, femaleBotConfig.EachVolumeMinValueCandles),
		EachVolumeMinValueMinVolume:   GetFloatFatherOrMomGen(maleBotConfig.EachVolumeMinValueMinVolume, femaleBotConfig.EachVolumeMinValueMinVolume),
		EachVolumeMinValueSkipCandles: GetIntFatherOrMomGen(maleBotConfig.EachVolumeMinValueSkipCandles, femaleBotConfig.EachVolumeMinValueSkipCandles),

		TrailingFixationActivatePercentage:  GetFloatFatherOrMomGen(maleBotConfig.TrailingFixationActivatePercentage, femaleBotConfig.TrailingFixationActivatePercentage),
		TrailingFixationPercentage:          GetFloatFatherOrMomGen(maleBotConfig.TrailingFixationPercentage, femaleBotConfig.TrailingFixationPercentage),
		TrailingSecondaryIncreasePercentage: GetFloatFatherOrMomGen(maleBotConfig.TrailingSecondaryIncreasePercentage, femaleBotConfig.TrailingSecondaryIncreasePercentage),

		AltCoinMaxCandles:    GetIntFatherOrMomGen(maleBotConfig.AltCoinMaxCandles, femaleBotConfig.AltCoinMaxCandles),
		AltCoinMaxPercentage: GetFloatFatherOrMomGen(maleBotConfig.AltCoinMaxPercentage, femaleBotConfig.AltCoinMaxPercentage),
	}

	for i := 0; i < 64; i++ {
		mutateGens(&childBotConfig, GetRandInt(0, 72))
	}

	return GetBotConfigMapInterface(childBotConfig)
}

func GetBotConfigMapInterface(botConfig BotConfig) map[string]interface{} {
	return map[string]interface{}{
		"HighSellPercentage": botConfig.HighSellPercentage,
		"LowSellPercentage":  botConfig.LowSellPercentage,

		"AltCoinMinBuyFirstPeriodMinutes":  botConfig.AltCoinMinBuyFirstPeriodMinutes,
		"AltCoinMinBuyFirstPercentage":     botConfig.AltCoinMinBuyFirstPercentage,
		"AltCoinMinBuySecondPeriodMinutes": botConfig.AltCoinMinBuySecondPeriodMinutes,
		"AltCoinMinBuySecondPercentage":    botConfig.AltCoinMinBuySecondPercentage,

		"BtcMinBuyPeriodMinutes": botConfig.BtcMinBuyPeriodMinutes,
		"BtcMinBuyPercentage":    botConfig.BtcMinBuyPercentage,
		"BtcSellPeriodMinutes":   botConfig.BtcSellPeriodMinutes,
		"BtcSellPercentage":      botConfig.BtcSellPercentage,

		"UnsoldFirstSellDurationMinutes": botConfig.UnsoldFirstSellDurationMinutes,
		"UnsoldFirstSellPercentage":      botConfig.UnsoldFirstSellPercentage,
		"UnsoldFinalSellDurationMinutes": botConfig.UnsoldFinalSellDurationMinutes,

		"AltCoinSuperTrendCandles": botConfig.AltCoinSuperTrendCandles,
		"AltCoinSuperMultiplier":   botConfig.AltCoinSuperMultiplier,

		"BtcSuperTrendCandles":    botConfig.BtcSuperTrendCandles,
		"BtcSuperTrendMultiplier": botConfig.BtcSuperTrendMultiplier,

		"AverageVolumeCandles": botConfig.AverageVolumeCandles,
		"AverageVolumeMinimal": botConfig.AverageVolumeMinimal,

		"AdxDiLen":               botConfig.AdxDiLen,
		"AdxBottomThreshold":     botConfig.AdxBottomThreshold,
		"AdxTopThreshold":        botConfig.AdxTopThreshold,
		"AdxMinGrowthPercentage": botConfig.AdxMinGrowthPercentage,

		"RealBuyTopResetReachRevenue":   botConfig.RealBuyTopResetReachRevenue,
		"RealBuyBottomStopReachRevenue": botConfig.RealBuyBottomStopReachRevenue,
		"FakeBuyReachStopRevenue":       botConfig.FakeBuyReachStopRevenue,

		"CandleBodyCandles":        botConfig.CandleBodyCandles,
		"CandleBodyHeightMinPrice": botConfig.CandleBodyHeightMinPrice,
		"CandleBodyHeightMaxPrice": botConfig.CandleBodyHeightMaxPrice,

		"BtcPriceGrowthCandles":       botConfig.BtcPriceGrowthCandles,
		"BtcPriceGrowthMinPercentage": botConfig.BtcPriceGrowthMinPercentage,
		"BtcPriceGrowthMaxPercentage": botConfig.BtcPriceGrowthMaxPercentage,

		"PriceFallCandles":       botConfig.PriceFallCandles,
		"PriceFallMinPercentage": botConfig.PriceFallMinPercentage,

		"TrailingLowPercentage": botConfig.TrailingLowPercentage,

		"FlatLineCandles":                botConfig.FlatLineCandles,
		"FlatLineSkipCandles":            botConfig.FlatLineSkipCandles,
		"FlatLineDispersionPercentage":   botConfig.FlatLineDispersionPercentage,
		"FlatLineOnLinePricesPercentage": botConfig.FlatLineOnLinePricesPercentage,

		"TwoLineCandles":           botConfig.TwoLineCandles,
		"TwoLineMaxDiffPercentage": botConfig.TwoLineMaxDiffPercentage,
		"TwoLineSkipCandles":       botConfig.TwoLineSkipCandles,

		"TrailingTopPercentage":      botConfig.TrailingTopPercentage,
		"TrailingReducePercentage":   botConfig.TrailingReducePercentage,
		"TrailingIncreasePercentage": botConfig.TrailingIncreasePercentage,

		"StopBuyAfterSellPeriodMinutes": botConfig.StopBuyAfterSellPeriodMinutes,

		"AltCoinMarketCandles":       botConfig.AltCoinMarketCandles,
		"AltCoinMarketMinPercentage": botConfig.AltCoinMarketMinPercentage,

		"AltCoinMinBuyMaxSecondPercentage": botConfig.AltCoinMinBuyMaxSecondPercentage,

		"WholeDayTotalVolumeCandles":   botConfig.WholeDayTotalVolumeCandles,
		"WholeDayTotalVolumeMinVolume": botConfig.WholeDayTotalVolumeMinVolume,

		"HalfVolumeFirstCandles":     botConfig.HalfVolumeFirstCandles,
		"HalfVolumeSecondCandles":    botConfig.HalfVolumeSecondCandles,
		"HalfVolumeGrowthPercentage": botConfig.HalfVolumeGrowthPercentage,

		"TrailingActivationPercentage": botConfig.TrailingActivationPercentage,

		"FlatLineSearchWindowCandles":          botConfig.FlatLineSearchWindowCandles,
		"FlatLineSearchWindowsCount":           botConfig.FlatLineSearchWindowsCount,
		"FlatLineSearchDispersionPercentage":   botConfig.FlatLineSearchDispersionPercentage,
		"FlatLineSearchOnLinePricesPercentage": botConfig.FlatLineSearchOnLinePricesPercentage,
		"FlatLineSearchRelativePeriodCandles":  botConfig.FlatLineSearchRelativePeriodCandles,

		"TripleGrowthCandles":          botConfig.TripleGrowthCandles,
		"TripleGrowthSecondPercentage": botConfig.TripleGrowthSecondPercentage,

		"PastMaxPricePeriod": botConfig.PastMaxPricePeriod,

		"SmoothGrowthCandles": botConfig.SmoothGrowthCandles,
		"SmoothGrowthAngle":   botConfig.SmoothGrowthAngle,

		"EachVolumeMinValueCandles":     botConfig.EachVolumeMinValueCandles,
		"EachVolumeMinValueMinVolume":   botConfig.EachVolumeMinValueMinVolume,
		"EachVolumeMinValueSkipCandles": botConfig.EachVolumeMinValueSkipCandles,

		"TrailingFixationActivatePercentage":  botConfig.TrailingFixationActivatePercentage,
		"TrailingFixationPercentage":          botConfig.TrailingFixationPercentage,
		"TrailingSecondaryIncreasePercentage": botConfig.TrailingSecondaryIncreasePercentage,

		"AltCoinMaxCandles":    botConfig.AltCoinMaxCandles,
		"AltCoinMaxPercentage": botConfig.AltCoinMaxPercentage,

		"TotalRevenue":      botConfig.TotalRevenue,
		"SuccessPercentage": botConfig.SuccessPercentage,
		"PlusRevenue":       botConfig.PlusRevenue,
		"MinusRevenue":      botConfig.MinusRevenue,

		"Selection": botConfig.Selection,
	}
}

func mutateGens(botConfig *BotConfig, randGenNumber int) {
	restrict := GetBotConfigRestrictions()

	mutateGenFloat64(randGenNumber, 0, &(botConfig.HighSellPercentage), restrict.HighSellPercentage)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.LowSellPercentage), restrict.LowSellPercentage)

	mutateGenInt(randGenNumber, 2, &(botConfig.AltCoinMinBuyFirstPeriodMinutes), restrict.AltCoinMinBuyFirstPeriodMinutes)
	mutateGenFloat64(randGenNumber, 3, &(botConfig.AltCoinMinBuyFirstPercentage), restrict.AltCoinMinBuyFirstPercentage)
	mutateGenInt(randGenNumber, 4, &(botConfig.AltCoinMinBuySecondPeriodMinutes), restrict.AltCoinMinBuySecondPeriodMinutes)
	mutateGenFloat64(randGenNumber, 5, &(botConfig.AltCoinMinBuySecondPercentage), restrict.AltCoinMinBuySecondPercentage)

	mutateGenInt(randGenNumber, 6, &(botConfig.BtcMinBuyPeriodMinutes), restrict.BtcMinBuyPeriodMinutes)
	mutateGenFloat64(randGenNumber, 7, &(botConfig.BtcMinBuyPercentage), restrict.BtcMinBuyPercentage)
	mutateGenInt(randGenNumber, 8, &(botConfig.BtcSellPeriodMinutes), restrict.BtcSellPeriodMinutes)
	mutateGenFloat64(randGenNumber, 9, &(botConfig.BtcSellPercentage), restrict.BtcSellPercentage)

	mutateGenInt(randGenNumber, 10, &(botConfig.UnsoldFirstSellDurationMinutes), restrict.UnsoldFirstSellDurationMinutes)
	mutateGenFloat64(randGenNumber, 11, &(botConfig.UnsoldFirstSellPercentage), restrict.UnsoldFirstSellPercentage)
	mutateGenInt(randGenNumber, 12, &(botConfig.UnsoldFinalSellDurationMinutes), restrict.UnsoldFinalSellDurationMinutes)

	mutateGenInt(randGenNumber, 13, &(botConfig.AltCoinSuperTrendCandles), restrict.AltCoinSuperTrendCandles)
	mutateGenFloat64(randGenNumber, 14, &(botConfig.AltCoinSuperMultiplier), restrict.AltCoinSuperMultiplier)

	mutateGenInt(randGenNumber, 15, &(botConfig.BtcSuperTrendCandles), restrict.BtcSuperTrendCandles)
	mutateGenFloat64(randGenNumber, 16, &(botConfig.BtcSuperTrendMultiplier), restrict.BtcSuperTrendMultiplier)

	mutateGenInt(randGenNumber, 17, &(botConfig.AverageVolumeCandles), restrict.AverageVolumeCandles)
	mutateGenFloat64(randGenNumber, 18, &(botConfig.AverageVolumeMinimal), restrict.AverageVolumeMinimal)

	mutateGenInt(randGenNumber, 19, &(botConfig.AdxDiLen), restrict.AdxDiLen)
	mutateGenFloat64(randGenNumber, 20, &(botConfig.AdxBottomThreshold), restrict.AdxBottomThreshold)
	mutateGenFloat64(randGenNumber, 21, &(botConfig.AdxTopThreshold), restrict.AdxTopThreshold)
	mutateGenFloat64(randGenNumber, 22, &(botConfig.AdxMinGrowthPercentage), restrict.AdxMinGrowthPercentage)

	mutateGenFloat64(randGenNumber, 23, &(botConfig.RealBuyTopResetReachRevenue), restrict.RealBuyTopResetReachRevenue)
	mutateGenFloat64(randGenNumber, 24, &(botConfig.RealBuyBottomStopReachRevenue), restrict.RealBuyBottomStopReachRevenue)
	mutateGenFloat64(randGenNumber, 25, &(botConfig.FakeBuyReachStopRevenue), restrict.FakeBuyReachStopRevenue)

	mutateGenInt(randGenNumber, 26, &(botConfig.CandleBodyCandles), restrict.CandleBodyCandles)
	mutateGenFloat64(randGenNumber, 27, &(botConfig.CandleBodyHeightMinPrice), restrict.CandleBodyHeightMinPrice)
	mutateGenFloat64(randGenNumber, 28, &(botConfig.CandleBodyHeightMaxPrice), restrict.CandleBodyHeightMaxPrice)

	mutateGenInt(randGenNumber, 29, &(botConfig.BtcPriceGrowthCandles), restrict.BtcPriceGrowthCandles)
	mutateGenFloat64(randGenNumber, 30, &(botConfig.BtcPriceGrowthMinPercentage), restrict.BtcPriceGrowthMinPercentage)
	mutateGenFloat64(randGenNumber, 31, &(botConfig.CandleBodyHeightMaxPrice), restrict.BtcPriceGrowthMaxPercentage)

	mutateGenInt(randGenNumber, 32, &(botConfig.PriceFallCandles), restrict.PriceFallCandles)
	mutateGenFloat64(randGenNumber, 33, &(botConfig.PriceFallMinPercentage), restrict.PriceFallMinPercentage)

	mutateGenFloat64(randGenNumber, 34, &(botConfig.TrailingLowPercentage), restrict.TrailingLowPercentage)

	mutateGenInt(randGenNumber, 35, &(botConfig.FlatLineCandles), restrict.FlatLineCandles)
	mutateGenInt(randGenNumber, 36, &(botConfig.FlatLineSkipCandles), restrict.FlatLineSkipCandles)
	mutateGenFloat64(randGenNumber, 37, &(botConfig.FlatLineDispersionPercentage), restrict.FlatLineDispersionPercentage)
	mutateGenFloat64(randGenNumber, 38, &(botConfig.FlatLineOnLinePricesPercentage), restrict.FlatLineOnLinePricesPercentage)

	mutateGenInt(randGenNumber, 39, &(botConfig.TwoLineCandles), restrict.TwoLineCandles)
	mutateGenFloat64(randGenNumber, 40, &(botConfig.TwoLineMaxDiffPercentage), restrict.TwoLineMaxDiffPercentage)
	mutateGenInt(randGenNumber, 41, &(botConfig.TwoLineSkipCandles), restrict.TwoLineSkipCandles)

	mutateGenFloat64(randGenNumber, 42, &(botConfig.TrailingTopPercentage), restrict.TrailingTopPercentage)
	mutateGenFloat64(randGenNumber, 43, &(botConfig.TrailingReducePercentage), restrict.TrailingReducePercentage)
	mutateGenFloat64(randGenNumber, 44, &(botConfig.TrailingIncreasePercentage), restrict.TrailingIncreasePercentage)

	mutateGenInt(randGenNumber, 45, &(botConfig.StopBuyAfterSellPeriodMinutes), restrict.StopBuyAfterSellPeriodMinutes)

	mutateGenInt(randGenNumber, 46, &(botConfig.AltCoinMarketCandles), restrict.AltCoinMarketCandles)
	mutateGenFloat64(randGenNumber, 47, &(botConfig.AltCoinMarketMinPercentage), restrict.AltCoinMarketMinPercentage)

	mutateGenFloat64(randGenNumber, 48, &(botConfig.AltCoinMinBuyMaxSecondPercentage), restrict.AltCoinMinBuyMaxSecondPercentage)

	mutateGenInt(randGenNumber, 49, &(botConfig.WholeDayTotalVolumeCandles), restrict.WholeDayTotalVolumeCandles)
	mutateGenFloat64(randGenNumber, 50, &(botConfig.WholeDayTotalVolumeMinVolume), restrict.WholeDayTotalVolumeMinVolume)

	mutateGenInt(randGenNumber, 51, &(botConfig.HalfVolumeFirstCandles), restrict.HalfVolumeFirstCandles)
	mutateGenInt(randGenNumber, 52, &(botConfig.HalfVolumeSecondCandles), restrict.HalfVolumeSecondCandles)
	mutateGenFloat64(randGenNumber, 53, &(botConfig.HalfVolumeGrowthPercentage), restrict.HalfVolumeGrowthPercentage)

	mutateGenFloat64(randGenNumber, 54, &(botConfig.TrailingActivationPercentage), restrict.TrailingActivationPercentage)

	mutateGenInt(randGenNumber, 55, &(botConfig.FlatLineSearchWindowCandles), restrict.FlatLineSearchWindowCandles)
	mutateGenInt(randGenNumber, 56, &(botConfig.FlatLineSearchWindowsCount), restrict.FlatLineSearchWindowsCount)
	mutateGenFloat64(randGenNumber, 57, &(botConfig.FlatLineSearchDispersionPercentage), restrict.FlatLineSearchDispersionPercentage)
	mutateGenFloat64(randGenNumber, 58, &(botConfig.FlatLineSearchOnLinePricesPercentage), restrict.FlatLineSearchOnLinePricesPercentage)
	mutateGenInt(randGenNumber, 59, &(botConfig.FlatLineSearchRelativePeriodCandles), restrict.FlatLineSearchRelativePeriodCandles)

	mutateGenInt(randGenNumber, 60, &(botConfig.TripleGrowthCandles), restrict.TripleGrowthCandles)
	mutateGenFloat64(randGenNumber, 61, &(botConfig.TripleGrowthSecondPercentage), restrict.TripleGrowthSecondPercentage)

	mutateGenInt(randGenNumber, 62, &(botConfig.PastMaxPricePeriod), restrict.PastMaxPricePeriod)

	mutateGenInt(randGenNumber, 63, &(botConfig.SmoothGrowthCandles), restrict.SmoothGrowthCandles)
	mutateGenFloat64(randGenNumber, 64, &(botConfig.SmoothGrowthAngle), restrict.SmoothGrowthAngle)

	mutateGenInt(randGenNumber, 65, &(botConfig.EachVolumeMinValueCandles), restrict.EachVolumeMinValueCandles)
	mutateGenFloat64(randGenNumber, 66, &(botConfig.EachVolumeMinValueMinVolume), restrict.EachVolumeMinValueMinVolume)
	mutateGenInt(randGenNumber, 67, &(botConfig.EachVolumeMinValueSkipCandles), restrict.EachVolumeMinValueSkipCandles)

	mutateGenFloat64(randGenNumber, 68, &(botConfig.TrailingFixationActivatePercentage), restrict.TrailingFixationActivatePercentage)
	mutateGenFloat64(randGenNumber, 69, &(botConfig.TrailingFixationPercentage), restrict.TrailingFixationPercentage)
	mutateGenFloat64(randGenNumber, 70, &(botConfig.TrailingSecondaryIncreasePercentage), restrict.TrailingSecondaryIncreasePercentage)

	mutateGenInt(randGenNumber, 71, &(botConfig.AltCoinMaxCandles), restrict.AltCoinMaxCandles)
	mutateGenFloat64(randGenNumber, 72, &(botConfig.AltCoinMaxPercentage), restrict.AltCoinMaxPercentage)
}

func mutateGenFloat64(randGenNumber, genNumber int, genValue *float64, restrictMinMax MinMaxFloat64) {
	if randGenNumber == genNumber {
		*genValue = MutateLittleFloat64(*genValue, restrictMinMax)
	}
}

func mutateGenInt(randGenNumber, genNumber int, genValue *int, restrictMinMax MinMaxInt) {
	if randGenNumber == genNumber {
		*genValue = MutateLittleInt(*genValue, restrictMinMax)
	}
}

func shuffleBots(bots *dataframe.DataFrame) *dataframe.DataFrame {
	shuffledBots := InitBotsDataFrame()
	shuffleNumbers := rand.Perm(bots.NRows())
	for i := 0; i < bots.NRows()-1; i++ {
		shuffledBots.Append(nil, bots.Row(shuffleNumbers[i], false))
	}
	return shuffledBots
}

func GetFloatFatherOrMomGen(maleGen, femaleGen float64) float64 {
	if GetRandInt(0, 1) == 1 {
		return maleGen
	}

	return femaleGen
}

func GetIntFatherOrMomGen(maleGen, femaleGen int) int {
	if GetRandInt(0, 1) == 1 {
		return maleGen
	}

	return femaleGen
}

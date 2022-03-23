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

		dataframe.NewSeriesFloat64("TotalRevenue", nil),
		dataframe.NewSeriesFloat64("SuccessPercentage", nil),
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

func SetBotTotalRevenue(bots *dataframe.DataFrame, botNumber int, revenue, successPercentage float64) {
	bots.UpdateRow(botNumber, nil, map[string]interface{}{
		"TotalRevenue":      revenue,
		"SuccessPercentage": successPercentage,
	})
}

func SortBestBots(bots *dataframe.DataFrame) *dataframe.DataFrame {
	sks := []dataframe.SortKey{
		{
			Key:  "TotalRevenue",
			Desc: true,
		},
		{
			Key:  "SuccessPercentage",
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

		"TotalRevenue":      bot["TotalRevenue"],
		"SuccessPercentage": bot["SuccessPercentage"],
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
	}

	for i := 0; i < 12; i++ {
		mutateGens(&childBotConfig, GetRandInt(0, 31))
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

		"TotalRevenue":      botConfig.TotalRevenue,
		"SuccessPercentage": botConfig.SuccessPercentage,
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

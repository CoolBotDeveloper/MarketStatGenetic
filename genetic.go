package main

import (
	"context"
	"github.com/rocketlaunchr/dataframe-go"
	"math/rand"
)

const BEST_BOTS_COUNT = 5
const BOTS_COUNT = 10
const GENERATION_COUNT = 100

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

		dataframe.NewSeriesFloat64("TotalRevenue", nil),
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

func SetBotTotalRevenue(bots *dataframe.DataFrame, botNumber int, revenue float64) {
	bots.UpdateRow(botNumber, nil, map[string]interface{}{
		"TotalRevenue": revenue,
	})
}

func SortBestBots(bots *dataframe.DataFrame) *dataframe.DataFrame {
	sks := []dataframe.SortKey{
		{
			Key:  "TotalRevenue",
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
	for {
		botNumber, bot, _ := iterator()
		if botNumber == nil || numberOfBots < (*botNumber+1) {
			break
		}
		botsDataFrame.Append(nil, map[string]interface{}{
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

			"AdxDiLen":           bot["AdxDiLen"],
			"AdxBottomThreshold": bot["AdxBottomThreshold"],
			"AdxTopThreshold":    bot["AdxTopThreshold"],

			"TotalRevenue": bot["TotalRevenue"],
		})
	}
	return botsDataFrame
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

		AdxDiLen:           GetIntFatherOrMomGen(maleBotConfig.AdxDiLen, femaleBotConfig.AdxDiLen),
		AdxBottomThreshold: GetFloatFatherOrMomGen(maleBotConfig.AdxBottomThreshold, femaleBotConfig.AdxBottomThreshold),
		AdxTopThreshold:    GetFloatFatherOrMomGen(maleBotConfig.AdxTopThreshold, femaleBotConfig.AdxTopThreshold),
	}

	for i := 0; i < 7; i++ {
		mutateGens(&childBotConfig, GetRandInt(0, 21))
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

		"AdxDiLen":           botConfig.AdxDiLen,
		"AdxBottomThreshold": botConfig.AdxBottomThreshold,
		"AdxTopThreshold":    botConfig.AdxTopThreshold,

		"TotalRevenue": 0.0,
	}
}

func mutateGens(botConfig *BotConfig, randGenNumber int) {
	restrict := GetBotConfigRestrictions()

	mutateGenFloat64(randGenNumber, 0, &(botConfig.HighSellPercentage), restrict.HighSellPercentage)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.LowSellPercentage), restrict.LowSellPercentage)

	mutateGenInt(randGenNumber, 1, &(botConfig.AltCoinMinBuyFirstPeriodMinutes), restrict.AltCoinMinBuyFirstPeriodMinutes)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.AltCoinMinBuyFirstPercentage), restrict.AltCoinMinBuyFirstPercentage)
	mutateGenInt(randGenNumber, 1, &(botConfig.AltCoinMinBuySecondPeriodMinutes), restrict.AltCoinMinBuySecondPeriodMinutes)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.AltCoinMinBuySecondPercentage), restrict.AltCoinMinBuySecondPercentage)

	mutateGenInt(randGenNumber, 1, &(botConfig.BtcMinBuyPeriodMinutes), restrict.BtcMinBuyPeriodMinutes)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.BtcMinBuyPercentage), restrict.BtcMinBuyPercentage)
	mutateGenInt(randGenNumber, 1, &(botConfig.BtcSellPeriodMinutes), restrict.BtcSellPeriodMinutes)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.BtcSellPercentage), restrict.BtcSellPercentage)

	mutateGenInt(randGenNumber, 1, &(botConfig.UnsoldFirstSellDurationMinutes), restrict.UnsoldFirstSellDurationMinutes)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.UnsoldFirstSellPercentage), restrict.UnsoldFirstSellPercentage)
	mutateGenInt(randGenNumber, 1, &(botConfig.UnsoldFinalSellDurationMinutes), restrict.UnsoldFinalSellDurationMinutes)

	mutateGenInt(randGenNumber, 1, &(botConfig.AltCoinSuperTrendCandles), restrict.AltCoinSuperTrendCandles)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.AltCoinSuperMultiplier), restrict.AltCoinSuperMultiplier)

	mutateGenInt(randGenNumber, 1, &(botConfig.BtcSuperTrendCandles), restrict.BtcSuperTrendCandles)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.BtcSuperTrendMultiplier), restrict.BtcSuperTrendMultiplier)

	mutateGenInt(randGenNumber, 1, &(botConfig.AverageVolumeCandles), restrict.AverageVolumeCandles)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.AverageVolumeMinimal), restrict.AverageVolumeMinimal)

	mutateGenInt(randGenNumber, 1, &(botConfig.AdxDiLen), restrict.AdxDiLen)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.AdxBottomThreshold), restrict.AdxBottomThreshold)
	mutateGenFloat64(randGenNumber, 1, &(botConfig.AdxTopThreshold), restrict.AdxTopThreshold)
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

package main

import (
	"github.com/MaxHalford/eaopt"
	"math/rand"
)

const BITCOIN_SYMBOL = "BTCUSDT"
const TOTAL_MONEY_AMOUNT = 100.0
const COMMISSION = 0.15
const MUTATE_RATE = 5

type BotConfig struct {
	HighSellPercentage float64
	LowSellPercentage  float64

	AltCoinMinBuyFirstPeriodMinutes  int
	AltCoinMinBuyFirstPercentage     float64
	AltCoinMinBuySecondPeriodMinutes int
	AltCoinMinBuySecondPercentage    float64

	BtcMinBuyPeriodMinutes int
	BtcMinBuyPercentage    float64
	BtcSellPeriodMinutes   int
	BtcSellPercentage      float64

	UnsoldFirstSellDurationMinutes int
	UnsoldFirstSellPercentage      float64
	UnsoldFinalSellDurationMinutes int

	AltCoinSuperTrendCandles int
	AltCoinSuperMultiplier   float64

	BtcSuperTrendCandles    int
	BtcSuperTrendMultiplier float64

	AverageVolumeCandles int
	AverageVolumeMinimal float64

	AdxDiLen           int
	AdxBottomThreshold float64
	AdxTopThreshold    float64

	Restrictions BotConfigRestriction
}

type BotConfigSlice struct {
	HighSellPercentage *eaopt.Float64Slice
	LowSellPercentage  *eaopt.Float64Slice

	AltCoinMinBuyFirstPeriodMinutes  *eaopt.IntSlice
	AltCoinMinBuyFirstPercentage     *eaopt.Float64Slice
	AltCoinMinBuySecondPeriodMinutes *eaopt.IntSlice
	AltCoinMinBuySecondPercentage    *eaopt.Float64Slice

	BtcMinBuyPeriodMinutes *eaopt.IntSlice
	BtcMinBuyPercentage    *eaopt.Float64Slice
	BtcSellPeriodMinutes   *eaopt.IntSlice
	BtcSellPercentage      *eaopt.Float64Slice

	UnsoldFirstSellDurationMinutes *eaopt.IntSlice
	UnsoldFirstSellPercentage      *eaopt.Float64Slice
	UnsoldFinalSellDurationMinutes *eaopt.IntSlice

	AltCoinSuperTrendCandles *eaopt.IntSlice
	AltCoinSuperMultiplier   *eaopt.Float64Slice

	BtcSuperTrendCandles    *eaopt.IntSlice
	BtcSuperTrendMultiplier *eaopt.Float64Slice

	AverageVolumeCandles *eaopt.IntSlice
	AverageVolumeMinimal *eaopt.Float64Slice

	AdxDiLen           *eaopt.IntSlice
	AdxBottomThreshold *eaopt.Float64Slice
	AdxTopThreshold    *eaopt.Float64Slice

	Restrictions BotConfigRestriction
}

func (bot BotConfigSlice) GetBotConfig() BotConfig {
	return BotConfig{
		HighSellPercentage: bot.HighSellPercentage.At(0).(float64),
		LowSellPercentage:  bot.LowSellPercentage.At(0).(float64),

		AltCoinMinBuyFirstPeriodMinutes:  bot.AltCoinMinBuyFirstPeriodMinutes.At(0).(int),
		AltCoinMinBuyFirstPercentage:     bot.AltCoinMinBuyFirstPercentage.At(0).(float64),
		AltCoinMinBuySecondPeriodMinutes: bot.AltCoinMinBuySecondPeriodMinutes.At(0).(int),
		AltCoinMinBuySecondPercentage:    bot.AltCoinMinBuySecondPercentage.At(0).(float64),

		BtcMinBuyPeriodMinutes: bot.BtcMinBuyPeriodMinutes.At(0).(int),
		BtcMinBuyPercentage:    bot.BtcMinBuyPercentage.At(0).(float64),
		BtcSellPeriodMinutes:   bot.BtcSellPeriodMinutes.At(0).(int),
		BtcSellPercentage:      bot.BtcSellPercentage.At(0).(float64),

		UnsoldFirstSellDurationMinutes: bot.UnsoldFirstSellDurationMinutes.At(0).(int),
		UnsoldFirstSellPercentage:      bot.UnsoldFirstSellPercentage.At(0).(float64),
		UnsoldFinalSellDurationMinutes: bot.UnsoldFinalSellDurationMinutes.At(0).(int),

		AltCoinSuperTrendCandles: bot.AltCoinSuperTrendCandles.At(0).(int),
		AltCoinSuperMultiplier:   bot.AltCoinSuperMultiplier.At(0).(float64),

		BtcSuperTrendCandles:    bot.BtcSuperTrendCandles.At(0).(int),
		BtcSuperTrendMultiplier: bot.BtcSuperTrendMultiplier.At(0).(float64),

		AverageVolumeCandles: bot.AverageVolumeCandles.At(0).(int),
		AverageVolumeMinimal: bot.AverageVolumeMinimal.At(0).(float64),

		AdxDiLen:           bot.AdxDiLen.At(0).(int),
		AdxBottomThreshold: bot.AdxBottomThreshold.At(0).(float64),
		AdxTopThreshold:    bot.AdxTopThreshold.At(0).(float64),

		Restrictions: bot.Restrictions,
	}
}

func NewBot() BotConfig {
	return BotConfig{}
}

//func BotFactory(rng *rand.Rand) eaopt.Genome {
//	return InitBotConfig()
//}

func BotSliceFactory(rng *rand.Rand) eaopt.Genome {
	return InitBotConfigSlice()
}

func (bot BotConfigSlice) Evaluate() (float64, error) {
	return 100 - Fitness(bot.GetBotConfig()), nil
}

func (bot BotConfigSlice) Mutate(rng *rand.Rand) {
	botConfig := bot.GetBotConfig()
	restrictions := bot.Restrictions

	*bot.HighSellPercentage = eaopt.Float64Slice{MutateLittleFloat64(botConfig.HighSellPercentage, restrictions.HighSellPercentage)}
	*bot.LowSellPercentage = eaopt.Float64Slice{MutateLittleFloat64(botConfig.LowSellPercentage, restrictions.LowSellPercentage)}

	*bot.AltCoinMinBuyFirstPeriodMinutes = eaopt.IntSlice{MutateLittleInt(botConfig.AltCoinMinBuyFirstPeriodMinutes, restrictions.AltCoinMinBuyFirstPeriodMinutes)}
	*bot.AltCoinMinBuyFirstPercentage = eaopt.Float64Slice{MutateLittleFloat64(botConfig.AltCoinMinBuyFirstPercentage, restrictions.AltCoinMinBuyFirstPercentage)}
	*bot.AltCoinMinBuySecondPeriodMinutes = eaopt.IntSlice{MutateLittleInt(botConfig.AltCoinMinBuySecondPeriodMinutes, restrictions.AltCoinMinBuySecondPeriodMinutes)}
	*bot.AltCoinMinBuySecondPercentage = eaopt.Float64Slice{MutateLittleFloat64(botConfig.AltCoinMinBuySecondPercentage, restrictions.AltCoinMinBuySecondPercentage)}

	*bot.BtcMinBuyPeriodMinutes = eaopt.IntSlice{MutateLittleInt(botConfig.BtcMinBuyPeriodMinutes, restrictions.BtcMinBuyPeriodMinutes)}
	*bot.BtcMinBuyPercentage = eaopt.Float64Slice{MutateLittleFloat64(botConfig.BtcMinBuyPercentage, restrictions.BtcMinBuyPercentage)}
	*bot.BtcSellPeriodMinutes = eaopt.IntSlice{MutateLittleInt(botConfig.BtcSellPeriodMinutes, restrictions.BtcSellPeriodMinutes)}
	*bot.BtcSellPercentage = eaopt.Float64Slice{MutateLittleFloat64(botConfig.BtcSellPercentage, restrictions.BtcSellPercentage)}

	*bot.UnsoldFirstSellDurationMinutes = eaopt.IntSlice{MutateLittleInt(botConfig.UnsoldFirstSellDurationMinutes, restrictions.UnsoldFirstSellDurationMinutes)}
	*bot.UnsoldFirstSellPercentage = eaopt.Float64Slice{MutateLittleFloat64(botConfig.UnsoldFirstSellPercentage, restrictions.UnsoldFirstSellPercentage)}
	*bot.UnsoldFinalSellDurationMinutes = eaopt.IntSlice{MutateLittleInt(botConfig.UnsoldFinalSellDurationMinutes, restrictions.UnsoldFinalSellDurationMinutes)}

	*bot.AltCoinSuperTrendCandles = eaopt.IntSlice{MutateLittleInt(botConfig.AltCoinSuperTrendCandles, restrictions.AltCoinSuperTrendCandles)}
	*bot.AltCoinSuperMultiplier = eaopt.Float64Slice{MutateLittleFloat64(botConfig.AltCoinSuperMultiplier, restrictions.AltCoinSuperMultiplier)}

	*bot.BtcSuperTrendCandles = eaopt.IntSlice{MutateLittleInt(botConfig.BtcSuperTrendCandles, restrictions.BtcSuperTrendCandles)}
	*bot.BtcSuperTrendMultiplier = eaopt.Float64Slice{MutateLittleFloat64(botConfig.BtcSuperTrendMultiplier, restrictions.BtcSuperTrendMultiplier)}

	*bot.AverageVolumeCandles = eaopt.IntSlice{MutateLittleInt(botConfig.AverageVolumeCandles, restrictions.AverageVolumeCandles)}
	*bot.AverageVolumeMinimal = eaopt.Float64Slice{MutateLittleFloat64(botConfig.AverageVolumeMinimal, restrictions.AverageVolumeMinimal)}

	*bot.AdxDiLen = eaopt.IntSlice{MutateLittleInt(botConfig.AdxDiLen, restrictions.AdxDiLen)}
	*bot.AdxBottomThreshold = eaopt.Float64Slice{MutateLittleFloat64(botConfig.AdxBottomThreshold, restrictions.AdxBottomThreshold)}
	*bot.AdxTopThreshold = eaopt.Float64Slice{MutateLittleFloat64(botConfig.AdxTopThreshold, restrictions.AdxTopThreshold)}
}

func (bot BotConfigSlice) Crossover(female eaopt.Genome, rng *rand.Rand) {
	femaleBotConfig := female.(BotConfigSlice)

	CrossGenFloat64Slice(bot.HighSellPercentage, femaleBotConfig.HighSellPercentage)
	CrossGenFloat64Slice(bot.LowSellPercentage, femaleBotConfig.LowSellPercentage)

	CrossGenIntSlice(bot.AltCoinMinBuyFirstPeriodMinutes, femaleBotConfig.AltCoinMinBuyFirstPeriodMinutes)
	CrossGenFloat64Slice(bot.AltCoinMinBuyFirstPercentage, femaleBotConfig.AltCoinMinBuyFirstPercentage)
	CrossGenIntSlice(bot.AltCoinMinBuySecondPeriodMinutes, femaleBotConfig.AltCoinMinBuySecondPeriodMinutes)
	CrossGenFloat64Slice(bot.AltCoinMinBuySecondPercentage, femaleBotConfig.AltCoinMinBuySecondPercentage)

	CrossGenIntSlice(bot.BtcMinBuyPeriodMinutes, femaleBotConfig.BtcMinBuyPeriodMinutes)
	CrossGenFloat64Slice(bot.BtcMinBuyPercentage, femaleBotConfig.BtcMinBuyPercentage)
	CrossGenIntSlice(bot.BtcSellPeriodMinutes, femaleBotConfig.BtcSellPeriodMinutes)
	CrossGenFloat64Slice(bot.BtcSellPercentage, femaleBotConfig.BtcSellPercentage)

	CrossGenIntSlice(bot.UnsoldFirstSellDurationMinutes, femaleBotConfig.UnsoldFirstSellDurationMinutes)
	CrossGenFloat64Slice(bot.UnsoldFirstSellPercentage, femaleBotConfig.UnsoldFirstSellPercentage)
	CrossGenIntSlice(bot.UnsoldFinalSellDurationMinutes, femaleBotConfig.UnsoldFinalSellDurationMinutes)

	CrossGenIntSlice(bot.AltCoinSuperTrendCandles, femaleBotConfig.AltCoinSuperTrendCandles)
	CrossGenFloat64Slice(bot.AltCoinSuperMultiplier, femaleBotConfig.AltCoinSuperMultiplier)

	CrossGenIntSlice(bot.BtcSuperTrendCandles, femaleBotConfig.BtcSuperTrendCandles)
	CrossGenFloat64Slice(bot.BtcSuperTrendMultiplier, femaleBotConfig.BtcSuperTrendMultiplier)

	CrossGenIntSlice(bot.AverageVolumeCandles, femaleBotConfig.AverageVolumeCandles)
	CrossGenFloat64Slice(bot.AverageVolumeMinimal, femaleBotConfig.AverageVolumeMinimal)

	CrossGenIntSlice(bot.AdxDiLen, femaleBotConfig.AdxDiLen)
	CrossGenFloat64Slice(bot.AdxBottomThreshold, femaleBotConfig.AdxBottomThreshold)
	CrossGenFloat64Slice(bot.AdxTopThreshold, femaleBotConfig.AdxTopThreshold)
}

func (bot BotConfigSlice) Clone() eaopt.Genome {
	Y := bot
	return Y
}

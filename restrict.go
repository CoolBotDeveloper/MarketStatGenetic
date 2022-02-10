package main

import (
	"math/rand"
	"time"
)

type BotConfigRestriction struct {
	HighSellPercentage MinMaxFloat64
	LowSellPercentage  MinMaxFloat64

	AltCoinMinBuyFirstPeriodMinutes  MinMaxInt
	AltCoinMinBuyFirstPercentage     MinMaxFloat64
	AltCoinMinBuySecondPeriodMinutes MinMaxInt
	AltCoinMinBuySecondPercentage    MinMaxFloat64

	BtcMinBuyPeriodMinutes MinMaxInt
	BtcMinBuyPercentage    MinMaxFloat64
	BtcSellPeriodMinutes   MinMaxInt
	BtcSellPercentage      MinMaxFloat64

	UnsoldFirstSellDurationMinutes MinMaxInt
	UnsoldFirstSellPercentage      MinMaxFloat64
	UnsoldFinalSellDurationMinutes MinMaxInt

	AltCoinSuperTrendCandles MinMaxInt
	AltCoinSuperMultiplier   MinMaxFloat64

	BtcSuperTrendCandles    MinMaxInt
	BtcSuperTrendMultiplier MinMaxFloat64

	AverageVolumeCandles MinMaxInt
	AverageVolumeMinimal MinMaxFloat64

	AdxDiLen           MinMaxInt
	AdxBottomThreshold MinMaxFloat64
	AdxTopThreshold    MinMaxFloat64
}

type MinMaxInt struct {
	min int
	max int
}

type MinMaxFloat64 struct {
	min float64
	max float64
}

func GetBotConfigRestrictions() BotConfigRestriction {
	return BotConfigRestriction{
		HighSellPercentage: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},
		LowSellPercentage: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},

		// -----------------------------------------------------
		AltCoinMinBuyFirstPeriodMinutes: MinMaxInt{
			min: 0,
			max: 0,
		},
		AltCoinMinBuyFirstPercentage: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},
		AltCoinMinBuySecondPeriodMinutes: MinMaxInt{
			min: 0,
			max: 0,
		},
		AltCoinMinBuySecondPercentage: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},

		// -----------------------------------------------------
		BtcMinBuyPeriodMinutes: MinMaxInt{
			min: 0,
			max: 0,
		},
		BtcMinBuyPercentage: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},
		BtcSellPeriodMinutes: MinMaxInt{
			min: 0,
			max: 0,
		},
		BtcSellPercentage: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},

		// -----------------------------------------------------
		UnsoldFirstSellDurationMinutes: MinMaxInt{
			min: 0,
			max: 0,
		},
		UnsoldFirstSellPercentage: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},
		UnsoldFinalSellDurationMinutes: MinMaxInt{
			min: 0,
			max: 0,
		},

		// -----------------------------------------------------
		AltCoinSuperTrendCandles: MinMaxInt{
			min: 0,
			max: 0,
		},
		AltCoinSuperMultiplier: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},

		// -----------------------------------------------------
		BtcSuperTrendCandles: MinMaxInt{
			min: 0,
			max: 0,
		},
		BtcSuperTrendMultiplier: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},

		// -----------------------------------------------------
		AverageVolumeCandles: MinMaxInt{
			min: 0,
			max: 0,
		},
		AverageVolumeMinimal: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},

		// -----------------------------------------------------
		AdxDiLen: MinMaxInt{
			min: 0,
			max: 0,
		},
		AdxBottomThreshold: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},
		AdxTopThreshold: MinMaxFloat64{
			min: 0.0,
			max: 0.0,
		},
	}
}

func GetRandIntConfig(minMax MinMaxInt) int {
	return GetRandInt(minMax.min, minMax.max)
}

func GetRandFloat64Config(minMax MinMaxFloat64) float64 {
	return GetRandFloat64(minMax.min, minMax.max)
}

func GetRandInt(lower int, upper int) int {
	rand.Seed(time.Now().UnixNano())
	return lower + rand.Intn(upper-lower+1)
}

func GetRandFloat64(lower float64, upper float64) float64 {
	rand.Seed(time.Now().UnixNano())
	return lower + rand.Float64()*(upper-lower)
}

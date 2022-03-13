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

	RealBuyTopResetReachRevenue   MinMaxFloat64
	RealBuyBottomStopReachRevenue MinMaxFloat64
	FakeBuyReachStopRevenue       MinMaxFloat64

	CandleBodyCandles        MinMaxInt
	CandleBodyHeightMinPrice MinMaxFloat64
	CandleBodyHeightMaxPrice MinMaxFloat64
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
			min: 0.2,
			max: 5.0,
		},
		LowSellPercentage: MinMaxFloat64{
			min: 0.5,
			max: 5.0,
		},

		// -----------------------------------------------------
		AltCoinMinBuyFirstPeriodMinutes: MinMaxInt{
			min: 60,
			max: 60 * 12,
		},
		AltCoinMinBuyFirstPercentage: MinMaxFloat64{
			min: 0.15,
			max: 10,
		},
		AltCoinMinBuySecondPeriodMinutes: MinMaxInt{
			min: 1,
			max: 20,
		},
		AltCoinMinBuySecondPercentage: MinMaxFloat64{
			min: 0.15,
			max: 10,
		},

		// -----------------------------------------------------
		BtcMinBuyPeriodMinutes: MinMaxInt{
			min: 60,
			max: 60 * 12,
		},
		BtcMinBuyPercentage: MinMaxFloat64{
			min: -0.5,
			max: 5,
		},
		BtcSellPeriodMinutes: MinMaxInt{
			min: 60,
			max: 60 * 15,
		},
		BtcSellPercentage: MinMaxFloat64{
			min: -0.6,
			max: 1.5,
		},

		// -----------------------------------------------------
		UnsoldFirstSellDurationMinutes: MinMaxInt{
			min: 1,
			max: 10,
		},
		UnsoldFirstSellPercentage: MinMaxFloat64{
			min: 0.2,
			max: 3,
		},
		UnsoldFinalSellDurationMinutes: MinMaxInt{
			min: 11,
			max: 45,
		},

		// -----------------------------------------------------
		AltCoinSuperTrendCandles: MinMaxInt{
			min: 1,
			max: 20,
		},
		AltCoinSuperMultiplier: MinMaxFloat64{
			min: 0.5,
			max: 50,
		},

		// -----------------------------------------------------
		BtcSuperTrendCandles: MinMaxInt{
			min: 5,
			max: 15,
		},
		BtcSuperTrendMultiplier: MinMaxFloat64{
			min: 10,
			max: 35,
		},

		// -----------------------------------------------------
		AverageVolumeCandles: MinMaxInt{
			min: 5,
			max: 50,
		},
		AverageVolumeMinimal: MinMaxFloat64{
			min: 35000,
			max: 100000,
		},

		// -----------------------------------------------------
		AdxDiLen: MinMaxInt{
			min: 5,
			max: 25,
		},
		AdxBottomThreshold: MinMaxFloat64{
			min: 20,
			max: 40,
		},
		AdxTopThreshold: MinMaxFloat64{
			min: 40,
			max: 90,
		},

		// -----------------------------------------------------
		RealBuyTopResetReachRevenue: MinMaxFloat64{
			min: 0.5,
			max: 5,
		},
		RealBuyBottomStopReachRevenue: MinMaxFloat64{
			min: -2,
			max: 3,
		},
		FakeBuyReachStopRevenue: MinMaxFloat64{
			min: -0.5,
			max: 5,
		},

		// -----------------------------------------------------
		CandleBodyCandles: MinMaxInt{
			min: 1,
			max: 100,
		},
		CandleBodyHeightMinPrice: MinMaxFloat64{ // В процентах
			min: 0.01,
			max: 50,
		},
		CandleBodyHeightMaxPrice: MinMaxFloat64{ // В процентах
			min: 50,
			max: 200,
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

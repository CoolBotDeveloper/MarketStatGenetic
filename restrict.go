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

	AdxDiLen               MinMaxInt
	AdxBottomThreshold     MinMaxFloat64
	AdxTopThreshold        MinMaxFloat64
	AdxMinGrowthPercentage MinMaxFloat64

	RealBuyTopResetReachRevenue   MinMaxFloat64
	RealBuyBottomStopReachRevenue MinMaxFloat64
	FakeBuyReachStopRevenue       MinMaxFloat64

	CandleBodyCandles        MinMaxInt
	CandleBodyHeightMinPrice MinMaxFloat64
	CandleBodyHeightMaxPrice MinMaxFloat64

	BtcPriceGrowthCandles       MinMaxInt
	BtcPriceGrowthMinPercentage MinMaxFloat64
	BtcPriceGrowthMaxPercentage MinMaxFloat64

	PriceFallCandles       MinMaxInt
	PriceFallMinPercentage MinMaxFloat64

	TrailingLowPercentage MinMaxFloat64

	FlatLineCandles                MinMaxInt
	FlatLineSkipCandles            MinMaxInt
	FlatLineDispersionPercentage   MinMaxFloat64
	FlatLineOnLinePricesPercentage MinMaxFloat64

	TwoLineCandles           MinMaxInt
	TwoLineMaxDiffPercentage MinMaxFloat64
	TwoLineSkipCandles       MinMaxInt

	TrailingTopPercentage      MinMaxFloat64
	TrailingReducePercentage   MinMaxFloat64
	TrailingIncreasePercentage MinMaxFloat64

	StopBuyAfterSellPeriodMinutes MinMaxInt

	AltCoinMarketCandles       MinMaxInt
	AltCoinMarketMinPercentage MinMaxFloat64
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
			min: 0.3,
			max: 5.0,
		},
		LowSellPercentage: MinMaxFloat64{
			min: 1.5,
			max: 5,
		},

		// -----------------------------------------------------
		AltCoinMinBuyFirstPeriodMinutes: MinMaxInt{
			min: 50,
			max: 500,
		},
		AltCoinMinBuyFirstPercentage: MinMaxFloat64{
			min: 1,
			max: 15,
		},
		AltCoinMinBuySecondPeriodMinutes: MinMaxInt{
			min: 5,
			max: 20,
		},
		AltCoinMinBuySecondPercentage: MinMaxFloat64{
			min: 2,
			max: 5,
		},

		// -----------------------------------------------------
		BtcMinBuyPeriodMinutes: MinMaxInt{
			min: 30,
			max: 60 * 12,
		},
		BtcMinBuyPercentage: MinMaxFloat64{
			min: -1,
			max: 5,
		},
		BtcSellPeriodMinutes: MinMaxInt{
			min: 30,
			max: 60 * 15,
		},
		BtcSellPercentage: MinMaxFloat64{
			min: -1,
			max: 1.5,
		},

		// -----------------------------------------------------
		UnsoldFirstSellDurationMinutes: MinMaxInt{
			min: 1,
			max: 10,
		},
		UnsoldFirstSellPercentage: MinMaxFloat64{
			min: 0.3,
			max: 3,
		},
		UnsoldFinalSellDurationMinutes: MinMaxInt{
			min: 11,
			max: 40,
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
			max: 50,
		},
		BtcSuperTrendMultiplier: MinMaxFloat64{
			min: 3,
			max: 45,
		},

		// -----------------------------------------------------
		AverageVolumeCandles: MinMaxInt{
			min: 8,
			max: 20,
		},
		AverageVolumeMinimal: MinMaxFloat64{
			min: 10000,
			max: 100000,
		},

		// -----------------------------------------------------
		AdxDiLen: MinMaxInt{
			min: 5,
			max: 25,
		},
		AdxBottomThreshold: MinMaxFloat64{
			min: 16,
			max: 40,
		},
		AdxTopThreshold: MinMaxFloat64{
			min: 40,
			max: 90,
		},
		AdxMinGrowthPercentage: MinMaxFloat64{
			min: 1,
			max: 50,
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

		// -----------------------------------------------------
		BtcPriceGrowthCandles: MinMaxInt{
			min: 3,
			max: 60,
		},
		BtcPriceGrowthMinPercentage: MinMaxFloat64{ // В процентах
			min: 0.1,
			max: 5,
		},
		BtcPriceGrowthMaxPercentage: MinMaxFloat64{ // В процентах
			min: 50,
			max: 200,
		},

		// -----------------------------------------------------
		PriceFallCandles: MinMaxInt{
			min: 1,
			max: 1,
		},
		PriceFallMinPercentage: MinMaxFloat64{ // В процентах, минусовые значения, можно и плюс писать
			min: 0.0,
			max: 0.5,
		},

		// -----------------------------------------------------
		TrailingLowPercentage: MinMaxFloat64{
			min: 2,
			max: 7,
		},
		TrailingTopPercentage: MinMaxFloat64{
			min: 0.7,
			max: 1,
		},
		TrailingReducePercentage: MinMaxFloat64{
			min: 0.2,
			max: 2.5,
		},
		TrailingIncreasePercentage: MinMaxFloat64{
			min: 0.2,
			max: 1,
		},

		// -----------------------------------------------------
		FlatLineCandles: MinMaxInt{
			min: 100,
			max: 400,
		},
		FlatLineSkipCandles: MinMaxInt{
			min: 0,
			max: 25,
		},
		FlatLineDispersionPercentage: MinMaxFloat64{
			min: 0.5,
			max: 2,
		},
		FlatLineOnLinePricesPercentage: MinMaxFloat64{
			min: 90,
			max: 100,
		},

		// -----------------------------------------------------
		TwoLineCandles: MinMaxInt{
			min: 200,
			max: 800,
		},
		TwoLineMaxDiffPercentage: MinMaxFloat64{
			min: 0.01,
			max: 0.03,
		},
		TwoLineSkipCandles: MinMaxInt{
			min: 0,
			max: 50,
		},

		// -----------------------------------------------------
		StopBuyAfterSellPeriodMinutes: MinMaxInt{
			min: 0,
			max: 60,
		},

		// -----------------------------------------------------
		AltCoinMarketCandles: MinMaxInt{
			min: 1,
			max: 20,
		},
		AltCoinMarketMinPercentage: MinMaxFloat64{
			min: 0.5,
			max: 50,
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

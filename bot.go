package main

const BITCOIN_SYMBOL = "BTCUSDT"
const TOTAL_MONEY_AMOUNT = 100.0
const COMMISSION = 0.15
const SIMULTANEOUS_BUYS_COUNT = 1

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

	AdxDiLen               int
	AdxBottomThreshold     float64
	AdxTopThreshold        float64
	AdxMinGrowthPercentage float64

	RealBuyTopResetReachRevenue   float64
	RealBuyBottomStopReachRevenue float64
	FakeBuyReachStopRevenue       float64

	CandleBodyCandles        int
	CandleBodyHeightMinPrice float64
	CandleBodyHeightMaxPrice float64

	BtcPriceGrowthCandles       int
	BtcPriceGrowthMinPercentage float64
	BtcPriceGrowthMaxPercentage float64

	PriceFallCandles       int
	PriceFallMinPercentage float64

	TrailingLowPercentage float64

	FlatLineCandles                int
	FlatLineSkipCandles            int
	FlatLineDispersionPercentage   float64
	FlatLineOnLinePricesPercentage float64

	TwoLineCandles           int
	TwoLineMaxDiffPercentage float64
	TwoLineSkipCandles       int

	TotalRevenue      float64
	SuccessPercentage float64

	Selection float64

	Restrictions BotConfigRestriction
}

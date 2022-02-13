package main

const BITCOIN_SYMBOL = "BTCUSDT"
const TOTAL_MONEY_AMOUNT = 100.0
const COMMISSION = 0.15

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

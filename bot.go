package main

import (
	"github.com/MaxHalford/eaopt"
	"math/rand"
)

type BotConfig struct {
	HighBuyPercentage float64
	LowSellPercentage float64

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
}

func NewBot() BotConfig {
	return BotConfig{}
}

func BotFactory(rng *rand.Rand) eaopt.Genome {
	return BotConfig{}
}

func (bot BotConfig) Evaluate() (float64, error) {
	return 0.0, nil
}

func (bot BotConfig) Mutate(rng *rand.Rand) {
	//fmt.Println("hello")
}

func (bot BotConfig) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	//eaopt.CrossUniformFloat64(X, Y.(Vector), rng)
}

func (bot BotConfig) Clone() eaopt.Genome {
	Y := bot
	return Y
}

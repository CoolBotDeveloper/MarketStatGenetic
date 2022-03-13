package main

func InitBotConfig() BotConfig {
	restrict := GetBotConfigRestrictions()

	return BotConfig{
		HighSellPercentage: GetRandFloat64Config(restrict.HighSellPercentage),
		LowSellPercentage:  GetRandFloat64Config(restrict.LowSellPercentage),

		AltCoinMinBuyFirstPeriodMinutes:  GetRandIntConfig(restrict.AltCoinMinBuyFirstPeriodMinutes),
		AltCoinMinBuyFirstPercentage:     GetRandFloat64Config(restrict.AltCoinMinBuyFirstPercentage),
		AltCoinMinBuySecondPeriodMinutes: GetRandIntConfig(restrict.AltCoinMinBuySecondPeriodMinutes),
		AltCoinMinBuySecondPercentage:    GetRandFloat64Config(restrict.AltCoinMinBuySecondPercentage),

		BtcMinBuyPeriodMinutes: GetRandIntConfig(restrict.BtcMinBuyPeriodMinutes),
		BtcMinBuyPercentage:    GetRandFloat64Config(restrict.BtcMinBuyPercentage),
		BtcSellPeriodMinutes:   GetRandIntConfig(restrict.BtcSellPeriodMinutes),
		BtcSellPercentage:      GetRandFloat64Config(restrict.BtcSellPercentage),

		UnsoldFirstSellDurationMinutes: GetRandIntConfig(restrict.UnsoldFirstSellDurationMinutes),
		UnsoldFirstSellPercentage:      GetRandFloat64Config(restrict.UnsoldFirstSellPercentage),
		UnsoldFinalSellDurationMinutes: GetRandIntConfig(restrict.UnsoldFinalSellDurationMinutes),

		AltCoinSuperTrendCandles: GetRandIntConfig(restrict.AltCoinSuperTrendCandles),
		AltCoinSuperMultiplier:   GetRandFloat64Config(restrict.AltCoinSuperMultiplier),

		BtcSuperTrendCandles:    GetRandIntConfig(restrict.BtcSuperTrendCandles),
		BtcSuperTrendMultiplier: GetRandFloat64Config(restrict.BtcSuperTrendMultiplier),

		AverageVolumeCandles: GetRandIntConfig(restrict.AverageVolumeCandles),
		AverageVolumeMinimal: GetRandFloat64Config(restrict.AverageVolumeMinimal),

		AdxDiLen:           GetRandIntConfig(restrict.AdxDiLen),
		AdxBottomThreshold: GetRandFloat64Config(restrict.AdxBottomThreshold),
		AdxTopThreshold:    GetRandFloat64Config(restrict.AdxTopThreshold),

		RealBuyTopResetReachRevenue:   GetRandFloat64Config(restrict.RealBuyTopResetReachRevenue),
		RealBuyBottomStopReachRevenue: GetRandFloat64Config(restrict.RealBuyBottomStopReachRevenue),
		FakeBuyReachStopRevenue:       GetRandFloat64Config(restrict.FakeBuyReachStopRevenue),

		CandleBodyCandles:        GetRandIntConfig(restrict.CandleBodyCandles),
		CandleBodyHeightMinPrice: GetRandFloat64Config(restrict.CandleBodyHeightMinPrice),
		CandleBodyHeightMaxPrice: GetRandFloat64Config(restrict.CandleBodyHeightMaxPrice),
	}
}

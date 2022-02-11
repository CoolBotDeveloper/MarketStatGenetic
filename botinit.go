package main

import "github.com/MaxHalford/eaopt"

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

		UnsoldFirstSellDurationMinutes: GetRandIntConfig(restrict.BtcMinBuyPeriodMinutes),
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
	}
}

func InitBotConfigSlice() BotConfigSlice {
	botConfig := InitBotConfig()

	return BotConfigSlice{
		HighSellPercentage: &eaopt.Float64Slice{botConfig.HighSellPercentage},
		LowSellPercentage:  &eaopt.Float64Slice{botConfig.LowSellPercentage},

		AltCoinMinBuyFirstPeriodMinutes:  &eaopt.IntSlice{botConfig.AltCoinMinBuyFirstPeriodMinutes},
		AltCoinMinBuyFirstPercentage:     &eaopt.Float64Slice{botConfig.AltCoinMinBuyFirstPercentage},
		AltCoinMinBuySecondPeriodMinutes: &eaopt.IntSlice{botConfig.AltCoinMinBuySecondPeriodMinutes},
		AltCoinMinBuySecondPercentage:    &eaopt.Float64Slice{botConfig.AltCoinMinBuySecondPercentage},

		BtcMinBuyPeriodMinutes: &eaopt.IntSlice{botConfig.BtcMinBuyPeriodMinutes},
		BtcMinBuyPercentage:    &eaopt.Float64Slice{botConfig.BtcMinBuyPercentage},
		BtcSellPeriodMinutes:   &eaopt.IntSlice{botConfig.BtcSellPeriodMinutes},
		BtcSellPercentage:      &eaopt.Float64Slice{botConfig.BtcSellPercentage},

		UnsoldFirstSellDurationMinutes: &eaopt.IntSlice{botConfig.UnsoldFirstSellDurationMinutes},
		UnsoldFirstSellPercentage:      &eaopt.Float64Slice{botConfig.UnsoldFirstSellPercentage},
		UnsoldFinalSellDurationMinutes: &eaopt.IntSlice{botConfig.UnsoldFinalSellDurationMinutes},

		AltCoinSuperTrendCandles: &eaopt.IntSlice{botConfig.AltCoinSuperTrendCandles},
		AltCoinSuperMultiplier:   &eaopt.Float64Slice{botConfig.AltCoinSuperMultiplier},

		BtcSuperTrendCandles:    &eaopt.IntSlice{botConfig.BtcSuperTrendCandles},
		BtcSuperTrendMultiplier: &eaopt.Float64Slice{botConfig.BtcSuperTrendMultiplier},

		AverageVolumeCandles: &eaopt.IntSlice{botConfig.AverageVolumeCandles},
		AverageVolumeMinimal: &eaopt.Float64Slice{botConfig.AverageVolumeMinimal},

		AdxDiLen:           &eaopt.IntSlice{botConfig.AdxDiLen},
		AdxBottomThreshold: &eaopt.Float64Slice{botConfig.AdxBottomThreshold},
		AdxTopThreshold:    &eaopt.Float64Slice{botConfig.AdxTopThreshold},

		Restrictions: GetBotConfigRestrictions(),
	}
}

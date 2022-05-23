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

		AdxDiLen:               GetRandIntConfig(restrict.AdxDiLen),
		AdxBottomThreshold:     GetRandFloat64Config(restrict.AdxBottomThreshold),
		AdxTopThreshold:        GetRandFloat64Config(restrict.AdxTopThreshold),
		AdxMinGrowthPercentage: GetRandFloat64Config(restrict.AdxMinGrowthPercentage),

		RealBuyTopResetReachRevenue:   GetRandFloat64Config(restrict.RealBuyTopResetReachRevenue),
		RealBuyBottomStopReachRevenue: GetRandFloat64Config(restrict.RealBuyBottomStopReachRevenue),
		FakeBuyReachStopRevenue:       GetRandFloat64Config(restrict.FakeBuyReachStopRevenue),

		CandleBodyCandles:        GetRandIntConfig(restrict.CandleBodyCandles),
		CandleBodyHeightMinPrice: GetRandFloat64Config(restrict.CandleBodyHeightMinPrice),
		CandleBodyHeightMaxPrice: GetRandFloat64Config(restrict.CandleBodyHeightMaxPrice),

		BtcPriceGrowthCandles:       GetRandIntConfig(restrict.BtcPriceGrowthCandles),
		BtcPriceGrowthMinPercentage: GetRandFloat64Config(restrict.BtcPriceGrowthMinPercentage),
		BtcPriceGrowthMaxPercentage: GetRandFloat64Config(restrict.BtcPriceGrowthMaxPercentage),

		PriceFallCandles:       GetRandIntConfig(restrict.PriceFallCandles),
		PriceFallMinPercentage: GetRandFloat64Config(restrict.PriceFallMinPercentage),

		TrailingLowPercentage: GetRandFloat64Config(restrict.TrailingLowPercentage),

		FlatLineCandles:                GetRandIntConfig(restrict.FlatLineCandles),
		FlatLineSkipCandles:            GetRandIntConfig(restrict.FlatLineSkipCandles),
		FlatLineDispersionPercentage:   GetRandFloat64Config(restrict.FlatLineDispersionPercentage),
		FlatLineOnLinePricesPercentage: GetRandFloat64Config(restrict.FlatLineOnLinePricesPercentage),

		TwoLineCandles:           GetRandIntConfig(restrict.TwoLineCandles),
		TwoLineMaxDiffPercentage: GetRandFloat64Config(restrict.TwoLineMaxDiffPercentage),
		TwoLineSkipCandles:       GetRandIntConfig(restrict.TwoLineSkipCandles),

		TrailingTopPercentage:      GetRandFloat64Config(restrict.TrailingTopPercentage),
		TrailingReducePercentage:   GetRandFloat64Config(restrict.TrailingReducePercentage),
		TrailingIncreasePercentage: GetRandFloat64Config(restrict.TrailingIncreasePercentage),

		StopBuyAfterSellPeriodMinutes: GetRandIntConfig(restrict.StopBuyAfterSellPeriodMinutes),

		AltCoinMarketCandles:       GetRandIntConfig(restrict.AltCoinMarketCandles),
		AltCoinMarketMinPercentage: GetRandFloat64Config(restrict.AltCoinMarketMinPercentage),

		AltCoinMinBuyMaxSecondPercentage: GetRandFloat64Config(restrict.AltCoinMinBuyMaxSecondPercentage),

		WholeDayTotalVolumeCandles:   GetRandIntConfig(restrict.WholeDayTotalVolumeCandles),
		WholeDayTotalVolumeMinVolume: GetRandFloat64Config(restrict.WholeDayTotalVolumeMinVolume),

		HalfVolumeFirstCandles:     GetRandIntConfig(restrict.HalfVolumeFirstCandles),
		HalfVolumeSecondCandles:    GetRandIntConfig(restrict.HalfVolumeSecondCandles),
		HalfVolumeGrowthPercentage: GetRandFloat64Config(restrict.HalfVolumeGrowthPercentage),

		TrailingActivationPercentage: GetRandFloat64Config(restrict.TrailingActivationPercentage),

		FlatLineSearchWindowCandles:          GetRandIntConfig(restrict.FlatLineSearchWindowCandles),
		FlatLineSearchWindowsCount:           GetRandIntConfig(restrict.FlatLineSearchWindowsCount),
		FlatLineSearchDispersionPercentage:   GetRandFloat64Config(restrict.FlatLineSearchDispersionPercentage),
		FlatLineSearchOnLinePricesPercentage: GetRandFloat64Config(restrict.FlatLineSearchOnLinePricesPercentage),
		FlatLineSearchRelativePeriodCandles:  GetRandIntConfig(restrict.FlatLineSearchRelativePeriodCandles),

		TripleGrowthCandles:          GetRandIntConfig(restrict.TripleGrowthCandles),
		TripleGrowthSecondPercentage: GetRandFloat64Config(restrict.TripleGrowthSecondPercentage),

		PastMaxPricePeriod: GetRandIntConfig(restrict.PastMaxPricePeriod),

		SmoothGrowthCandles: GetRandIntConfig(restrict.SmoothGrowthCandles),
		SmoothGrowthAngle:   GetRandFloat64Config(restrict.SmoothGrowthAngle),

		EachVolumeMinValueCandles:     GetRandIntConfig(restrict.EachVolumeMinValueCandles),
		EachVolumeMinValueMinVolume:   GetRandFloat64Config(restrict.EachVolumeMinValueMinVolume),
		EachVolumeMinValueSkipCandles: GetRandIntConfig(restrict.EachVolumeMinValueSkipCandles),

		TrailingFixationActivatePercentage:  GetRandFloat64Config(restrict.TrailingFixationActivatePercentage),
		TrailingFixationPercentage:          GetRandFloat64Config(restrict.TrailingFixationPercentage),
		TrailingSecondaryIncreasePercentage: GetRandFloat64Config(restrict.TrailingSecondaryIncreasePercentage),

		AltCoinMaxCandles:    GetRandIntConfig(restrict.AltCoinMaxCandles),
		AltCoinMaxPercentage: GetRandFloat64Config(restrict.AltCoinMaxPercentage),
	}
}

package main

type PositiveApproach struct {
	botConfig                 BotConfig
	exchangeManager           *ExchangeManager
	candleMarketStat          *CandleMarketStat
	isRealBuyBlocked          bool
	realReachedStartPointTime string // By default, is empty, after each reset updates the time to new one.
}

func NewPositiveApproach(botConfig BotConfig, exchangeManager *ExchangeManager, candleMarketStat *CandleMarketStat) PositiveApproach {
	return PositiveApproach{
		botConfig:        botConfig,
		exchangeManager:  exchangeManager,
		candleMarketStat: candleMarketStat,
		isRealBuyBlocked: false,
	}
}

func (pa *PositiveApproach) HasSignal(candle Candle) bool {
	symbol := candle.Symbol
	exchangeRate := candle.ClosePrice

	if pa.isRealBuyBlocked {
		if pa.isFakeBuyReached(candle) {
			pa.activateRealBuy(candle)
			return true
		}

		pa.FakeBuy(symbol, exchangeRate, candle.CloseTime)
		return false
	}

	if pa.isRealBuyPositive(candle) {
		if pa.isRealBuyReached(candle) {
			pa.resetRealBuyReached(candle)
		}

		pa.activateRealBuy(candle)
		return true
	}

	pa.FakeBuy(symbol, exchangeRate, candle.CloseTime)
	return false
}

func (pa *PositiveApproach) UpdateBuys(candle Candle) {
	if !pa.isRealBuyBlocked {
		return
	}

	pa.updateNormalBuys(candle)
	pa.updateFirstSellZombies(candle)
	pa.updateExitZombies(candle)

	if pa.candleMarketStat.HasBtcSellPercentage() {
		pa.updateAllExitSymbols(candle)
	}
}

func (pa *PositiveApproach) updateNormalBuys(candle Candle) {
	exchangeRate := candle.ClosePrice
	fakeUnsoldBuys := pa.getStorage().FindFakeUnsoldBuys(
		candle.Symbol,
		exchangeRate,
		pa.botConfig.HighSellPercentage,
		pa.botConfig.LowSellPercentage,
		candle.CloseTime,
	)

	for _, buy := range fakeUnsoldBuys {
		calcRevenue := pa.exchangeManager.calcRevenue(buy.coins, exchangeRate)
		pa.getStorage().AddFakeSell(
			candle.Symbol,
			buy.coins,
			exchangeRate,
			calcRevenue,
			buy.id,
			candle.CloseTime,
		)
	}
}

func (pa *PositiveApproach) updateFirstSellZombies(candle Candle) {
	exchangeRate := candle.ClosePrice
	fakeFirstSellZombies := pa.getStorage().FindFakeFirstSellZombies(
		candle.Symbol,
		exchangeRate,
		candle.CloseTime,
		pa.botConfig.UnsoldFirstSellDurationMinutes,
		pa.botConfig.UnsoldFirstSellPercentage,
	)

	for _, expiredBuy := range fakeFirstSellZombies {
		calcRevenue := expiredBuy.coins * exchangeRate
		pa.getStorage().AddFakeSell(
			candle.Symbol,
			expiredBuy.coins,
			exchangeRate,
			calcRevenue,
			expiredBuy.id,
			candle.CloseTime,
		)
	}
}

func (pa *PositiveApproach) updateExitZombies(candle Candle) {
	exchangeRate := candle.ClosePrice
	fakeExitZombies := pa.getStorage().FindFakeExitZombies(
		candle.Symbol,
		candle.CloseTime,
		pa.botConfig.UnsoldFinalSellDurationMinutes,
	)

	for _, expiredBuy := range fakeExitZombies {
		calcRevenue := expiredBuy.coins * exchangeRate
		pa.getStorage().AddFakeSell(
			candle.Symbol,
			expiredBuy.coins,
			exchangeRate,
			calcRevenue,
			expiredBuy.id,
			candle.CloseTime,
		)
	}
}

func (pa *PositiveApproach) updateAllExitSymbols(candle Candle) {
	exchangeRate := candle.ClosePrice
	fakeExitZombies := pa.getStorage().FindFakeExitZombies(candle.Symbol, candle.CloseTime, 0)

	for _, expiredBuy := range fakeExitZombies {
		calcRevenue := expiredBuy.coins * exchangeRate
		pa.getStorage().AddFakeSell(
			candle.Symbol,
			expiredBuy.coins,
			exchangeRate,
			calcRevenue,
			expiredBuy.id,
			candle.CloseTime,
		)
	}
}

func (pa *PositiveApproach) isRealBuyPositive(candle Candle) bool {
	startTime := pa.realReachedStartPointTime
	if startTime == "" {
		startTime = "2000-01-23 12:09:59"
	}

	startTimeRevenue := pa.getStorage().CalculateRevenueFromStartTime(candle.Symbol, startTime)
	realBuysCount := pa.getStorage().GetBuysCount()

	return realBuysCount == 0 || startTimeRevenue > pa.botConfig.RealBuyBottomStopReachRevenue
}

func (pa *PositiveApproach) isRealBuyReached(candle Candle) bool {
	startTime := pa.realReachedStartPointTime
	if startTime == "" {
		startTime = "2000-01-23 12:09:59"
	}

	startTimeRevenue := pa.getStorage().CalculateRevenueFromStartTime(candle.Symbol, startTime)

	return startTimeRevenue >= pa.botConfig.RealBuyTopResetReachRevenue
}

func (pa *PositiveApproach) isFakeBuyReached(candle Candle) bool {
	fakeRevenue := pa.getStorage().GetFakeTotalRevenue(candle.Symbol)

	return fakeRevenue >= pa.botConfig.FakeBuyReachStopRevenue
}

func (pa *PositiveApproach) resetRealBuyReached(candle Candle) {
	pa.realReachedStartPointTime = candle.CloseTime
}

func (pa *PositiveApproach) FakeBuy(symbol string, exchangeRate float64, createdAt string) {
	pa.blockRealBuy()

	unsoldBuysCount := pa.getStorage().CountFakeUnsoldBuys(symbol)
	if SIMULTANEOUS_BUYS_COUNT > unsoldBuysCount {
		coinsCount := TOTAL_MONEY_AMOUNT / exchangeRate
		pa.getStorage().AddFakeBuy(symbol, coinsCount, exchangeRate, createdAt)
	}
}

func (pa *PositiveApproach) FakeSell(candle Candle) {
	if pa.isRealBuyBlocked {
		// TODO: do fake sell
	}
}

func (pa *PositiveApproach) activateRealBuy(candle Candle) {
	pa.isRealBuyBlocked = false
	pa.getStorage().CleanFakeBuySellTables(candle.Symbol)
}

func (pa *PositiveApproach) blockRealBuy() {
	pa.isRealBuyBlocked = true
}

func (pa *PositiveApproach) getStorage() *Storage {
	return &pa.exchangeManager.storage
}

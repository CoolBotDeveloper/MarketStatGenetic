package main

import "time"

type PositiveApproach struct {
	botConfig                 BotConfig
	exchangeManager           ExchangeManager
	candleMarketStat          CandleMarketStat
	realBuyReachRevenue       float64
	fakeBuyReachRevenue       float64
	isRealBuyBlocked          bool
	realReachedStartPointTime string // By default, is empty, after each reset updates the time to new one.
}

func NewPositiveApproach(botConfig BotConfig, exchangeManager ExchangeManager, candleMarketStat CandleMarketStat) PositiveApproach {
	return PositiveApproach{
		botConfig:           botConfig,
		exchangeManager:     exchangeManager,
		candleMarketStat:    candleMarketStat,
		realBuyReachRevenue: botConfig.RealBuyReachRevenue,
		fakeBuyReachRevenue: botConfig.FakeBuyReachRevenue,
		isRealBuyBlocked:    false,
	}
}

func (pa *PositiveApproach) HasSignal(symbol string, exchangeRate float64) bool {
	pa.updateBuys(symbol, exchangeRate)

	if pa.isRealBuyBlocked {
		if pa.isFakeBuyReached() {
			pa.activateRealBuy()
			return true
		}

		pa.FakeBuy(symbol, exchangeRate)
		return false
	}

	if pa.isRealBuyPositive() {
		if pa.isRealBuyReached() {
			pa.resetRealBuyReached()
		}
		return true
	}

	pa.FakeBuy(symbol, exchangeRate)
	return false
}

func (pa *PositiveApproach) updateBuys(symbol string, exchangeRate float64) {
	// TODO: do fake sells by different conditions
}

func (pa *PositiveApproach) isRealBuyPositive() bool {
	return false
}

func (pa *PositiveApproach) isRealBuyReached() bool {
	return false
}

func (pa *PositiveApproach) isFakeBuyReached() bool {
	return false
}

func (pa *PositiveApproach) resetRealBuyReached() {

}

func (pa *PositiveApproach) FakeBuy(symbol string, exchangeRate float64) {
	pa.blockRealBuy()

	createdAt := time.Now().Format("2006-01-02 15:04:05")
	coinsCount := TOTAL_MONEY_AMOUNT / exchangeRate
	pa.getStorage().AddFakeBuy(symbol, coinsCount, exchangeRate, createdAt)
}

func (pa *PositiveApproach) FakeSell() {
	if pa.isRealBuyBlocked {
		// TODO: do fake sell
	}
}

func (pa *PositiveApproach) activateRealBuy() {
	pa.isRealBuyBlocked = false
}

func (pa *PositiveApproach) blockRealBuy() {
	pa.isRealBuyBlocked = true
}

func (pa *PositiveApproach) getStorage() *Storage {
	return &pa.exchangeManager.storage
}

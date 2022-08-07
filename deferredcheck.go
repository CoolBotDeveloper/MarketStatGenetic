package main

type DeferredCheck struct {
	signals    map[string]Candle
	config     *BotConfig
	dataSource *DataSource
}

func NewDeferredCheck(config *BotConfig, dataSource *DataSource) DeferredCheck {
	return DeferredCheck{
		config:     config,
		dataSource: dataSource,
		signals:    map[string]Candle{},
	}
}

func (deferred *DeferredCheck) AddForCandle(candle Candle) {
	if _, ok := deferred.signals[candle.Symbol]; !ok {
		deferred.signals[candle.Symbol] = candle
	}
}

func (deferred *DeferredCheck) DeleteForSymbol(symbol string) {
	if _, ok := deferred.signals[symbol]; ok {
		delete(deferred.signals, symbol)
	}
}

func (deferred *DeferredCheck) HasDeferred(candle Candle) bool {
	_, ok := deferred.signals[candle.Symbol]

	return ok
}

func (deferred *DeferredCheck) CheckForCandle(candle Candle) bool {
	if deferred.config.DeferredCheckInterval == 0 {
		return true
	}

	if signalCandle, ok := deferred.signals[candle.Symbol]; ok {
		cur := ConvertDateStringToTime(candle.CloseTime)
		diff := cur.Unix() - ConvertDateStringToTime(signalCandle.CloseTime).Unix()

		if diff < int64(60*deferred.config.DeferredCheckInterval) {
			return false
		}

		candles := deferred.dataSource.GetCandlesFor(candle.Symbol)
		closePrices := GetHighPrice(candles, deferred.config.DeferredCheckInterval)
		max := Max(closePrices)

		isMore := max < candle.ClosePrice
		deferred.DeleteForSymbol(candle.Symbol)

		return isMore
	}

	return false
}

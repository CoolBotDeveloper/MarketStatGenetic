package main

const INTERVAL = 5

type DeferredCheck struct {
	signals    map[string]Candle
	dataSource *DataSource
}

func NewDeferredCheck(dataSource *DataSource) DeferredCheck {
	return DeferredCheck{
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
	if signalCandle, ok := deferred.signals[candle.Symbol]; ok {
		cur := ConvertDateStringToTime(candle.CloseTime)
		diff := cur.Unix() - ConvertDateStringToTime(signalCandle.CloseTime).Unix()

		if diff < (60 * INTERVAL) {
			return false
		}

		isMore := candle.ClosePrice > signalCandle.ClosePrice
		deferred.DeleteForSymbol(candle.Symbol)

		return isMore
	}

	return false
}

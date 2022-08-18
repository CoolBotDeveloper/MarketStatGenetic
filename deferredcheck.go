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
		closePrices := GetClosePrice(candles, deferred.config.DeferredCheckInterval)

		isAllGrowing := true
		for index := range closePrices {
			if 0 == index {
				continue
			}

			prev := closePrices[index-1]
			current := closePrices[index]

			if prev > current {
				isAllGrowing = false
				break
			}
		}

		deferred.DeleteForSymbol(candle.Symbol)

		return isAllGrowing
	}

	return false
}

func (deferred *DeferredCheck) CheckForCandleByNeural(candle Candle) bool {
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
		neuralIndicator := NewNeuralNetworkIndicator(*deferred.config)
		hasBuySignal := neuralIndicator.HasBuySignal(candles)

		deferred.DeleteForSymbol(candle.Symbol)

		return hasBuySignal
	}

	return false
}

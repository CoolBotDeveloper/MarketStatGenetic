package main

type DeferredSell struct {
	signals map[string]Candle
	config  *BotConfig
}

func NewDeferredSell(config *BotConfig) DeferredSell {
	return DeferredSell{
		signals: map[string]Candle{},
		config:  config,
	}
}

func (deferred *DeferredSell) Start(candle Candle) {
	if _, ok := deferred.signals[candle.Symbol]; !ok {
		deferred.signals[candle.Symbol] = candle
	}
}

func (deferred *DeferredSell) Finish(symbol string) {
	if _, ok := deferred.signals[symbol]; ok {
		delete(deferred.signals, symbol)
	}
}

func (deferred *DeferredSell) CanSell(candle Candle) bool {
	if signalCandle, ok := deferred.signals[candle.Symbol]; ok {
		cur := ConvertDateStringToTime(candle.CloseTime)
		diff := cur.Unix() - ConvertDateStringToTime(signalCandle.CloseTime).Unix()

		return diff >= int64(60*deferred.config.DeferredSellInterval)
	}

	return false
}

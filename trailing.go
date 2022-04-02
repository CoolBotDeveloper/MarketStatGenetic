package main

type Trailing struct {
	LowPercentage float64
	Items         map[string]*TrailingSymbol
}

type TrailingSymbol struct {
	Symbol         string
	StopPrice      float64
	PreviousPrices []float64
}

func NewTrailingSymbol(lowPercentage float64) Trailing {
	return Trailing{
		Items:         map[string]*TrailingSymbol{},
		LowPercentage: lowPercentage,
	}
}

func (trailing *Trailing) Start(candle Candle) {
	trailing.initiateSymbolTrailing(candle)
}

func (trailing *Trailing) Update(candle Candle) bool {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		// if there is not enough prices just skip it
		if len(trailingSymbol.PreviousPrices) < 1 {
			return false
		}

		trailing.appendPrice(candle)
		if trailing.isGrowing(candle) {
			newStopPrice := trailing.calculateStopPrice(candle.ClosePrice)

			// just update if the new calculated stop price higher than old one
			if newStopPrice > trailingSymbol.StopPrice {
				trailingSymbol.StopPrice = trailing.calculateStopPrice(candle.ClosePrice)

				return true
			}
		}
		return false
	}

	return false
	//panic(fmt.Sprintf("There is no Trailing data for Symbol <%s>", candle.Symbol))
}

func (trailing *Trailing) Finish(candle Candle) {
	if _, ok := trailing.Items[candle.Symbol]; ok {
		delete(trailing.Items, candle.Symbol)
	}
}

func (trailing *Trailing) CanSellByStop(candle Candle) bool {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		return candle.ClosePrice <= trailingSymbol.StopPrice
	}

	return false
}

func (trailing *Trailing) initiateSymbolTrailing(candle Candle) {
	trailing.Items[candle.Symbol] = &TrailingSymbol{
		Symbol:         candle.Symbol,
		StopPrice:      trailing.calculateStopPrice(candle.ClosePrice),
		PreviousPrices: []float64{candle.ClosePrice},
	}
}

func (trailing *Trailing) calculateStopPrice(closePrice float64) float64 {
	return closePrice - ((closePrice * trailing.LowPercentage) / 100)
}

func (trailing *Trailing) appendPrice(candle Candle) {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		prices := append(trailingSymbol.PreviousPrices, candle.ClosePrice)
		last := len(prices) - 1

		trailingSymbol.PreviousPrices = []float64{
			prices[last-1],
			prices[last],
		}
	}
}

func (trailing *Trailing) isGrowing(candle Candle) bool {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		last := len(trailingSymbol.PreviousPrices) - 1

		currentPrice := trailingSymbol.PreviousPrices[last]
		previousPrice := trailingSymbol.PreviousPrices[last-1]

		return currentPrice > previousPrice
	}

	panic("There is no prices")
}

package main

import "fmt"

type Trailing struct {
	Items map[string]*TrailingSymbol

	TopPercentage      float64
	BottomPercentage   float64
	ReducePercentage   float64
	IncreasePercentage float64
}

type TrailingSymbol struct {
	Symbol            string
	StopPrice         float64
	PreviousPrices    []float64
	CurrentPercentage float64
	LastMaxPrice      float64
}

func NewTrailingSymbol(config BotConfig) Trailing {
	return Trailing{
		Items: map[string]*TrailingSymbol{},

		TopPercentage:      config.TrailingTopPercentage,
		BottomPercentage:   config.TrailingLowPercentage,
		ReducePercentage:   config.TrailingReducePercentage,
		IncreasePercentage: config.TrailingIncreasePercentage,
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
			isHigherThanLastMaxPrice := candle.ClosePrice > trailingSymbol.LastMaxPrice
			if isHigherThanLastMaxPrice {
				trailingSymbol.LastMaxPrice = candle.ClosePrice
			}

			if isHigherThanLastMaxPrice {
				trailing.increasePercentage(trailingSymbol)

				fmt.Println(fmt.Sprintf("Trailing INCREASED %f, StopPrice: %f:, COIN: %s, EXCHANGE_RATE: %f, TIME: %s",
					trailingSymbol.CurrentPercentage, trailingSymbol.StopPrice, candle.Symbol, candle.ClosePrice, candle.CloseTime))
			} else {
				// if not growing, step by step reduce low percentage
				trailing.reducePercentage(trailingSymbol)

				fmt.Println(fmt.Sprintf("Trailing REDUCED %f, StopPrice: %f, : COIN: %s, EXCHANGE_RATE: %f, TIME: %s",
					trailingSymbol.CurrentPercentage, trailingSymbol.StopPrice, candle.Symbol, candle.ClosePrice, candle.CloseTime))
			}
		} else {
			// if not growing, step by step reduce low percentage
			trailing.reducePercentage(trailingSymbol)

			fmt.Println(fmt.Sprintf("Trailing REDUCED %f, StopPrice: %f, : COIN: %s, EXCHANGE_RATE: %f, TIME: %s",
				trailingSymbol.CurrentPercentage, trailingSymbol.StopPrice, candle.Symbol, candle.ClosePrice, candle.CloseTime))
		}

		newStopPrice := trailing.calculateStopPrice(trailingSymbol.LastMaxPrice, trailingSymbol.CurrentPercentage)
		// just update if the new calculated stop price higher than old one
		if newStopPrice > trailingSymbol.StopPrice {
			trailingSymbol.StopPrice = newStopPrice

			return true
		}
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

func (trailing *Trailing) GetStopPrice(candle Candle) (float64, bool) {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		return trailingSymbol.StopPrice, true
	}

	return 0.0, false
}

func (trailing *Trailing) reducePercentage(trailingSymbol *TrailingSymbol) {
	if trailing.TopPercentage == trailingSymbol.CurrentPercentage {
		return
	}

	reducedPercentage := trailingSymbol.CurrentPercentage - trailing.ReducePercentage
	if trailing.TopPercentage > reducedPercentage {
		reducedPercentage = trailing.TopPercentage
	}

	trailingSymbol.CurrentPercentage = reducedPercentage
}

func (trailing *Trailing) increasePercentage(trailingSymbol *TrailingSymbol) {
	if trailing.BottomPercentage == trailingSymbol.CurrentPercentage {
		return
	}

	increasedPercentage := trailingSymbol.CurrentPercentage + trailing.IncreasePercentage
	if trailing.BottomPercentage < increasedPercentage {
		increasedPercentage = trailing.BottomPercentage
	}

	trailingSymbol.CurrentPercentage = increasedPercentage
}

func (trailing *Trailing) initiateSymbolTrailing(candle Candle) {
	trailing.Items[candle.Symbol] = &TrailingSymbol{
		Symbol:            candle.Symbol,
		PreviousPrices:    []float64{candle.ClosePrice},
		StopPrice:         trailing.calculateStopPrice(candle.ClosePrice, trailing.BottomPercentage),
		CurrentPercentage: trailing.BottomPercentage,
		LastMaxPrice:      candle.ClosePrice,
	}
}

func (trailing *Trailing) calculateStopPrice(closePrice, percentage float64) float64 {
	return closePrice - ((closePrice * percentage) / 100)
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

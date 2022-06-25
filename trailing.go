package main

import (
	"fmt"
	"math"
)

type Trailing struct {
	Items                       map[string]*TrailingSymbol
	TopPercentage               float64
	BottomPercentage            float64
	ReducePercentage            float64
	IncreasePercentage          float64
	ActivationPercentage        float64
	FixationActivatePercentage  float64
	FixationPercentage          float64
	SecondaryIncreasePercentage float64
	IncreaseSpeedCoefficient    float64
	ReduceSpeedCoefficient      float64

	dataSource *DataSource
}

type TrailingSymbol struct {
	Symbol               string
	StopPrice            float64
	PreviousPrices       []float64
	CurrentPercentage    float64
	LastMaxPrice         float64
	IsPrevLastPercentage bool
	PrevLastSellPrice    float64
	UpdatesCount         int
	ReduceNumber         int
	FirstPrice           float64
	FixationEnabled      bool
	LastPrice            float64
}

func NewTrailingSymbol(config BotConfig, dataSource *DataSource) Trailing {
	return Trailing{
		Items:                       map[string]*TrailingSymbol{},
		TopPercentage:               config.TrailingTopPercentage,
		BottomPercentage:            config.TrailingLowPercentage,
		ReducePercentage:            config.TrailingReducePercentage,
		IncreasePercentage:          config.TrailingIncreasePercentage,
		ActivationPercentage:        config.TrailingActivationPercentage,
		FixationActivatePercentage:  config.TrailingFixationActivatePercentage,
		FixationPercentage:          config.TrailingFixationPercentage,
		SecondaryIncreasePercentage: config.TrailingSecondaryIncreasePercentage,
		IncreaseSpeedCoefficient:    config.TrailingIncreaseSpeedCoefficient,
		ReduceSpeedCoefficient:      config.TrailingReduceSpeedCoefficient,

		dataSource: dataSource,
	}
}

func (trailing *Trailing) Start(candle Candle) {
	trailing.initiateSymbolTrailing(candle)
}

func (trailing *Trailing) Update(candle Candle) bool {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		trailingSymbol.LastPrice = candle.ClosePrice

		trailing.appendPrice(candle)

		// just update the fixation activation flag
		trailing.processFixation(candle, trailingSymbol)

		// if there is not enough prices just skip it
		if len(trailingSymbol.PreviousPrices) < 1 {
			return false
		}

		if trailing.isGrowing(candle) {
			isHigherThanLastMaxPrice := candle.ClosePrice > trailingSymbol.LastMaxPrice
			if isHigherThanLastMaxPrice {
				trailingSymbol.LastMaxPrice = candle.ClosePrice

				trailing.increasePercentage(trailingSymbol)

				fmt.Println(fmt.Sprintf("Trailing INCREASED %f, StopPrice: %f:, COIN: %s, EXCHANGE_RATE: %f, TIME: %s",
					trailingSymbol.CurrentPercentage, trailingSymbol.StopPrice, candle.Symbol, candle.ClosePrice, candle.CloseTime))
			} else {
				// if not growing, step by step reduce low percentage
				//trailing.reducePercentage(trailingSymbol)
				trailing.increaseSecondaryPercentage(trailingSymbol) //без Coefficient
				//trailing.increasePercentageByCoefficient(trailingSymbol) //c Coefficient

				fmt.Println(fmt.Sprintf("Trailing INCREASE COEFFICIENT %f, StopPrice: %f, : COIN: %s, EXCHANGE_RATE: %f, TIME: %s",
					trailingSymbol.CurrentPercentage, trailingSymbol.StopPrice, candle.Symbol, candle.ClosePrice, candle.CloseTime))
			}
		} else {
			// if not growing, step by step reduce low percentage
			trailing.reducePercentage(trailingSymbol) //без Coefficient
			//trailing.reducePercentageByCoefficient(trailingSymbol) //c Coefficient

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
		//if trailingSymbol.UpdatesCount == 1 && trailingSymbol.ReduceNumber == 1 {
		//	return true
		//}

		//if ok2, _ := trailing.willBeReduceToFinal(candle); ok2 {
		//	return true
		//}

		//if trailingSymbol.FixationEnabled {
		//	fixationPrice := trailing.calculateFixationPrice(
		//		trailingSymbol.FirstPrice,
		//		trailing.FixationPercentage,
		//	)
		//
		//	// if we have crossed the fixation price and the stopPrice less than fixation price
		//	if candle.ClosePrice <= fixationPrice && fixationPrice > trailingSymbol.StopPrice {
		//		return true
		//	}
		//}

		//if trailing.isTriggerReached(candle) {
		//	return true
		//}

		return trailingSymbol.StopPrice >= candle.ClosePrice || trailingSymbol.StopPrice >= candle.LowPrice
	}

	return false
}

func (trailing *Trailing) GetStopPrice(candle Candle) (float64, bool) {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {

		if ok2, _ := trailing.willBeReduceToFinal(candle); ok2 {
			offsetPrice := trailing.calculateOffsetPrice(candle.ClosePrice, 0.0) // от последней

			return offsetPrice, true
		}

		//if trailingSymbol.FixationEnabled {
		//	fixationPrice := trailing.calculateFixationPrice(trailingSymbol.FirstPrice, trailing.FixationPercentage)
		//
		//	if fixationPrice >= candle.ClosePrice && fixationPrice > trailingSymbol.StopPrice {
		//		offsetPrice := trailing.calculateOffsetPrice(fixationPrice, 0.0)
		//
		//		return offsetPrice, true
		//	}
		//}

		//if trailing.isTriggerReached(candle) {
		//	offsetPrice := trailing.calculateOffsetPrice(trailingSymbol.StopPrice, 0.15) // от последней
		//
		//	return offsetPrice, true
		//}

		offsetPrice := trailing.calculateOffsetPrice(trailingSymbol.StopPrice, 0.0) // от последней

		return offsetPrice, true
	}

	return 0.0, false
}

func (trailing *Trailing) isTriggerReached(candle Candle) bool {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		triggerPercentage := CalcGrowth(trailingSymbol.StopPrice, candle.ClosePrice)
		//triggerPercentageLow := CalcGrowth(trailingSymbol.StopPrice, candle.LowPrice)

		return trailing.ActivationPercentage >= triggerPercentage /* ||
		trailing.ActivationPercentage >= triggerPercentageLow*/
	}

	return false
}

func (trailing *Trailing) IsPrevLastPercentage(candle Candle) bool {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		//if trailingSymbol.IsPrevLastPercentage {
		//	return true
		//}

		prevLastPercentage := trailing.TopPercentage + trailing.ReducePercentage
		if trailing.TopPercentage < trailingSymbol.CurrentPercentage && trailingSymbol.CurrentPercentage <= prevLastPercentage {
			trailingSymbol.IsPrevLastPercentage = true
		}

		return trailingSymbol.IsPrevLastPercentage
	}

	return false
}

func (trailing *Trailing) IsLastPercentage(candle Candle) bool {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		return trailing.TopPercentage == trailingSymbol.CurrentPercentage
	}

	return false
}

// in case of zero top final is also zero
func (trailing *Trailing) willBeReduceToFinal(candle Candle) (bool, float64) {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		if trailing.isGrowingFinal(candle) {
			return false, 0.0
		}

		reducePercentage := trailing.getReducePercentage(trailingSymbol)
		newStopPrice := trailing.calculateStopPrice(trailingSymbol.LastMaxPrice, reducePercentage)

		if trailingSymbol.LastMaxPrice == candle.ClosePrice {
			return false, 0.0
		}

		if newStopPrice > trailingSymbol.StopPrice {
			isCrossedStop := newStopPrice >= candle.ClosePrice || newStopPrice >= candle.LowPrice

			return isCrossedStop, newStopPrice
		}

		//if reducePercentage == trailing.TopPercentage {
		//	return true, newStopPrice
		//}

		return false, 0.0
	}

	return false, 0.0
}

func (trailing *Trailing) getReducePercentage(trailingSymbol *TrailingSymbol) float64 {
	if trailing.TopPercentage == trailingSymbol.CurrentPercentage {
		return trailing.TopPercentage
	}

	reducedPercentage := trailingSymbol.CurrentPercentage - trailing.ReducePercentage
	if trailing.TopPercentage > reducedPercentage {
		reducedPercentage = trailing.TopPercentage
	}

	return reducedPercentage
}

func (trailing *Trailing) isGrowingFinal(candle Candle) bool {
	if trailingSymbol, ok := trailing.Items[candle.Symbol]; ok {
		if len(trailingSymbol.PreviousPrices) < 2 {
			return true
		}

		last := len(trailingSymbol.PreviousPrices) - 1
		//firstPrice := trailingSymbol.PreviousPrices[0]
		lastPrice := trailingSymbol.PreviousPrices[last]

		if futureCandle, ok2 := trailing.dataSource.GetNextCandle(candle.Symbol, candle.Index); ok2 {
			return futureCandle.ClosePrice > lastPrice
		}

		return true
	}

	panic("There is no prices")
}

func (trailing *Trailing) reducePercentage(trailingSymbol *TrailingSymbol) {
	trailingSymbol.ReduceNumber++

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

func (trailing *Trailing) increasePercentageByCoefficient(trailingSymbol *TrailingSymbol) {
	if trailing.BottomPercentage == trailingSymbol.CurrentPercentage {
		return
	}

	speed := trailing.calcSpeed(trailingSymbol)
	percentageBySpeed := trailing.IncreaseSpeedCoefficient * speed
	increasedPercentage := trailingSymbol.CurrentPercentage + percentageBySpeed

	if trailing.BottomPercentage < increasedPercentage {
		increasedPercentage = trailing.BottomPercentage
	}

	trailingSymbol.CurrentPercentage = increasedPercentage
}

func (trailing *Trailing) reducePercentageByCoefficient(trailingSymbol *TrailingSymbol) {
	trailingSymbol.ReduceNumber++

	if trailing.TopPercentage == trailingSymbol.CurrentPercentage {
		return
	}

	speed := trailing.calcSpeed(trailingSymbol)
	percentageBySpeed := trailing.ReduceSpeedCoefficient * speed
	reducedPercentage := trailingSymbol.CurrentPercentage - percentageBySpeed

	if trailing.TopPercentage > reducedPercentage {
		reducedPercentage = trailing.TopPercentage
	}

	trailingSymbol.CurrentPercentage = reducedPercentage
}

func (trailing *Trailing) calcSpeed(trailingSymbol *TrailingSymbol) float64 {
	last := len(trailingSymbol.PreviousPrices) - 1
	startPrice := trailingSymbol.PreviousPrices[last-1] // previous price
	endPrice := trailingSymbol.PreviousPrices[last]     // current price

	delta := CalcGrowth(startPrice, endPrice)
	x2 := 1.0

	return math.Abs(delta / x2)
}

func (trailing *Trailing) increaseSecondaryPercentage(trailingSymbol *TrailingSymbol) {
	if trailing.BottomPercentage == trailingSymbol.CurrentPercentage {
		return
	}

	secondaryIncreaseValue := (trailing.IncreasePercentage * trailing.SecondaryIncreasePercentage) / 100
	increasedPercentage := trailingSymbol.CurrentPercentage + secondaryIncreaseValue

	if trailing.BottomPercentage < increasedPercentage {
		increasedPercentage = trailing.BottomPercentage
	}

	trailingSymbol.CurrentPercentage = increasedPercentage
}

func (trailing *Trailing) initiateSymbolTrailing(candle Candle) {
	trailing.Items[candle.Symbol] = &TrailingSymbol{
		Symbol:               candle.Symbol,
		PreviousPrices:       []float64{candle.ClosePrice},
		StopPrice:            trailing.calculateStopPrice(candle.ClosePrice, trailing.BottomPercentage),
		CurrentPercentage:    trailing.BottomPercentage,
		LastMaxPrice:         candle.ClosePrice,
		IsPrevLastPercentage: false,
		PrevLastSellPrice:    0.0,
		FirstPrice:           candle.ClosePrice,
		FixationEnabled:      false,
		LastPrice:            candle.ClosePrice,
	}
}

func (trailing *Trailing) processFixation(candle Candle, trailingSymbol *TrailingSymbol) {
	fixationEnablePrice := trailing.calculateFixationPrice(
		trailingSymbol.FirstPrice,
		trailing.FixationActivatePercentage,
	)

	if candle.ClosePrice >= fixationEnablePrice {
		trailingSymbol.FixationEnabled = true
	}
}

func (trailing *Trailing) calculateFixationPrice(closePrice, percentage float64) float64 {
	return closePrice + ((closePrice * percentage) / 100)
}

func (trailing *Trailing) calculateStopPrice(closePrice, percentage float64) float64 {
	return closePrice - ((closePrice * percentage) / 100)
}

func (trailing *Trailing) calculateOffsetPrice(closePrice, percentage float64) float64 {
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

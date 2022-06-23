package main

type DataSource struct {
	symbolCandlesMap map[string][]Candle
}

func NewDataSource() DataSource {
	return DataSource{map[string][]Candle{}}
}

func (dataSource *DataSource) AddCandleFor(symbol string, candle Candle) {
	if _, ok := dataSource.symbolCandlesMap[symbol]; !ok {
		dataSource.symbolCandlesMap[symbol] = []Candle{}
	}

	dataSource.symbolCandlesMap[symbol] = append(
		dataSource.symbolCandlesMap[symbol],
		candle,
	)
}

func (dataSource *DataSource) GetPrevCandle(symbol string, index int) Candle {
	candles := dataSource.GetCandlesFor(symbol)

	return candles[index-1]
}

func (dataSource *DataSource) GetNextCandle(symbol string, index int) (Candle, bool) {
	candles := dataSource.GetCandlesFor(symbol)
	lastIdx := len(candles)
	if lastIdx == index {
		return Candle{}, false
	}

	return candles[index+1], true
}

func (dataSource *DataSource) GetCandlesFor(symbol string) []Candle {
	if candles, ok := dataSource.symbolCandlesMap[symbol]; ok {
		return candles
	}

	return []Candle{}
}

func (dataSource *DataSource) DeleteCandlesFor(symbol string) {
	if _, ok := dataSource.symbolCandlesMap[symbol]; ok {
		delete(dataSource.symbolCandlesMap, symbol)
	}
}

func (dataSource *DataSource) GetLastCandleFor(symbol string) (Candle, bool) {
	if candles, ok := dataSource.symbolCandlesMap[symbol]; ok {
		lastIdx := len(candles) - 1

		return candles[lastIdx], true
	}

	return Candle{}, false
}

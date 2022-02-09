package main

import (
	"github.com/go-gota/gota/dataframe"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

type Candle struct {
	Symbol                   string
	OpenTime                 string
	CloseTime                string
	OpenPrice                float64
	HighPrice                float64
	LowPrice                 float64
	ClosePrice               float64
	Volume                   float64
	QuoteAssetVolume         float64
	NumberOfTrades           int
	TakerBuyBaseAssetVolume  float64
	TakerBuyQuoteAssetVolume float64
	Ignore                   int
}

const OPEN_TIME = 0
const OPEN_PRICE = 1
const HIGH_PRICE = 2
const LOW_PRICE = 3
const CLOSE_PRICE = 4
const VOLUME = 5
const CLOSE_TIME = 6
const QUOTE_ASSET_VOLUME = 7
const NUMBER_OF_TRADES = 8
const TAKER_BUY_BASE_ASSET_VOLUME = 9
const TAKER_BUY_QUOTE_ASSET_VOLUME = 10
const IGNORE = 11

func CsvFileToCandles(fileName string, symbol string) []Candle {
	var candles []Candle

	file, _ := os.Open(fileName)
	csvDataFrame := dataframe.ReadCSV(file)

	for i := 0; i < csvDataFrame.Nrow(); i++ {
		row := csvDataFrame.Subset(i)
		candles = append(candles, CsvRowToCandle(row, symbol))
	}

	return candles
}

func CsvRowToCandle(candleDataFrame dataframe.DataFrame, symbol string) Candle {
	firstRow := 0
	openTime, _ := candleDataFrame.Elem(firstRow, OPEN_TIME).Int()
	closeTime, _ := candleDataFrame.Elem(firstRow, CLOSE_TIME).Int()

	return Candle{
		Symbol:                   symbol,
		OpenTime:                 FormatTimestamp(int64(openTime)),
		OpenPrice:                candleDataFrame.Elem(firstRow, OPEN_PRICE).Float(),
		HighPrice:                candleDataFrame.Elem(firstRow, HIGH_PRICE).Float(),
		LowPrice:                 candleDataFrame.Elem(firstRow, LOW_PRICE).Float(),
		ClosePrice:               candleDataFrame.Elem(firstRow, CLOSE_PRICE).Float(),
		Volume:                   candleDataFrame.Elem(firstRow, VOLUME).Float(),
		CloseTime:                FormatTimestamp(int64(closeTime)),
		QuoteAssetVolume:         candleDataFrame.Elem(firstRow, QUOTE_ASSET_VOLUME).Float(),
		NumberOfTrades:           0,
		TakerBuyBaseAssetVolume:  candleDataFrame.Elem(firstRow, TAKER_BUY_BASE_ASSET_VOLUME).Float(),
		TakerBuyQuoteAssetVolume: candleDataFrame.Elem(firstRow, TAKER_BUY_QUOTE_ASSET_VOLUME).Float(),
		Ignore:                   0,
	}
}

func GetCsvFileNamesInDir(dir string, date string) []string {
	filesNames := []string{}
	files, _ := os.ReadDir(dir)
	path, _ := filepath.Abs(dir)

	for _, file := range files {
		name := file.Name()
		matchedPrefix, _ := regexp.MatchString(`^BTCUSDT-`, name)
		matchedSuffix, _ := regexp.MatchString(`1m-`+date+`\.csv$`, name)
		if !matchedPrefix && matchedSuffix {
			filesNames = append(filesNames, filepath.Join(path, file.Name()))
		}
	}

	return filesNames
}

func GetCoinSymbolFromCsvFileName(fileName string) string {
	compile := regexp.MustCompile(`datasets/(\w+)`)
	groups := compile.FindStringSubmatch(fileName)

	return groups[1]
}

func FormatTimestamp(timestamp int64) string {
	date := ParseMilliTimestamp(timestamp)
	//date = date.Add(time.Hour * 3)

	return date.Format("2006-01-02 15:04:05")
}

func ParseMilliTimestamp(tm int64) time.Time {
	sec := tm / 1000
	msec := tm % 1000
	return time.Unix(sec, msec*int64(time.Millisecond))
}

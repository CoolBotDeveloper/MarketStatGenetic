package main

type Dataset struct {
	AltCoinName    string
	AltCoinCandles []Candle
	BtcCandles     *[]Candle
}

var datasets *[]Dataset

func ImportDatasets() *[]Dataset {
	if datasets == nil {
		dir := "./datasets"
		dates := []string{
			"2021-10",
			"2022-04",
		}

		locDatasets := []Dataset{}
		altCoinCandles := []Candle{}

		for _, date := range dates {
			csvFiles := GetCsvFileNamesInDir(dir, date)

			for _, fileName := range csvFiles {
				symbol := GetCoinSymbolFromCsvFileName(fileName)
				altCoinCandles = append(altCoinCandles, CsvFileToCandles(fileName, symbol)...)
			}
		}

		locDatasets = append(locDatasets, Dataset{
			AltCoinName:    "BTCUSDT",
			AltCoinCandles: altCoinCandles,
			//BtcCandles:     btcDatasetCandles,
		})

		datasets = &locDatasets
	}
	return datasets
}

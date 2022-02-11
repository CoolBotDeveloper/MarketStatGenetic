package main

type Dataset struct {
	AltCoinName    string
	AltCoinCandles []Candle
	BtcCandles     []Candle
}

func ImportDatasets() []Dataset {
	dir := "./datasets"
	dates := []string{
		"2022-01",
	}

	datasets := []Dataset{}
	for _, date := range dates {
		btcFileName := "./datasets/BTCUSDT-1m-" + date + ".csv"
		csvFiles := GetCsvFileNamesInDir(dir, date)

		for _, fileName := range csvFiles {
			symbol := GetCoinSymbolFromCsvFileName(fileName)
			datasets = append(datasets, Dataset{
				AltCoinName:    symbol,
				AltCoinCandles: CsvFileToCandles(fileName, symbol),
				BtcCandles:     CsvFileToCandles(btcFileName, "BTCUSDT"),
			})
		}
	}

	return datasets
}

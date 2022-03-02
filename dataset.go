package main

type Dataset struct {
	AltCoinName    string
	AltCoinCandles []Candle
	BtcCandles     *[]Candle
}

var datasets *[]Dataset
var btcDatasetCandles *[]Candle

func ImportDatasets() *[]Dataset {
	if datasets == nil {
		dir := "./datasets"
		dates := []string{
			"2022-01",
			"2021-11",
			"2021-07",
			"2021-03",
		}
		locDatasets := []Dataset{}
		for _, date := range dates {
			btcFileName := "./datasets/BTCUSDT-1m-" + date + ".csv"
			csvFiles := GetCsvFileNamesInDir(dir, date)

			for _, fileName := range csvFiles {
				symbol := GetCoinSymbolFromCsvFileName(fileName)
				if btcDatasetCandles == nil {
					bd := CsvFileToCandles(btcFileName, "BTCUSDT")
					btcDatasetCandles = &bd
				}

				locDatasets = append(locDatasets, Dataset{
					AltCoinName:    symbol,
					AltCoinCandles: CsvFileToCandles(fileName, symbol),
					BtcCandles:     btcDatasetCandles,
				})
			}
		}
		datasets = &locDatasets
	}
	return datasets
}

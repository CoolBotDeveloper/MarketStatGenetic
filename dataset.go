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
			"2020-01",
			"2020-02",
			"2020-03",
			"2020-04",
			"2020-05",
			"2020-06",
			"2020-07",
			"2020-08",
			"2020-09",
			"2020-10",
			"2020-11",
			"2020-12",
			"2021-01",
			"2021-02",
			"2021-03",
			"2021-04",
			"2021-05",
			"2021-06",
			"2021-07",
			"2021-08",
			"2021-09",
			"2021-10",
			"2021-11",
			"2021-12",
			"2022-01",
			"2022-02",
			"2022-03",
			"2022-04",
			"2022-05",
		}
		locDatasets := []Dataset{}
		for _, date := range dates {
			btcFileName := "./datasets/BTCUSDT-15m-" + date + ".csv"
			csvFiles := GetCsvFileNamesInDir(dir, date)
			bd := CsvFileToCandles(btcFileName, "BTCUSDT")

			for _, fileName := range csvFiles {
				symbol := GetCoinSymbolFromCsvFileName(fileName)
				btcDatasetCandles = &bd

				altCoinCandles := CsvFileToCandles(fileName, symbol)
				if len(*btcDatasetCandles) == len(altCoinCandles) {
					locDatasets = append(locDatasets, Dataset{
						AltCoinName:    symbol,
						AltCoinCandles: altCoinCandles,
						BtcCandles:     btcDatasetCandles,
					})
				}
			}
		}
		datasets = &locDatasets
	}
	return datasets
}

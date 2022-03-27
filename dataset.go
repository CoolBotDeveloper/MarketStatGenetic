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
			//"2022-01",
			//"2022-02",
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

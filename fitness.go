package main

import "fmt"

func Fitness(botConfig BotConfig) float64 {
	totalRevenue := 0.0
	datasets := ImportDatasets()

	for _, dataset := range datasets {
		fmt.Println(dataset.AltCoinName)
		totalRevenue += doBuysAndSells(dataset)
	}

	// todo: import datasets
	// todo: iterate through each alt coin dataset using bitcoin dataset
	// todo: do buys and sells
	// todo: return result

	return totalRevenue
}

func doBuysAndSells(dataset Dataset) float64 {
	return 0
}

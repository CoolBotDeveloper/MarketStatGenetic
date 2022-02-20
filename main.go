package main

import (
	"context"
	"fmt"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/exports"
	"os"
)

func main() {
	bots := GetInitialBots()

	for generation := 0; generation < GENERATION_COUNT; generation++ {
		iterator := bots.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
		for {
			botNumber, bot, _ := iterator()
			if botNumber == nil {
				break
			}
			fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d", generation, *botNumber))
			botConfig := ConvertDataFrameToBotConfig(bot)
			botRevenue := fixRevenue(Fitness(botConfig))
			SetBotTotalRevenue(bots, *botNumber, botRevenue)
			fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d, Revenue: %f\n", generation, *botNumber, botRevenue))
		}
		bots = SortBestBots(bots)
		botsCsvFile, _ := os.Create(fmt.Sprintf("generation_%d.csv", generation))
		exports.ExportToCSV(context.Background(), botsCsvFile, bots)

		bestBots := SelectNBots(BEST_BOTS_COUNT, bots)
		bots = MakeChildren(bestBots)
	}

	fmt.Println("Done")
}

func fixRevenue(revenue float64) float64 {
	if revenue == 0.0 || revenue == -0.15 {
		return DEFAULT_REVENUE
	}
	return revenue
}

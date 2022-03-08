package main

import (
	"context"
	"fmt"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/exports"
	"os"
)

func main() {
	bots := GetInitialBotsFromFile("generation_19.csv")
	//bots := GetInitialBots()
	fitnessDatasets := ImportDatasets()

	for generation := 0; generation < GENERATION_COUNT; generation++ {
		var botRevenueChan = make(chan BotRevenue, 5)

		iterator := bots.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
		for {
			botNumber, bot, _ := iterator()
			if botNumber == nil {
				break
			}

			if *botNumber < BEST_BOTS_FROM_PREV_GEN && generation > 0 {
				rev := convertToFloat64(bot["TotalRevenue"])
				fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d", generation, *botNumber))
				fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d, Revenue: %f\n", generation, *botNumber, rev))
				SetBotTotalRevenue(bots, *botNumber, rev)
				continue
			}

			fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d", generation, *botNumber))
			botConfig := ConvertDataFrameToBotConfig(bot)
			go Fitness(botConfig, *botNumber, botRevenueChan, fitnessDatasets)
		}

		channelsCount := bots.NRows()
		if generation > 0 {
			channelsCount = channelsCount - BEST_BOTS_FROM_PREV_GEN
		}

		for i := 0; i < channelsCount; i++ {
			botRevenue := <-botRevenueChan
			rev := fixRevenue(botRevenue.Revenue)
			SetBotTotalRevenue(bots, botRevenue.BotNumber, rev)
			fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d, Buys Count: %d, Success: %d, Failed: %d, Revenue: %f\n", generation, botRevenue.BotNumber, botRevenue.TotalBuysCount, botRevenue.SuccessBuysCount, botRevenue.FailedBuysCount, rev))
		}
		close(botRevenueChan)

		parentBots := SortBestBots(bots)
		botsCsvFile, _ := os.Create(fmt.Sprintf("generation_%d.csv", generation))
		exports.ExportToCSV(context.Background(), botsCsvFile, parentBots)

		bestBots := SelectNBots(BEST_BOTS_COUNT, parentBots)
		childBots := MakeChildren(bestBots)

		bots = CombineParentAndChildBots(
			SelectNBots(BEST_BOTS_FROM_PREV_GEN, bestBots),
			SelectNBots(BOTS_COUNT-BEST_BOTS_FROM_PREV_GEN, childBots),
		)
	}

	fmt.Println("Done")
}

func fixRevenue(revenue float64) float64 {
	if revenue == 0.0 {
		return DEFAULT_REVENUE
	}
	return revenue
}

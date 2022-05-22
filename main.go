package main

import (
	"context"
	"fmt"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/exports"
	"os"
)

func main() {
	//bots := GetInitialBotsFromFile("initial.csv")
	bots := GetInitialBots()
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
				successPercentage := convertToFloat64(bot["SuccessPercentage"])
				plusRevenue := convertToFloat64(bot["PlusRevenue"])
				minusRevenue := convertToFloat64(bot["MinusRevenue"])

				fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d", generation, *botNumber))
				fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d, Revenue: %f, SuccessPercentage: %f\n", generation, *botNumber, rev, successPercentage))
				SetBotTotalRevenue(bots, *botNumber, rev, successPercentage, plusRevenue, minusRevenue)
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
			successPercentage := CalcSuccessBuysPercentage(botRevenue)
			SetBotTotalRevenue(bots, botRevenue.BotNumber, rev, successPercentage, botRevenue.PlusRevenue, botRevenue.MinusRevenue)
			fmt.Println(fmt.Sprintf("Gen: %d, Bot: %d, Buys Count: %d, Success: %d, Failed: %d, Revenue: %f, SuccessPercentage: %f, Selection: %f\n", generation, botRevenue.BotNumber, botRevenue.TotalBuysCount, botRevenue.SuccessBuysCount, botRevenue.FailedBuysCount, rev, successPercentage, CalcSelection(rev, successPercentage)))
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

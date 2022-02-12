package main

import (
	"fmt"
	"github.com/MaxHalford/eaopt"
)

func main() {
	var ga, err = eaopt.NewDefaultGAConfig().NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set the number of generations to run for
	ga.PopSize = 4
	ga.NGenerations = 1000
	ga.ParallelEval = true

	// Add a custom print function to track progress
	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	// Find the minimum
	err = ga.Minimize(BotSliceFactory)
	if err != nil {
		fmt.Println(err)
		return
	}
}

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/patrikeh/go-deep"
	"github.com/patrikeh/go-deep/training"
)

func testMain() {

	rand.Seed(time.Now().UnixNano())

	dataS, err := load("./normalized_neural_network_dataset.csv")
	if err != nil {
		panic(err)
	}

	for i := range dataS {
		deep.Standardize(dataS[i].Input)
	}
	dataS.Shuffle()

	//split := dataS.SplitN(2)
	data, predict := dataS.Split(0.8)
	//data := split[0]
	//train := split[1]

	neural := deep.NewNeural(&deep.Config{
		Inputs:     len(data[0].Input),
		Layout:     []int{25, 5, 1},
		Activation: deep.ActivationSigmoid,
		Mode:       deep.ModeBinary,
		Weight:     deep.NewNormal(1, 0),
		Bias:       true,
	})

	//trainer := training.NewTrainer(training.NewSGD(0.005, 0.5, 1e-6, true), 50)
	trainer := training.NewBatchTrainer(training.NewSGD(0.005, 0.1, 0, true), 50, 300, 16)
	//trainer := training.NewTrainer(training.NewAdam(0.1, 0, 0, 0), 50)
	//trainer := training.NewBatchTrainer(training.NewAdam(0.1, 0, 0, 0), 50, len(data)/2, 0)
	////data, heldout := data.Split(0.5)
	trainer.Train(neural, data, data, 3000)

	okCount := 0
	fmt.Println(len(predict))
	for idx, _ := range predict {
		fmt.Print(idx)
		r := neural.Predict(predict[idx].Input)
		fmt.Println(r, predict[idx].Response)

		moreThanPercentage := 0.6
		if r[0] >= moreThanPercentage && predict[idx].Response[0] == 1 {
			okCount++
		}

		if r[0] < moreThanPercentage && predict[idx].Response[0] == 0 {
			okCount++
		}
	}

	percentage := float64((okCount * 100.0) / len(predict))
	fmt.Println(fmt.Sprintf("Total: %d, Success: %d, Percentage: %f", len(predict), okCount, percentage))
}

func load(path string) (training.Examples, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(bufio.NewReader(f))

	var examples training.Examples
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if record[0] == "" {
			continue
		}

		examples = append(examples, toExample(record))
	}

	return examples, nil
}

func toExample(in []string) training.Example {
	isGoodCol := len(in) - 1
	res, err := strconv.ParseFloat(in[isGoodCol], 64)
	if err != nil {
		panic(err)
	}

	var features []float64
	for i := 1; i < len(in)-1; i++ {
		res, err := strconv.ParseFloat(in[i], 64)
		if err != nil {
			panic(err)
		}
		features = append(features, res)
	}

	return training.Example{
		Response: []float64{res},
		Input:    features,
	}
}

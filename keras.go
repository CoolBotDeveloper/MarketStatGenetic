package main

import (
	"encoding/csv"
	"fmt"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	"os"
)

func RunKeras() {
	// replace myModel and myTag with the appropriate exported names in the chestrays-keras-binary-classification.ipynb
	model, err := tf.LoadSavedModel("test_model", []string{"serve"}, nil)

	if err != nil {
		fmt.Printf("Error loading saved model: %s\n", err.Error())
		return
	}

	defer model.Session.Close()

	r := ImportDataset()

	result, err := model.Session.Run(
		map[tf.Output]*tf.Tensor{
			model.Graph.Operation("inputLayer_input").Output(0): r[0], // Replace this with your input layer name
		},
		[]tf.Output{
			model.Graph.Operation("inferenceLayer/Sigmoid").Output(0), // Replace this with your output layer name
		},
		nil,
	)

	if err != nil {
		fmt.Printf("Error running the session with input, err: %s\n", err.Error())
		return
	}

	fmt.Printf("Result value: %v \n", result[0].Value())
}

func ImportDataset() []*tf.Tensor {
	fileName := "perfect_5_minutes.csv"
	file, err := os.Open(fileName)
	if err != nil {
		panic("Can not load initial bots from file.")
	}

	csvReader := csv.NewReader(file)
	rows, err := csvReader.ReadAll()

	tensors := []*tf.Tensor{}

	for index, candles := range rows {
		if index == 0 {
			continue
		}

		beforeCount := 3
		afterCount := 1
		candlesCount := 100
		datum := []float64{}

		for dataIndex, candle := range candles {
			afterIndex := beforeCount + candlesCount*2 + afterCount
			if dataIndex < beforeCount || dataIndex > afterIndex {
				continue
			}

			datum = append(datum, convertStringToFloat64(candle))
		}

		newTensor, err := tf.NewTensor(datum)
		if err != nil {
			fmt.Println("Error tensor")
		}

		tensors = append(tensors, newTensor)
	}

	return tensors
}

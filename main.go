package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Weight struct {
	Conv2D    []float64
	W0        []float64
	W1        []float64
	BatchNorm []float64
}

type Data struct {
	Weights []Weight
}

func Load() Data {
	jsonFile, err := os.Open("weights.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened weights.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	json.Unmarshal([]byte(byteValue), &data)
	return data
}

func main() {

	data := Load()

	dataFile, err_file_in := os.Create("integerdata.gob")
	if err_file_in != nil {
		fmt.Println(err_file_in)
		os.Exit(1)
	}
	enc := gob.NewEncoder(dataFile) // Will write to network.
	err := enc.Encode(data)
	if err != nil {
		log.Fatal("encode error:", err)
	}
	dataFile.Close()
}

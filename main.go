package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
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
	// jsonFile, err := openuri.Open("http://localhost:3000/weights.json")
	jsonFile, err := os.Open("weights.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var data Data
	json.Unmarshal([]byte(byteValue), &data)
	return data
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {

	data := Load()
	// log.Println(data)

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

func writeFile(bytes []byte) int {
	file, _ := os.Create("test.bin")
	defer file.Close()

	number_byte := writeNextBytes(file, bytes)
	return number_byte
}

func readFile(number_byte int) []byte {
	file, _ := os.Open("test.bin")
	defer file.Close()

	bytes_array := readNextBytes(file, number_byte)
	return bytes_array
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}

// func writeFile() {
// 	file, err := os.Create("test.bin")
// 	defer file.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	r := Load()

// 	for i := 0; i < 10; i++ {

// 		s := r[i]
// 		var bin_buf bytes.Buffer
// 		binary.Write(&bin_buf, binary.BigEndian, s)
// 		//b :=bin\_buf.Bytes()
// 		//l := len(b)
// 		//fmt.Println(l)
// 		writeNextBytes(file, bin_buf.Bytes())
// 	}
// }

func writeNextBytes(file *os.File, bytes []byte) int {

	_, err := file.Write(bytes)

	if err != nil {
		log.Fatal(err)
	}

	return len(bytes)
}

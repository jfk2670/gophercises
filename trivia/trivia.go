package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"io"
	"log"
	"time"
	"strconv"
)

const PROBLEMS = "/Users/jack/Downloads/_git/gophercises/ex1/problems.csv"

var correct int
var incorrect int

func askUser(line []string) {
	fmt.Print(line[0], ": ")

	var response string
	_, error := fmt.Scanln(&response)
	if error != nil {
		log.Fatal(error)
		return
	}
	if response == line[1] {
		correct += 1
	} else {
		incorrect += 1
	}
}

func checkTime(start int64, timeLimit int64) {
	now := time.Now()
	cur := now.Unix()
	if cur - start > timeLimit {
		fmt.Println("You did not finish the quiz in time. Better luck next time!")
		printResults()
		os.Exit(0)
	}
	return
}

func printResults() {
	fmt.Println("RESULTS:")
	fmt.Println("Total correct:", correct)
	fmt.Println("Total incorrect:", incorrect)
}

func main () {
	var timeLimit int64 = 30

	if len(os.Args) == 2 {
		timeLimit,_ = strconv.ParseInt(os.Args[1], 10, 64)
	}

	file, error := os.Open(PROBLEMS)

	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()
	defer printResults()

	reader := csv.NewReader(file)

	fmt.Printf("You will have %vs to complete the quiz.\nPress ENTER to begin", timeLimit)
	var enter string
	fmt.Scanln(&enter)

	now := time.Now()
	start := now.Unix()

	for {
		line, error := reader.Read()
		if error == io.EOF{
			break
		}
		checkTime(start, timeLimit)
		askUser(line)
	}
	
}
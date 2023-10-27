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

const PROBLEMS = "/Users/jack/Downloads/git/gophercises/ex1/problems.csv"

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

func finishedInTime(start int64, timeLimit int64) bool {
	now := time.Now()
	stop := now.Unix()
	if stop - start < timeLimit {
		return true
	} else {
		return false
	}
}

func main () {
	var timeLimit int64 = 30

	if len(os.Args) == 2 {
		timeLimit,_ = strconv.ParseInt(os.Args[1], 10, 64)
	}

	now := time.Now()
	start := now.Unix()

	file, error := os.Open(PROBLEMS)

	if error != nil {
		log.Fatal(error)
	}
	defer file.Close()
	
	reader := csv.NewReader(file)
	
	for {
		line, error := reader.Read()
		if error == io.EOF{
			break
		}
		//fmt.Println(line)
		askUser(line)
	}
	
	if finishedInTime(start, timeLimit) {
		fmt.Println("The quiz is over! Here are your results:")
		fmt.Println("Total correct:", correct)
		fmt.Println("Total incorrect:", incorrect)
	} else {
		fmt.Println("You did not finish the quiz in time. Better luck next time!")
	}
}
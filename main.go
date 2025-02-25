/*
Name: Zachary Perry
Date: 2-25-25
Class: COSC 581

Lab 1: Median of Medians Quicksort Assignment

main.go - Contains driver code for reading in the file and calling quicksort. Also supports the '-benchmark' cmd line arg to run benchmark tests.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// readFile will just read the numbers in from the file.
// This assumes that the file contains a single number on each line (seperated by newlines).
func readFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error encounter trying to open the file..")
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Error converting str to int...")
			os.Exit(1)
		}

		numbers = append(numbers, num)
	}
	return numbers
}

// printOutput just does exactly what is says -> prints the output
func printOutput(sortedNumbers []int) {
	for i := 0; i < len(sortedNumbers); i++ {
		fmt.Print(sortedNumbers[i], " ")
	}

	fmt.Print("\n")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./bin/main filename [-benchmark]")
		os.Exit(1)
	}

	// If benchmark enabled and file provided, run benchmark using the file as the input
	// Otherwise, run the benchmark using random arrays of 1 million integers
	if len(os.Args) == 3 && os.Args[2] == "-benchmark" {
		numbers := readFile(os.Args[1])
		benchmarkOnFile(numbers)
		return
	} else if len(os.Args) == 2 && os.Args[1] == "-benchmark" {
		benchmark()
		return
	}

	// Defaulting to an r of 5 here (benchmark tests 3,5,7,9,11)
	r := 5
	numbers := readFile(os.Args[1])
	sortedNumbers := quicksort(numbers, 0, len(numbers)-1, r)
	printOutput(sortedNumbers)
}

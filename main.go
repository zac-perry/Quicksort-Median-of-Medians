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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./bin/main filename [-benchmark]")
		os.Exit(1)
	}

	// Open the file and read the contents
	file, err := os.Open(os.Args[1])
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

	// if benchmark enabled, run that
	if len(os.Args) == 3 && os.Args[2] == "-benchmark" {
		benchmark(numbers)
		return
	}

	r := 5
	sortedNumbers := quicksort(numbers, 0, len(numbers)-1, r)
	printOutput(sortedNumbers)
}

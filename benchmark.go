/*
Name: Zachary Perry
Date: 2-25-25
Class: COSC 581

Lab 1: Median of Medians Quicksort Assignment

benchmark.go - contains code for printing results + benchmarking performance for various values of r. It will also generate plots for the results.
*/

package main

import (
	"fmt"
	"sort"
	"time"
)

/*
printOutput will just print out the sorted numbers in a semi-readable fashion
*/
func printOutput(sortedNumbers []int) {
	for i := 0; i < len(sortedNumbers); i++ {
		fmt.Print(sortedNumbers[i], " ")
	}

	fmt.Print("\n")
}

// TODO: Plotting the results for my report
func plotResults() {}

// TODO: Generate random array of 1 mil elements to test with.
// Will end up running the benchmark multiple times
func generateTestArray() {}

// TODO: Average multiple runs for each r using various 1mil element arrays
func averageResults() {}

/*
benchmark function will run quicksort with various values of r.
TODO: run on random 1 mil arrays 10 times for each value of r.
TODO: run on the input specified as well i suppose.
*/
func benchmark(numbers []int) {
	fmt.Println("===========================================")
	fmt.Println("         RUNNING IN BENCHMARK MODE")
	fmt.Println("===========================================")

	r := 3
	for r < 13 {
		copyNumbers := make([]int, len(numbers))
		sortedCopy := make([]int, len(numbers))
		copy(copyNumbers, numbers)
		copy(sortedCopy, numbers)
		sort.Ints(sortedCopy)

		startingTime := time.Now()
		quicksort(copyNumbers, 0, len(copyNumbers)-1, r)
		finalTime := time.Since(startingTime)

		fmt.Printf(" r = %3d  | Final Time: %6d ns\n", r, finalTime.Nanoseconds())
		r += 2

		for i := 0; i < len(sortedCopy); i++ {
			if sortedCopy[i] != copyNumbers[i] {
				fmt.Println("ERROR ------ DID NOT SORT CORRECTLY!")
			}
		}
	}
}

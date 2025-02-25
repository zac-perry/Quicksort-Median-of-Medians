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
	"math/rand"
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

// TODO: Generate random array of 1 mil elements to test with.
// Will end up running the benchmark multiple times
func generateTestArray() []int {
	numbers := make([]int, 1_000_000)

	for i := 0; i < 1_000_000; i++ {
		numbers[i] = rand.Int()
	}

	return numbers
}

// TODO: Plotting the results for my report
func plotResults() {}

/*
averageResults just averages all times recorded for each r value during the benchmark.
 */
func averageResults(data []float64) float64 {
	sum := 0.0

	for i := 0; i < len(data); i++ {
		sum += data[i]
	}

	return sum / float64(len(data))
}

/*
benchmark function will run quicksort with various values of r on 10 different, random arrays of 1 million elements.
*/
func benchmark() {
	data := make(map[int][]float64)

	for i := 0; i < 10; i++ {
		fmt.Println("\n===========================================")
		fmt.Println("         RUNNING IN BENCHMARK MODE")
		fmt.Println("                 RUN #", i)
		fmt.Println("===========================================")

		r := 3
		numbers := generateTestArray()
		for r < 13 {
			copyNumbers := make([]int, len(numbers))
			sortedCopy := make([]int, len(numbers))
			copy(copyNumbers, numbers)
			copy(sortedCopy, numbers)
			sort.Ints(sortedCopy)

			startingTime := time.Now()
			quicksort(copyNumbers, 0, len(copyNumbers)-1, r)
			finalTime := time.Since(startingTime).Seconds()

			fmt.Printf(" r = %3d  | Final Time: %6.10f s\n", r, finalTime)
			data[r] = append(data[r], finalTime)

			r += 2

			for i := 0; i < len(sortedCopy); i++ {
				if sortedCopy[i] != copyNumbers[i] {
					fmt.Println("ERROR ------ DID NOT SORT CORRECTLY!")
				}
			}
		}
	}
	fmt.Println("\n\n===========================================")
	fmt.Println("        AVERAGE TIMES FOR EACH R")
	fmt.Println("===========================================")

	for k, v := range data {
		average := averageResults(v)
		fmt.Printf(" r = %3d  | Average Time: %6.10f s\n", k, average)
	}
}

/*
benchmarkOnFile will take the file as input and run the sort using different r values.
*/
func benchmarkOnFile(numbers []int) {
	fmt.Println("===========================================")
	fmt.Println("         RUNNING IN BENCHMARK MODE")
	fmt.Println("        USING PROVIDED FILE AS INPUT")
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

		fmt.Printf(" r = %3d  | Final Time: %6.10f s\n", r, finalTime.Seconds())
		r += 2

		for i := 0; i < len(sortedCopy); i++ {
			if sortedCopy[i] != copyNumbers[i] {
				fmt.Println("ERROR ------ DID NOT SORT CORRECTLY!")
			}
		}
	}
}

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
	"os"
	"sort"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// generateTestArray will generate a slice of 1,000,000 random integers
func generateTestArray() []int {
	numbers := make([]int, 1_000_000)

	for i := 0; i < 1_000_000; i++ {
		numbers[i] = rand.Int()
	}

	return numbers
}

// plotResults will plot r vs time (in seconds)
// note: this uses an external library to create the graphs
func plotResults(data map[int]float64, title string, fileName string) {
	barChart := charts.NewBar()
	barChart.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    title,
		Subtitle: "by Zac Perry",
	}),
		charts.WithXAxisOpts(opts.XAxis{
			Name:         "r values",
			NameLocation: "middle",
			NameGap:      40,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:         "Time (in seconds)",
			NameLocation: "middle",
			NameGap:      70,
		}))

	keys := []int{}
	values := []opts.BarData{}
	for k := range data {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		values = append(values, opts.BarData{Value: data[k]})
	}

	barChart.SetXAxis(keys).AddSeries("Times", values)

	f, _ := os.Create(fileName)
	barChart.Render(f)
}

// averageResults just averages all times recorded for each r value during the benchmark.
func averageResults(data []float64) float64 {
	sum := 0.0

	for i := 0; i < len(data); i++ {
		sum += data[i]
	}

	return sum / float64(len(data))
}

// benchmark function will run quicksort with various values of r on 10 different, random arrays of 1 million elements.
// It will calculate the average time to sort for each value of r over the 10 runs.
// It will then plot the results and save the image to the images/ dir.
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

			startingTime := time.Now()
			quicksort(copyNumbers, 0, len(copyNumbers)-1, r)
			finalTime := time.Since(startingTime).Seconds()

			fmt.Printf(" r = %3d  | Final Time: %6.10f s\n", r, finalTime)
			data[r] = append(data[r], finalTime)

			sort.Ints(sortedCopy)
			checkSort(copyNumbers, sortedCopy)

			r += 2
		}
	}
	fmt.Println("\n\n===========================================")
	fmt.Println("        AVERAGE TIMES FOR EACH R")
	fmt.Println("===========================================")

	avgData := make(map[int]float64)
	for k, v := range data {
		average := averageResults(v)
		avgData[k] = average
		fmt.Printf(" r = %3d  | Average Time: %6.10f s\n", k, average)
	}

	plotResults(avgData, "Average Sorting Times for Each R", "./images/avg_bar_char.html")
}

/*
benchmarkOnFile will take the file as input and run the sort using different r values.
It will then plot the results and save the image to the images/ dir.
*/
func benchmarkOnFile(numbers []int) {
	fmt.Println("===========================================")
	fmt.Println("         RUNNING IN BENCHMARK MODE")
	fmt.Println("        USING PROVIDED FILE AS INPUT")
	fmt.Println("===========================================")
	data := make(map[int]float64)
	r := 3

	for r < 13 {
		copyNumbers := make([]int, len(numbers))
		sortedCopy := make([]int, len(numbers))
		copy(copyNumbers, numbers)
		copy(sortedCopy, numbers)

		startingTime := time.Now()
		quicksort(copyNumbers, 0, len(copyNumbers)-1, r)
		finalTime := time.Since(startingTime).Seconds()

		fmt.Printf(" r = %3d  | Final Time: %6.10f s\n", r, finalTime)
		data[r] = finalTime

		sort.Ints(sortedCopy)
		checkSort(copyNumbers, sortedCopy)

		r += 2
	}

	plotResults(data, "Sorting Time for Each R for Input File", "./images/input_bar_char.html")
}

// checkSort will just ensure that my sorting algorithm actually worked.
func checkSort(mine []int, correct []int) {
	for i := 0; i < len(correct); i++ {
		if correct[i] != mine[i] {
			fmt.Println("ERROR ------ DID NOT SORT CORRECTLY!")
			return
		}
	}
}

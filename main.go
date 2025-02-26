/*
Name: Zachary Perry
Date: 2-25-25
Class: COSC 581

Lab 1: Median of Medians Quicksort Assignment

main.go
- Contains driver code for reading in the file and calling quicksort.
- Supports the '-benchmark' cmd line arg to run benchmark tests.
- Additionally, contains implementation for quicksort & median of medians algorithm.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// FindSubMedian will recursively find the median of the subarray passed to it.
// First, it will partition the subarray to get the pivotIndex.
// Partially sorts the array by partitioning portions that could contain the median value.
// Once the pivotIndex used == the medianIndex(len/2) -> median element has been found.
// sources: Textbook (Chapter 9), https://en.wikipedia.org/wiki/Quickselect.
func findSubMedian(numbers []int, low int, high int, medianIndex int) int {
	// If the low and high are equal, then we can just return the number at this index as it is the median
	if low == high {
		return numbers[low]
	}

	// Partition the subarray using a starting pivot index at index low
	pivotIndex := low
	pivotIndex = partition(numbers, low, high, pivotIndex)

	// Will recursivley call + partition until we find the index at which the median will exist
	// If the medianIndex is less than the pivotIndex, search the left side. Otherwise, search right.
	if medianIndex == pivotIndex {
		return numbers[medianIndex]
	} else if medianIndex < pivotIndex {
		return findSubMedian(numbers, low, pivotIndex-1, medianIndex)
	}
	return findSubMedian(numbers, pivotIndex+1, high, medianIndex)
}

// MedianofMedians will find the median of all medians in the array
// It does this by looping over the number of subarrays based on r.
// It will create subarrays, find the median of these subarrays, and store those values.
// Finally, it will recursively call the function to continue and find the median of all medians.
// sources: Textbook (Chapter 9), https://en.wikipedia.org/wiki/Median_of_medians
func medianOfMedians(numbers []int, r int) int {
	length := len(numbers)

	// For the numbers array with current length <= r, manually sort and return median of this subarray
	if length <= r {
		copyNumbers := make([]int, length)
		copy(copyNumbers, numbers)
		return findSubMedian(copyNumbers, 0, length-1, length/2)
	}

	numSubArrays := (length + r - 1) / r
	medians := make([]int, numSubArrays)

	// Loop through and calculate new indicies for each subarray.
	// Then, for that subarray, find the median and store into the 'medians' array
	for i := 0; i < numSubArrays; i++ {
		newLow := i * r
		newHigh := newLow + r
		if newHigh > length {
			newHigh = length
		}

		subgroup := numbers[newLow:newHigh]
		medians[i] = findSubMedian(subgroup, 0, len(subgroup)-1, len(subgroup)/2)
	}

	return medianOfMedians(medians, r)
}

// Partition will use the given pivotIndex and separate out the array such that all elements <
// the item at the pivot index are on the left and the elements > are on the right side.
// partiton will return the index of the pivot item.
func partition(numbers []int, low int, high int, pivotIndex int) int {
	// Get pivot and move to the back
	pivot := numbers[pivotIndex]
	numbers[pivotIndex], numbers[high] = numbers[high], numbers[pivotIndex]

	// Swapping
	i := low
	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
		}
	}

	// Move the pivot back
	numbers[i], numbers[high] = numbers[high], numbers[i]
	return i
}

// Quicksort algorithm using median of medians to find the pivot.
// Will find the pivot + pivot index using median of medians.
// Then, it will quicksort using this as the starting pivotIndex.
// Partition -> Quicksort(left side) -> Quicksort (right side).
func quicksort(numbers []int, low int, high int, r int) []int {
	if low < high {
		// copy array to use for medianOfMedians
		copyNumbers := make([]int, high-low)
		copy(copyNumbers, numbers[low:high])

		// Find the median of medians to use as the starting pivot
		pivot := medianOfMedians(copyNumbers, r)
		pivotIndex := low

		// Find the pivot index (median index) to use
		for i := low; i <= high; i++ {
			if numbers[i] == pivot {
				pivotIndex = i
				break
			}
		}

		// Regular quicksorting
		pivotIndex = partition(numbers, low, high, pivotIndex)
		numbers = quicksort(numbers, low, pivotIndex-1, r)
		numbers = quicksort(numbers, pivotIndex+1, high, r)
	}

	return numbers
}

// ReadFile will just read the numbers in from the file.
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

// PrintOutput just prints the sortedNumbers to stdout.
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

	// Defaulting to an r of 7 here (benchmark tests 3,5,7,9,11)
	r := 7
	numbers := readFile(os.Args[1])
	sortedNumbers := quicksort(numbers, 0, len(numbers)-1, r)
	printOutput(sortedNumbers)
}

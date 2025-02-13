// Zachary Perry
// 2/20/25
// Median of medians pivot selection for quick sort

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// TODO - median of median algo to select a good pivot value
// general idea:
// - Divide the arr into len(arr)/r sub arrays
// - find the median in each
// - then, take the median of these median values to use as the pivot?
// r is hte number of elements in each sub array
// consider returning the index of the median value found to make life easier
func medianOfMedians(numbers []int, r int) int {

	// if the current length <= r, then go ahead and solve, return
	// otherwise:

	medians := []int{}
	medianIndex := []int{}
	numSubArrays := len(numbers) / r

	fmt.Println(medians, numSubArrays, medianIndex)

	// loop over this, get the median of the elements (helper function)
	// ever r elements, anything left over, solve seperetly
	// account for any elements that dont fall into the r!
	// make sure to add to the medians slice

	// then, once i have all medians, recursive call to medians
	pivot := medianOfMedians(medians, r)

	return pivot
}

// calculation the pivot index
// Then, call quicksort on the 'left' lower elements
// Then, call quicksort on the 'right' lower elements
func quicksort(numbers []int, low int, high int) []int {
	if low < high {
		pivot, numbers := partition(numbers, low, high)
		numbers = quicksort(numbers, low, pivot-1)
		numbers = quicksort(numbers, pivot+1, high)
	}
	return numbers
}

// partition the data, return the pivot index & numbers slice
// Take the rightmost value as the pivot
// From low to high, if the number is <= pivot --> swap value positions
// Then, swap the final elements at i and high indexes, return the pivot index and modified slice
func partition(numbers []int, low int, high int) (int, []int) {
	pivot := numbers[high]
	// here, call the median of medians function i guess

	i := low
	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
		}
	}

	numbers[i], numbers[high] = numbers[high], numbers[i]

	return i, numbers
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./bin/main filename")
		os.Exit(1)
	}

	// Open the file
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error encounter trying to open the file..")
		os.Exit(1)
	}
	defer file.Close()

	// read in the file contents (assumes input is a singular line in the file)
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

	fmt.Println("Numbers before sorting: ", numbers)
	fmt.Println("Numbers after sorting : ", quicksort(numbers, 0, len(numbers)-1))
}

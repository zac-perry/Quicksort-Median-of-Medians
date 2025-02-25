/*
Name: Zachary Perry
Date: 2-25-25
Class: COSC 581

Lab 1: Median of Medians Quicksort Assignment

quicksort.go - contains my implementation for quicksort w/ the median of medians algorithm for pivot selection.
*/

package main

/*
findSubMedian will manually (recursively) find the median of the subarray passed to it
*/
func findSubMedian(numbers []int, low int, high int, medianIndex int) int {
	if low == high {
		return numbers[low]
	}

	pivotIndex := low
	pivotIndex = partition(numbers, low, high, pivotIndex)

	// will keep partitioning until we find the partition index where the median exists
	if medianIndex == pivotIndex {
		return numbers[medianIndex]
	} else if medianIndex < pivotIndex {
		return findSubMedian(numbers, low, pivotIndex-1, medianIndex)
	} else {
		return findSubMedian(numbers, pivotIndex+1, high, medianIndex)
	}
}

/*
medianOfMedians will find the median of all medians in the array
*/
func medianOfMedians(numbers []int, r int) int {
	length := len(numbers)

	// for array with current length <= r, manually sort and return median of this subarray
	if length <= r {
		tmp := make([]int, length)
		copy(tmp, numbers)
		return findSubMedian(tmp, 0, length-1, length/2)
	}

	numSubArrays := (length + r - 1) / r
	medians := make([]int, numSubArrays)

	// otherwise, create the sub arrays, find the median of the subarray, and store
	for i := 0; i < numSubArrays; i++ {
		// Determine the start and end of this group
		// getting the current start + end for the current subarray
		newLow := i * r
		newEnd := newLow + r
		if newEnd > length {
			newEnd = length
		}

		// Extract the current group
		group := numbers[newLow:newEnd]

		// insert into the medians array to recursivley find the median of sub medians
		medians[i] = findSubMedian(group, 0, len(group)-1, len(group)/2)
	}

	// recursivley find the median of all medians found
	return medianOfMedians(medians, r)
}

// partition the data, return the pivot index
// Take the current median value as the pivot
// From low to high, if the number is <= pivot --> swap value positions
// Then, swap the final elements at i and high indexes, return the pivot index
func partition(numbers []int, low int, high int, pivotIndex int) int {
	pivot := numbers[pivotIndex]
	numbers[pivotIndex], numbers[high] = numbers[high], numbers[pivotIndex]

	i := low
	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			numbers[i], numbers[j] = numbers[j], numbers[i]
			i++
		}
	}

	numbers[i], numbers[high] = numbers[high], numbers[i]
	return i
}

/*
  Quicksort using the median of medians algorithm to find the best pivot
    - Will find the median of medians & its pivot index value
    - Then, it uses this for partitioning
*/
// Finding the pivot index of the median of medians
// Then, call quicksort on the 'left' lower elements
// Then, call quicksort on the 'right' lower elements
func quicksort(numbers []int, low int, high int, r int) []int {
	if low < high {
		subArray := make([]int, high-low+1)
		copy(subArray, numbers[low:high+1])

		// find the median of medians to use as the starting pivot
		pivot := medianOfMedians(subArray, r)
		pivotIndex := low

		// find the pivot index (median index) to use
		for i := low; i <= high; i++ {
			if numbers[i] == pivot {
				pivotIndex = i
				break
			}
		}
		pivotIndex = partition(numbers, low, high, pivotIndex)
		numbers = quicksort(numbers, low, pivotIndex-1, r)
		numbers = quicksort(numbers, pivotIndex+1, high, r)
	}
	return numbers
}

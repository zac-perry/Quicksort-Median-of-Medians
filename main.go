package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// TODO
func quickSort() {}

// TODO
func partition() {}

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
}

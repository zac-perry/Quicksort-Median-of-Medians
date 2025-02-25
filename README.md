# QuickSort with Median of Medians
Quick Sort implementation using the median-of-medians pivot selection algorithm.

# TODO: 
- [x] Makefile, general setup
- [x] Research more about the median-of-medians method
- [x] Basic implementation
- [x] Implementation with median-of-medians
- [ ] Update README with details, how to run, etc.
- [ ] Finish up comments (provided sources used, etc.)
- [ ] Test on the hydra/tesla machines (before submission)
- [ ] Writeup (include graphs, method, etc)
- [x] Probably best to merge the qucksort file with main for simplicity
- [ ] Potentially change the manual median selection for the subarray?
- [x] consider adding benchmark tests? 
- [x] finish benchmark setup
  ~~Involves adding randomized arrays~~
  ~~Involves saving the data~~
  ~~Involves using this data and plotting~~

## Implementation Details
Quick sort, pivot selection method: median of medians with some value 
    - r represents the number of elements in each subarray

## How to Run
Note: this requires go version 1.22.9. This has to be ran on the hydra/tesla machines and this is the version that they have installed currently.

### To compile
To compile the program, just run ` make `. This will compile all three files used for this.
To remove binary files, run `make clean`.

### To run quicksort on a specified file
` ./bin/main inputFile [-benchmark]`
Two things to note here. It assumes that the input file contains a single number on each line (newline seperated).
Additionally, the benchmark argument is optional.

If you just want to see the sorted output: `./bin/main inputFile`

NOTE: in the code, the r value is defaulted to 5.

### Benchmarking
To benchmark performance, i time the amount of time it takes the program to sort the given input for each value of r (3, 5, 7, 9, 11). Then, I create a bar chart displaying the results. This is saved into an HTML file within the `images/` directory.

If you want to evaluate the performance of the program on the given input file (with each value of r): `./bin/main inputFile -benchmark`

If you want to evaluate the performance of the program in general (with each value of r): `./bin/main -benchmark`
This will run the program 10 different times with 10 different arrays containing 1,000,000 random elements. It will evaluate each value of r for each run. 

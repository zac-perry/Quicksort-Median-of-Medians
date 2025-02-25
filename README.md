# CS581-QuickSort
Quick Sort implementation using the median-of-medians pivot selection algorithm

# TODO: 
- [x] Makefile, general setup
- [x] Research more about the median-of-medians method
- [x] Basic implementation
- [x] Implementation with median-of-medians
- [ ] Update README with details, how to run, etc.
- [ ] Test on the hydra/tesla machines
- [x] consider adding benchmark tests? 
- [ ] finish benchmark setup
    - Involves adding randomized arrays
    - Involves saving the data
    - Involves using this data and plotting

## Implementation Details
Quick sort, pivot selection method: median of medians with some value ----------
    - r represents the number of elements in each subarray

## How to Run
note: this requires go version 1.22.9. This has to be ran on the hydra/tesla machines and this is the version that they have installed currently.

#### To compile
` make `
will compile all three go files being used here

#### To run quicksort on a specified file
` ./bin/main inputFile [-benchmark]`
optional benchmark parameter if you want to see the performance for various values of r on 1 million random elements


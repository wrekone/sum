package main

// Sum adds all numbers in a given slice
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll returns a sum of slices
func SumAll(numbersToSum ...[]int) (sums []int) {
	return
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Numbers struct {
	Nums []int
}

// Average
func (n *Numbers) Average() float64 {
	if len(n.Nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range n.Nums {
		sum += num
	}
	return float64(sum) / float64(len(n.Nums))
}

// Median
func (n *Numbers) Median() int {
	if len(n.Nums) == 0 {
		return 0
	}
	sort.Ints(n.Nums)
	mid := len(n.Nums) / 2
	if len(n.Nums)%2 != 0 {
		return n.Nums[mid]
	}
	med := float64(n.Nums[mid-1]+n.Nums[mid]) / 2
	return int(math.Round(med))
}

// Variance
func (n *Numbers) Variance() float64 {
	if len(n.Nums) == 0 {
		return 0
	}
	average := n.Average()
	sumOfSquares := 0.0
	for _, num := range n.Nums {
		diff := float64(num) - average
		sumOfSquares += diff * diff
	}
	return sumOfSquares / float64(len(n.Nums))
}

// StandardDeviation
func (n *Numbers) StandardDeviation() float64 {
	return math.Sqrt(n.Variance())
}

// Guess method
func (n *Numbers) Guess() (int, int) {
	var max, min float64
	start := len(n.Nums) - 4
	if start < 0 {
		start = 0
	}

	// Create a new Numbers instance for the last up to 4 elements
	newSl := n.Nums[start:]
	newNumbers := &Numbers{Nums: newSl}

	avg := newNumbers.Average()
	stdev := newNumbers.StandardDeviation()

	max = avg + stdev
	min = avg - stdev

	if min < 0 {
		min = 0 // Ensure min is not negative
	}
	return int(min), int(max)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums []int
	for scanner.Scan() {
		conres, err := strconv.Atoi(scanner.Text())
		if err != nil {
			continue
		}
		nums = append(nums, conres)
		if len(nums) <= 1 {
			continue
		}
		numbers := &Numbers{Nums: nums} // Create a Numbers instance
		min, max := numbers.Guess()     // Call Guess on the instance
		fmt.Println(min, max)
	}
}

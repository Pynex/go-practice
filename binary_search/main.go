package main

import (
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"sort"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)

var verbosity = flag.Bool("verbosity", false, "enable print array and target. Default is false")

func main() {
	flag.Parse()
	fmt.Println("starting binary search...")
	max := 1000
	randomLen := rand.Intn(max)
	nums := make([]int, randomLen)

	for i := 0; i < len(nums); i++ {
		max := 1000000
		randomNum := rand.Intn(max)
		nums[i] = randomNum
	}

	randomIndexTarget := rand.Intn(randomLen)
	val := nums[randomIndexTarget]

	sort.Ints(nums)

	fmt.Println("target value: ", val)
	if *verbosity {
		printArrayWithHighlight(nums, val)
	}

	result, err := binary_search(nums, val)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Found at: ", ColorGreen, result, ColorReset)
}

func printArrayWithHighlight(arr []int, target int) {
	fmt.Print("Array: [")
	for i, num := range arr {
		if num == target {
			fmt.Printf("%s%d%s", ColorGreen, num, ColorReset)
		} else {
			fmt.Printf("%d", num)
		}

		if i < len(arr)-1 {
			fmt.Print(", ")
		}

		if i > 0 && i%20 == 0 {
			fmt.Println()
			fmt.Print("       ")
		}
	}
	fmt.Println("]")
}

func binary_search(array []int, target int) (index int, err error) {
	if len(array) == 0 {
		return -1, errors.New("array is empty")
	}
	left, right := 0, len(array)-1

	for left <= right {
		mid := left + (right-left)/2
		mid_value := array[mid]

		if mid_value == target {
			return mid, nil
		} else if mid_value < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1, errors.New("target element not found")
}

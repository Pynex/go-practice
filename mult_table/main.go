package main

import (
	"flag"
	"fmt"
)

//TODO: 1)modify this code to optional print results
// 			done!
//		2)give user option to choose last table element
//			done!

var res = flag.Bool("res", false, "enable result mode")
var amount = flag.Int("amount", 10, "choose amount of tables. Default is 9")

func main() {
	flag.Parse()
	*amount++

	nums := make([]int, *amount)

	for i := 2; i < *amount; i++ {
		nums[i] = i
	}

	var results []int

	fmt.Print("a\\b\t")

	for _, num := range nums {
		fmt.Printf("%6d", num)
	}

	fmt.Print("\n\n")

	for _, a := range nums {
		fmt.Printf("%d\t", a)

		for _, b := range nums {
			res := a * b
			fmt.Printf("%6d", res)

			results = append(results, res)
		}

		fmt.Println()
	}

	if *res {
		fmt.Println("\n\tResults:")
		for index, result := range results {
			fmt.Printf("%6d\t%6d\n", index, result)
		}
	}
}

package main

import(
	"fmt"
)
//TODO: 1)modify this code to optional print results
//		2)give user option to choose last table element
func main() {
	nums := [...]int{2, 3, 4, 5, 6, 7, 8, 9, 10}

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

	fmt.Println("\n\tResults:")
	for index, result := range results {
		fmt.Printf("%6d\t%6d\n",index, result)
	}
}
// see: https://gobyexample.com/range
package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	// ignore the index
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum) //=> 9

	// Using the index
	for i, num := range nums {
		if num == 3 {
			fmt.Printf("Found 3 at index %d\n", i) //=> Found 3 at index 1
		}
	}

	// Using range with map
	pairs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range pairs {
		fmt.Printf("%s -> %s\n", k, v) //=> a -> apple, b -> banana
	}
	// Using range with string
	for i, c := range "go" {
		fmt.Println(i, c) //=> 0 101
		//=> 1 111
	}
}

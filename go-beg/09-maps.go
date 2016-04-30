// see: https://gobyexample.com/maps
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m) //=> map[k1:7 k2:13]

	v1 := m["k1"]
	fmt.Println(v1) //=> 7

	fmt.Println(len(m)) //=> 2

	delete(m, "k2")
	fmt.Println(m)      //=> map[k1:7]
	fmt.Println(len(m)) //=> 1

	_, prs := m["k2"]
	fmt.Println("prs", prs) //=> false

	// declare and initialize map
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n) //=> map[foo:1, bar:2]
}

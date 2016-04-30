package main

import "fmt"

func main() {
	dayMonths := map[string]int{
		"Jan": 31,
		"Feb": 28, // simplify
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 31,
		"Dec": 31}

	fmt.Printf("Days in Feb : %d\n", dayMonths["Feb"])

	// zero type fo invalid keys
	fmt.Printf("Default value if not valid: %d\n", dayMonths["xFeb"]) //=> 0

	// check if the key is valid
	days, ok := dayMonths["January!"]
	if !ok {
		fmt.Printf("Cant' get days for January")
	} else {
		fmt.Printf("%d\n", days)
	}

	// using ranges to get the key pair (note the order is random!)
	has31 := 0
	for month, days := range dayMonths {
		fmt.Printf("%s has %d days\n", month, days)
		if days == 31 {
			has31 += 1
		}
	}

	fmt.Printf("No of month with 31 days: %d\n", has31)
}

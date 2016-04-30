package main

import "fmt"

func main() {
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28 // simplify
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["June"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 31
	dayMonths["Dec"] = 31

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

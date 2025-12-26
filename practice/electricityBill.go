package practice

import "fmt"

func Bill() {
	var units float64
	fmt.Println("Enter the number of units: ")
	fmt.Scan(&units)

	if units <= 100 {
		amount := units * 1
		fmt.Println("Your Bill is", amount, "rupees.")
		return
	} else if units <= 200 {
		amount1 := (units-100)*2 + 100
		fmt.Println("Your Bill is", amount1, "rupees.")
		return
	} else if units > 200 {
		amount2 := (units-200)*3 + (100 * 2) + (100)
		fmt.Println("Your Bill is", amount2, "rupees.")
	}
}

package practice

import "fmt"

func Temperatur() {
	for {
		fmt.Println("<Choose a conversion>")
		fmt.Println("1.Celsius-->Fahrenheit")
		fmt.Println("2.Fahrenheit-->Celsius")

		var i int
		fmt.Println("Enter conversion type:")
		fmt.Scan(&i)

		var n float64
		fmt.Println("Enter number to convert:")
		fmt.Scan(&n)

		switch i {
		case 1:
			result := (n * 9 / 5) + 32
			fmt.Println(n, "In Fahrenheit is", result)

		case 2:
			result := (n - 32) * 9 / 5
			fmt.Println(n, "In celsius is:", result)

		}
	}
}

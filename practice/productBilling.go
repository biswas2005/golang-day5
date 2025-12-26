package practice

import "fmt"

type Product struct {
	Name     string
	Quantity int
	Price    float64
}

func TotalBill() {
	var product Product
	var gstRate float64

	fmt.Print("Enter product name: ")
	fmt.Scan(&product.Name)
	fmt.Print("\nEnter quantity: ")
	fmt.Scan(&product.Quantity)
	fmt.Print("\nEnter price: ")
	fmt.Scan(&product.Price)
	fmt.Print("\nEnter GST: ")
	fmt.Scan(&gstRate)
	total := calculatingBill(product, gstRate)

	fmt.Println("<Final Bill>")
	fmt.Println("Total Bill including GST:", total)

}

func calculatingBill(n Product, gstRate float64) float64 {

	baseAmount := float64(n.Quantity) * n.Price
	finalAmount := baseAmount + (baseAmount * gstRate / 100)
	return finalAmount

}

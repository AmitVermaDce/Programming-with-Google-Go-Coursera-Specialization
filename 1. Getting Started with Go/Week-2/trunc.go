package main

import "fmt"

func main() {
	var number float64
	fmt.Print("Enter any decimal point number: ")
	fmt.Scan(&number)
	fmt.Print("Truncated form:")
	print(int64(number))

}

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var s []int = make([]int, 3)
	var inp string
	fmt.Println("Please enter an integer(X to exit):")
	for true {
		fmt.Scanln(&inp)
		if inp == "X" {
			break
		}
		ap, err := strconv.Atoi(inp)
		if err != nil {
			fmt.Println("Wrong input")
			continue
		}
		s = append(s, ap)
		sort.Ints(s[:])
		fmt.Println(s)
		fmt.Println("Please again enter an integer(X to exit):")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func swap(x, y int) (int, int) {
	temp := 0
	temp = x
	x = y
	y = temp
	return x, y
}

func bubbleSort(x []int) {
	for i := 0; i < len(x)-1; i++ {
		for j := i + 1; j < len(x); j++ {
			if x[i] > x[j] {
				x[i], x[j] = swap(x[i], x[j])
			}
		}
	}
	fmt.Println("Sorted slice using Bubble sort: ", x)

}

func main() {
	sli := make([]int, 0)
	stdin := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		fmt.Printf("Enter element %d: ", i)
		val, _, _ := stdin.ReadLine()
		number, _ := strconv.Atoi(string(val))
		sli = append(sli, number)
	}
	bubbleSort(sli)
}

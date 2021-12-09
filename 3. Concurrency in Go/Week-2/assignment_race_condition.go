package main

/*
Race condition is when multiple threads are trying to access and manipulate the same variable.
In the code below, main, add_one and sub_one are all accessing and changing the value of x.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.
*/

import "fmt"

func incr_x(x *int) {
	(*x)++
	fmt.Println("Incrementing value of x by 1")
}
func decr_x(x *int) {
	(*x)--
	fmt.Println("Incrementing value of x by 1")
}

func main() {
	i := 0
	go incr_x(&i)
	go decr_x(&i)
	i++

}

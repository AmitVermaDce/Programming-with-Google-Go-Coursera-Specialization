package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GenDisplaceFn(acceleration, init_v, init_disp float64) func(float64) float64 {
	fn := func(t float64) float64 { return 0.5*acceleration*t*t + init_v*t + init_disp }
	return fn
}
func stringToFloat(x []byte) float64 {
	value, _ := strconv.ParseFloat(string(x), 8)
	return value
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	// Input Acceleration
	fmt.Println("Enter acceleration: ")
	acceleration, _, _ := stdin.ReadLine()
	acc := stringToFloat(acceleration)

	// Input Initial Velocity
	fmt.Println("Enter Initial Velocity: ")
	initial_velocity, _, _ := stdin.ReadLine()
	ini_v := stringToFloat(initial_velocity)

	// Input Initial Displacement
	fmt.Println("Enter Initial Displacement: ")
	initial_displacement, _, _ := stdin.ReadLine()
	ini_d := stringToFloat(initial_displacement)

	// Input Time
	fmt.Println("Enter Time: ")
	time, _, _ := stdin.ReadLine()
	t := stringToFloat(time)

	// Function calling
	initial_func := GenDisplaceFn(acc, ini_v, ini_d)

	fmt.Println("Displacement travelled after t time: ", initial_func(t))

}

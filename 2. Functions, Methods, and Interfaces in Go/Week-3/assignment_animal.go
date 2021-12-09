package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (x Animal) Eat() {
	fmt.Println("Animal's food is ", x.food)

}
func (x Animal) Move() {
	fmt.Println("Animal's Locomotion is ", x.locomotion)
}
func (x Animal) Speak() {
	fmt.Println("Animal spoken sound is ", x.noise)
}
func main() {
	dict := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		inp, _, _ := stdin.ReadLine()
		inputArr := strings.Split(string(inp), " ")
		if len(inputArr) != 2 {
			fmt.Println("Incorrect input, please enter animal name and information...")
		} else {
			animalName, information := inputArr[0], inputArr[1]
			if information == "eat" {
				dict[animalName].Eat()
			} else if information == "move" {
				dict[animalName].Move()
			} else if information == "speak" {
				dict[animalName].Speak()
			}
		}

	}

}

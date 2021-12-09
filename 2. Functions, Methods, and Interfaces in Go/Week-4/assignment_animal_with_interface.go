package main

import "fmt"

type AnimalInterface interface {
	Eat()
	Move()
	Speak()
}
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

	for {
		var command, requestAni, requestType string
		var generalAni AnimalInterface
		fmt.Print(">")
		fmt.Scan(&command, &requestAni, &requestType)
		if command == "query" {
			generalAni = dict[requestAni]
			switch requestType {
			case "eat":
				generalAni.Eat()
			case "move":
				generalAni.Move()
			case "speak":
				generalAni.Speak()
			}

		}
		if command == "newanimal" {
			dict[requestAni] = dict[requestType]
			fmt.Println("Created it!")
		}

	}

}

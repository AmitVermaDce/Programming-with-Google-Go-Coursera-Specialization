package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	idMap1 := make(map[string]string)
	stdin := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter name: ")
	name, _, _ := stdin.ReadLine()

	fmt.Printf("Enter address: ")
	address, _, _ := stdin.ReadLine()

	idMap1["name"] = string(name)
	idMap1["address"] = string(address)

	json_obj, _ := json.Marshal(idMap1)
	fmt.Print("JSON object: ", string(json_obj))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter any text: ")
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	word = strings.ToLower(word)
	trimWord := strings.TrimSpace(word)
	word_len := len(trimWord) - 1
	if string(trimWord[0]) == "i" && string(trimWord[word_len]) == "n" && strings.Contains(trimWord, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}

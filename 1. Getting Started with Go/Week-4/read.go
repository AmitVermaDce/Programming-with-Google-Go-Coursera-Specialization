package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func trim_name(name string) string {
	if len(name) > 20 {
		return name[:20]
	}
	return name
}

func main() {

	path := ""
	fmt.Printf("Enter name of the file: ")
	fmt.Scan(&path)

	// Open a file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	f := bufio.NewReader(file)
	var name []Name

	for {
		data, _, err := f.ReadLine()
		if err != nil || io.EOF == err {
			break
		}
		temp_name := strings.Split(string(data), " ")
		temp_name_struct := Name{
			trim_name(temp_name[0]),
			trim_name(temp_name[1]),
		}
		name = append(name, temp_name_struct)

	}
	for i := 0; i < len(name); i++ {
		fmt.Println(i+1,".", " First Name:", name[i].fname, " and Last Name:", name[i].lname)
	}
}

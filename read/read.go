package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Name contain first and last name in a struct
type Name struct {
	firstName string
	lastName  string
}

// Set to set first name and last name to struct also trim if more than 20 characters
func (n *Name) Set(firstName, lastName string) {
	n.firstName = firstName
	n.lastName = lastName

	if len(firstName) > 20 {
		n.firstName = firstName[:20]
	}
	if len(lastName) > 20 {
		n.lastName = lastName[:20]
	}
}

func main() {
	var fileName string
	nameSlice := make([]Name, 0)

	fmt.Printf("Please input file name: ")
	fmt.Scanf("%s", &fileName)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var name Name
		line := strings.Split(scanner.Text(), " ")
		name.Set(line[0], line[1])
		nameSlice = append(nameSlice, name)
	}

	for _, val := range nameSlice {
		fmt.Println(val.firstName, val.lastName)
	}
}

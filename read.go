package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var fileName string
	fmt.Println("Enter file name: ")
	fmt.Scan(&fileName)
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	type data struct {
		firstName string
		lastName  string
	}

	s := make([]data, 0)

	for scanner.Scan() {
		result := strings.Split(scanner.Text(), " ")
		var a data
		a.firstName = result[0]
		a.lastName = result[1]
		s = append(s, a)
	}
	for _, v := range s {
		fmt.Printf("First Name: %s, Last Name: %s\n", v.firstName, v.lastName)
	}

}

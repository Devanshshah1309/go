package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}
type Bird struct {
	name string
}
type Snake struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Bird) Eat() {
	fmt.Println("worms")
}
func (c Snake) Eat() {
	fmt.Println("mice")
}

func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Bird) Move() {
	fmt.Println("fly")
}
func (c Snake) Move() {
	fmt.Println("slither")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}
func (c Bird) Speak() {
	fmt.Println("peep")
}
func (c Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		str := strings.Split(s, " ")
		if str[0] == "query" {
			name := str[1]
			animal, p := animals[name]
			if p {
				if str[2] == "eat" {
					animal.Eat()
				} else if str[2] == "move" {
					animal.Move()
				} else if str[2] == "speak" {
					animal.Speak()
				} else {
					print("Invalid! Program aborting!")
					break
				}
			}

		} else if str[0] == "newanimal" {
			name := str[1]
			t := str[2]
			var animal Animal
			if t == "cow" {
				animal = Cow{name}
			} else if t == "bird" {
				animal = Bird{name}
			} else if t == "snake" {
				animal = Snake{name}
			} else {
				print("Invalid! Program aborting!")
				break
			}
			println("Created it!")
			animals[name] = animal
		} else {
			print("Invalid! Program aborting!")
			break
		}
	}
}

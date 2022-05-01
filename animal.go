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

func (a Animal) Eat() string {
	return a.food
}
func (a Animal) Move() string {
	return a.locomotion
}
func (a Animal) Speak() string {
	return a.noise
}
func main() {
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		s := scanner.Text()
		str := strings.Split(s, " ")
		a := new(Animal)
		if str[0] == "cow" {
			a.food = "grass"
			a.locomotion = "walk"
			a.noise = "moo"

		} else if str[0] == "bird" {
			a.food = "worms"
			a.locomotion = "fly"
			a.noise = "peep"

		} else if str[0] == "snake" {
			a.food = "mice"
			a.locomotion = "slither"
			a.noise = "hiss"

		} else {
			fmt.Println("Invalid string! Program Aborting")
			break
		}
		if str[1] == "eat" {
			fmt.Println(a.Eat())
		} else if str[1] == "move" {
			fmt.Println(a.Move())
		} else if str[1] == "speak" {
			fmt.Println(a.Speak())
		} else {
			fmt.Println("Invalid string! Program Aborting")
			break
		}
	}
}

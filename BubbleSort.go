package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func BubbleSort(arr []int) {
// 	for i := len(arr) - 2; i >= 0; i-- {
// 		for j := 0; j < i; j++ {
// 			if arr[j] > arr[j+1] {
// 				Swap(arr, j)
// 			}
// 		}
// 	}
// }

// func Swap(arr []int, index int) {
// 	temp := arr[index]
// 	arr[index] = arr[index+1]
// 	arr[index+1] = temp
// }
// func StringToIntSlice(inputSequence string) []int {
// 	var numbers = []int{}
// 	var stringSequence = strings.Split(inputSequence, " ")
// 	for _, value := range stringSequence {
// 		intValue, _ := strconv.Atoi(value)
// 		numbers = append(numbers, intValue)
// 	}
// 	return numbers
// }

// func main() {
// 	// var array []int
// 	// for i := 0; i < 10; i++ {
// 	// 	fmt.Scanln(&array[i])
// 	// }
// 	var inputSequence string
// 	var numbers = []int{}
// 	scanner := bufio.NewScanner(os.Stdin)

// 	fmt.Println("Enter a Sequence of upto 10 Integers separated by spaces:")
// 	scanner.Scan()                 // moves scanner to just before the first token
// 	inputSequence = scanner.Text() // returns the most recent token generated
// 	// by a call to Scan as a newly allocated string holding its bytes.

// 	numbers = StringToIntSlice(inputSequence)
// 	BubbleSort(numbers)
// 	fmt.Print(numbers)
// }

package main

import (
	"fmt"
)

func GenDisplaceFn(acc, velocity, init float64) func(float64) float64 {
	f := func(time float64) float64 {
		return init + (velocity * time) + (0.5 * acc * time * time)
	}
	return f
}
func main() {
	print("Enter acceleration, velocity and initial displacement: ")
	var acc float64
	var velocity float64
	var init float64
	fmt.Scan(&acc, &velocity, &init)
	fn := GenDisplaceFn(acc, velocity, init)
	print("Enter time: ")
	var time float64
	fmt.Scan(&time)
	fmt.Printf("Displacement after %f time is %f", time, fn(time))
}

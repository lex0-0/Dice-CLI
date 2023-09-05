package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	return rand.Intn(sides) + 1
}
func main() {
	rand.Seed(time.Now().Unix())
	dice, sides := 2, 12
	rolls := 2
	fmt.Println("    __o  ")
	fmt.Println(" __| |_  ___ ___ ")
	fmt.Println("/ _` | |/ __/ _ |")
	fmt.Println("|(_| | | (_|  __/")
	fmt.Println("|__,_|_||___|___|")

	for r := 1; r <= rolls; r++ {
		sum := 0
		for d := 1; d <= dice; d++ {
			rolled := roll(sides)
			sum += rolled
			fmt.Println("Roll #", r, ", Die #", d, ":", rolled)
		}
		fmt.Println("total rolled:", sum)
		switch sum := sum; {
		case sum == 2 && dice == 2:
			fmt.Println("snake eyes")
		case sum == 7:
			fmt.Println("lucky number 7")
		case sum%2 == 0:
			fmt.Println("even")
		case sum%2 == 1:
			fmt.Println("odd")
		}
	}
}

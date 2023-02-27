package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in ,just press enter when ready"

func main() {

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2
	var secondNumber = rand.Intn(8) + 2

	var subtraction = rand.Intn(8) + 2

	// place to store the answer
	var answer = firstNumber*secondNumber - subtraction

	logic(firstNumber, secondNumber, subtraction, answer)

	// store the answer

	fmt.Println("the number u guess is", answer)

}

func logic(firstNumber, secondNumber, subtraction, answer int) {
	// guess the number game
	// step-1 display a welcome/instructioon
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("guess the number game ") // welcome msg
	fmt.Println("----------------------")
	fmt.Println("")

	// instruction
	fmt.Println("think of a number from 1 and 10 ", prompt)

	reader.ReadString('\n')

	// write logic for the game
	
	fmt.Println("multiply your number by", firstNumber, prompt)

	reader.ReadString('\n')

	fmt.Println("multiply the result by", secondNumber, prompt)

	reader.ReadString('\n')

	fmt.Println("divide the result by the number orginally throught of ", prompt)

	reader.ReadString('\n')

	fmt.Println("now subtract", subtraction, prompt)

	reader.ReadString('\n')
}

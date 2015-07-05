// we all know the classic "guessing game" with higher or lower prompts.
// lets do a role reversal; you create a program that will guess numbers
// between 1-100, and respond appropriately based on whether users say that
// the number is too high or too low. Try to make a program that can guess
// your number based on user input and great code!

package main

import "fmt"
import "math/rand"
import "time"

func random(max int, min int) int {
	rand.Seed(time.Now().UnixNano()) // Generate current unix time(to nanoseconds), and use it to seed time

	return rand.Intn(max-min) + min // return random number within range
}

func main() {
	var number int
	var guess int
	high := 100
	low := 1

	number = random(high, low)

	for guess != number {
		fmt.Print("Your Guess: ")
		fmt.Scanf("%d\n", &guess)
	}

	fmt.Printf("Your guess of %d was correct!", guess)
	fmt.Scanln()
}

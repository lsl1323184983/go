package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().Unix())
	secrerNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is ", secrerNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("An error occured with reading input. please try again", err)
			// return
			continue
		}
		input = strings.TrimSuffix(input, "\r\n")

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Print("Invalid input. Please enter an integer value", err)
			// return
			continue
		}
		fmt.Println("Your guess is", guess)
		if guess > secrerNumber {
			fmt.Println("Your guess is bigger than the secret number.Please try again")
		} else if guess < secrerNumber {
			fmt.Println("Your guess is samller than the secret number.Please try again")
		} else {
			fmt.Println("Correct, you legend!")
			break
		}
	}

}

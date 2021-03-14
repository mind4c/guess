package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	"strconv"
	"bufio"
	"os"
	"log"
)

func main(){

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	min_num, max_num := 0, 100
	target := rand.Intn(max_num) + 1
	fmt.Printf("Please guess between %d - %d \n", min_num, max_num)

	reader := bufio.NewReader(os.Stdin)

	right_guess := false
	max_try := 6
	for i := 0; i < max_try; i++ {

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("Error while read stdin", err)
		}
		
		candidate, err := strconv.Atoi( strings.TrimSpace(input) )

		if err != nil {
			log.Printf("Input is not a number: %v \n", input)
			continue
		}

		if candidate == target {
			right_guess = true
			fmt.Println("Guess right!")
			break
		}

		if i == max_try -1 {
			break
		}

		if candidate < target {
			fmt.Println("Try something larger.")
			continue
		}
		fmt.Println("Try something smaller.")
	}

	if !right_guess {
		fmt.Printf("Game over! Target is: %d, try next time. \n", target)
	}


}

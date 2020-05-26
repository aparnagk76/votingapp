package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	voteMap := make(map[string]int)
	var winner string
	maxNumber := 0

	fmt.Println("voting app")
	for {
		fmt.Println("Please cast your vote with the following options")
		fmt.Println("1 Donald Trump")
		fmt.Println("2 Hilory Clinton")
		fmt.Println("Enter a number")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = strings.Replace(input, "\n", "", -1)
		if input == "q" {
			for key, value := range voteMap {
				if value > maxNumber {
					maxNumber = value
					winner = key
				}
			}

			fmt.Println("the summary of votin app is ", winner, "with total votes ", maxNumber)
			fmt.Println("we are done. Thank you")
			fmt.Println("bye")
			break
		}

		voteOption, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}

		if voteOption == 1 {
			voteMap["DT"] = voteMap["DT"] + 1
			fmt.Println("you have casted vote for Donald Trump")
			fmt.Println("Thank you for casting vote")
		} else if voteOption == 2 {
			voteMap["HC"] = voteMap["HC"] + 1
			fmt.Println("you have casted vote for Hilory Clinton")
			fmt.Println("Thank you for casting vote")
		} else {
			fmt.Println("Invalid option")
		}

	}
}

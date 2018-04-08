//Colm Woodlock
//G00341460
//Graph Theory Project

package main

import (
	"bufio"
	"fmt" //https://golang.org/pkg/bufio/

	"os"
)

func main() {

	exit := true
	option := 0

	fmt.Println(" Please enter 1 to parse the regular expression from infix to postfix notation \n Please enter 2 to parse regular expression to nfa \n Please enter 3 to exit")
	fmt.Scanln(&option)

	for exit {
		if option == 1 {
			fmt.Println("Please enter the regular expression you want to convert into postfix notation: ")
			reader := bufio.NewReader(os.Stdin)
			regex, _ := reader.ReadString('\n')

			fmt.Println("Infix:  ", regex)
			fmt.Println("Postfix: ", intopost(regex))

		} else if option == 2 {
			fmt.Println("Please enter the regular expression you want to convert to nfa: ")

		} else if option == 3 {
			fmt.Println("Exiting...")
			exit = false
		} else {
			fmt.Println("Please enter a valid option (e.g 3 to exit)")
		}
	}
}

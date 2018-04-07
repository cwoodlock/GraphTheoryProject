//Colm Woodlock
//G00341460
//Graph Theory Project

package main

import (
	"fmt"
)

func main() {

	exit := true
	option := 0

	fmt.Println("Please enter 1 to parse the regular expression from infix to postfix notation \n Please enter 2 to parse regular expression to nfa \n Please enter 3 to exit")
	fmt.Scanln(&option)

	for exit {
		if option == 1 {

		} else if option == 2 {

		} else if option == 3 {
			exit = false
		} else {
			fmt.Println("Please enter a valid option (e.g 3 to exit)")
		}
	}
}

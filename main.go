//Colm Woodlock
//G00341460
//Graph Theory Project
//This code is adapted from the video lectures given to us

package main

import (
	"bufio" //https://golang.org/pkg/bufio/
	"fmt"   //https://golang.org/pkg/bufio/
	"os"    //https://golang.org/pkg/os/
)

type state struct {
	symbol rune
	edge1  *state
	edge2  *state
}

type nfa struct {
	initial *state
	accept  *state
}

func main() {

	exit := true
	option := 0

	//prompt user to enter a number which corresponds to a function of the program
	fmt.Println(" Please enter 1 to parse the regular expression from infix to postfix notation \n Please enter 2 to parse regular expression to nfa \n Please enter 3 to exit")
	fmt.Scanln(&option) //read in the users selection

	for exit {
		if option == 1 {
			fmt.Println("Please enter the regular expression you want to convert into postfix notation: ")
			reader := bufio.NewReader(os.Stdin) //read in the regular expression the user entered
			regex, _ := reader.ReadString('\n')

			fmt.Println("Infix:  ", regex)            //display the infix regular expression
			fmt.Println("Postfix: ", Intopost(regex)) //display the postfix regular expression

		} else if option == 2 {
			fmt.Println("Please enter the regular expression you want to convert to nfa: ")
			reader := bufio.NewReader(os.Stdin) //read in user regular expression
			regex, _ := reader.ReadString('\n')

			fmt.Println("Postfix:  ", regex)        //display the postfix regular expression
			fmt.Println("NFA: ", Poregtonfa(regex)) //display tthe converted regular expression as an nfa

			fmt.Println("Please enter the string to see if it matches the nfa: ")
			userString, _ := reader.ReadString('\n') //read in string to compare to the regular expression
			userString = Intopost(userString)

			if pomatch(regex, userString) == false { //pomatch compares the regular expression and the user entered string

				fmt.Println("The string does not match")

			} else if pomatch(regex, userString) == true {

				fmt.Println("String: ", userString, " matches the regulare expression: ", regex)

			} else {
				fmt.Println("Error..")
			}

		} else if option == 3 {
			fmt.Println("Exiting...") //option 3 exits the program
			exit = false
		} else {
			fmt.Println("Please enter a valid option (e.g 3 to exit)") //catch invalid inputs
		}
	}
} // end main

func Intopost(infix string) string {
	specials := map[rune]int{'*': 10, '.': 9, '|': 8} //Map to keep track of special chars, and order of precedence

	pofix := []rune{}
	s := []rune{} //s is stack

	for _, r := range infix { //Range on string will convert it to array of runes
		switch {
		case r == '(':
			s = append(s, r) //temp put open bracket on the end of the stack

		case r == ')':
			for s[len(s)-1] != '(' { //If we see a closin bracket we are going to pop things off the stack until we see a close bracket
				pofix = append(pofix, s[len(s)-1]) //last element of s still need to pop it off
				s = s[:len(s)-1]                   //everything up to last element and set it to s
			}
			s = s[:len(s)-1]

		case specials[r] > 0:
			for len(s) > 0 && specials[r] <= specials[s[len(s)-1]] {
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
			}
			s = append(s, r)

		default:
			pofix = append(pofix, r)
		}
	}

	for len(s) > 0 {
		pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1]
	}

	return string(pofix)
} //end intopost

func Poregtonfa(pofix string) *nfa {
	nfastack := []*nfa{}

	for _, r := range pofix {
		switch r {
		case '.': //concatonate
			frag2 := nfastack[len(nfastack)-1] //create two fragments
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|': //or
			frag2 := nfastack[len(nfastack)-1] //create two fragments
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*': //Kleene star (any number of)
			frag := nfastack[len(nfastack)-1] //create one fragment
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '+': //at least one of
			frag := nfastack[len(nfastack)-1] //pop one frag off the nfa nfastack
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}                                     //create a new accept states
			initial := state{edge1: frag.initial, edge2: &accept} //create new state initial where edge1 pounts to frag.initial and edge2 points at the new accept state

			frag.accept.edge1 = &initial                                              //frag edge1 points to the initial states
			nfastack = append(nfastack, &nfa{initial: frag.initial, accept: &accept}) //push the new frag onto the stack

		case '?': //zero or one of
			frag := nfastack[len(nfastack)-1] //pop one frag off the nfa nfastack

			initial := state{edge1: frag.initial, edge2: frag.accept} //create new state initial

			nfastack = append(nfastack, &nfa{initial: &initial, accept: frag.accept}) //push new frag onto the stack

		default:
			accept := state{}
			initial := state{symbol: r, edge1: &accept}

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})
		}
	}
	if len(nfastack) != 1 { //if there is an error catch it and alert the user
		fmt.Println("uh oh: ", len(nfastack), nfastack)
	}

	return nfastack[0]
} //end Poregtonfa

func addState(l []*state, s *state, a *state) []*state { //helper function
	l = append(l, s)

	if s != a && s.symbol == 0 { //if it is a state that has an E arrow coming from it, and check we are not in accept state
		l = addState(l, s.edge1, a)
		if s.edge2 != nil {
			l = addState(l, s.edge2, a)
		}
	}
	return l
} //end addState

func pomatch(po string, s string) bool { //does po match s
	ismatch := false
	ponfa := Poregtonfa(po)

	current := []*state{ponfa.initial} //current set of states im in
	next := []*state{}                 //next set of states, generate next from current

	current = addState(current[:], ponfa.initial, ponfa.accept) //[:] turns current into a slice from an array

	for _, r := range s {
		for _, c := range current { //take all current states that you're in
			if c.symbol == r { //check if they are labeled by rune i just read
				next = addState(next[:], c.edge1, ponfa.accept)
			}

		}
		current, next = next, []*state{} //swap current for next and replace next with empty array of states
	}

	for _, c := range current { //loop through and check if they are the accpet state
		if c == ponfa.accept {
			ismatch = true
			break
		}
	}

	return ismatch
} //end pomatch

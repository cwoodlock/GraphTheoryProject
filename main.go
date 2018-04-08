//Colm Woodlock
//G00341460
//Graph Theory Project

package main

import (
	"bufio"
	"fmt" //https://golang.org/pkg/bufio/
	"os"
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

	fmt.Println(" Please enter 1 to parse the regular expression from infix to postfix notation \n Please enter 2 to parse regular expression to nfa \n Please enter 3 to exit")
	fmt.Scanln(&option)

	for exit {
		if option == 1 {
			fmt.Println("Please enter the regular expression you want to convert into postfix notation: ")
			reader := bufio.NewReader(os.Stdin)
			regex, _ := reader.ReadString('\n')

			fmt.Println("Infix:  ", regex)
			fmt.Println("Postfix: ", Intopost(regex))

		} else if option == 2 {
			fmt.Println("Please enter the regular expression you want to convert to nfa: ")

		} else if option == 3 {
			fmt.Println("Exiting...")
			exit = false
		} else {
			fmt.Println("Please enter a valid option (e.g 3 to exit)")
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
				pofix, s = append(pofix, s[len(s)-1]), s[:len(s)-1] //Condensed version of lines 20 + 21
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
		case '.':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			frag1.accept.edge1 = frag2.initial

			nfastack = append(nfastack, &nfa{initial: frag1.initial, accept: frag2.accept})

		case '|':
			frag2 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]
			frag1 := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			initial := state{edge1: frag1.initial, edge2: frag2.initial}
			accept := state{}
			frag1.accept.edge1 = &accept
			frag2.accept.edge1 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

		case '*':
			frag := nfastack[len(nfastack)-1]
			nfastack = nfastack[:len(nfastack)-1]

			accept := state{}
			initial := state{edge1: frag.initial, edge2: &accept}
			frag.accept.edge1 = frag.initial
			frag.accept.edge2 = &accept

			nfastack = append(nfastack, &nfa{initial: &initial, accept: &accept})

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

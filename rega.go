//Colm Woodlock
//G00341460
//Graph Theory Project
//Code adapted from video lectures : https://web.microsoftstream.com/video/bad665ee-3417-4350-9d31-6db35cf5f80d

package methods

import (
	"fmt"
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

/*
func main() {
	fmt.Println(pomatch("ab.c*|", "cccc"))
} //end main
*/

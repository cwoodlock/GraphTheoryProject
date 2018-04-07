//G00341460
//Graph Theory Project
//Code adapted from video lectures : https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e

package methods

func intopost(infix string) string {
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

/*
func main() {

	//Answer ab.c*.
	fmt.Println("Infix: ", "a.b.c*")
	fmt.Println("Postfix: ", intopost("a.b.c*"))

	//Answer abd|.*
	fmt.Println("Infix: ", "(a.(b|d))*")
	fmt.Println("Postfix: ", intopost("(a.(b|d))*"))

	//Answer abd|.c*
	fmt.Println("Infix: ", "a.(b|d).c*")
	fmt.Println("Postfix: ", intopost("a.(b|d).c*"))

	//Answer abb.+.c.
	fmt.Println("Infix: ", "a.(b.b)+.c")
	fmt.Println("Postfix: ", intopost("a.(b.b)+.c"))

}
*/

// modified https://gist.github.com/uchan-nos/cd594b3a4c88af136bd4
package main

import (
	"strconv"
	"unicode"
	"unicode/utf8"
)

func skipSpaces(s []byte) []byte {
	c, w := utf8.DecodeRune(s)
	for w > 0 && unicode.IsSpace(c) {
		s = s[w:]
		c, w = utf8.DecodeRune(s)
	}
	return s
}

func readDigits(s []byte) (numStr, remain []byte) {
	numStr = s
	totalW := 0
	c, w := utf8.DecodeRune(s)
	for w > 0 && unicode.IsDigit(c) {
		s = s[w:]
		totalW += w
		c, w = utf8.DecodeRune(s)
	}
	return numStr[:totalW], s
}

func pop(stack []int) (int, []int) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func eval(s []byte) []int {
	stack := make([]int, 0)
	var a, b int
	var token []byte

	s = skipSpaces(s)
	for len(s) > 0 {
		c, w := utf8.DecodeRune(s)
		switch {
		case unicode.IsDigit(c):
			token, s = readDigits(s)
			num, _ := strconv.Atoi(string(token))
			stack = append(stack, num)
		case c == '+':
			b, stack = pop(stack)
			a, stack = pop(stack)
			stack = append(stack, a+b)
			s = s[w:]
		case c == '-':
			b, stack = pop(stack)
			a, stack = pop(stack)
			stack = append(stack, a-b)
			s = s[w:]
		case c == '*':
			b, stack = pop(stack)
			a, stack = pop(stack)
			stack = append(stack, a*b)
			s = s[w:]
		case c == '/':
			b, stack = pop(stack)
			a, stack = pop(stack)
			stack = append(stack, a/b)
			s = s[w:]
		case c == '%':
			b, stack = pop(stack)
			a, stack = pop(stack)
			stack = append(stack, a%b)
			s = s[w:]
		default:
			s = s[w:]
		}
		s = skipSpaces(s)
	}

	return stack
}

// func main() {
// 	// 2 * 21 - 30 = 12
// 	eval([]byte("2 21 * 30-"))

// 	// 1 + ... + 10 = 55
// 	eval([]byte("1 2 3 4 5 6 7 8 9 10+++++++++"))
// }

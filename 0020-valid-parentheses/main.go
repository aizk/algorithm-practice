package main

import "fmt"

func main() {
	fmt.Println(isValid("[]"))
}

type Stack []rune

func (s Stack) Push(c rune) Stack {
	return append(s, c)
}

func (s Stack) Pop() (Stack, rune) {
	l := len(s)
	if l == 0 {
		return s, rune(0)
	}
	return s[:l - 1], s[l-1]
}

func isValid(s string) bool {
	var st Stack
	for _, c := range s {
		if c == '[' || c == '(' || c == '{' {
			st = st.Push(c)
		} else {
			var w rune
			st, w = st.Pop()
			if canClose(string(w) + string(c)) {
				continue
			} else {
				return false
			}
		}
	}
	if len(st) > 0 {
		return false
	}

	return true
}

func canClose(s string) bool {
	switch {
	case s == "()":
		return true
	case s == "{}":
		return true
	case s == "[]":
		return true
	default:
		return false
	}
	return false
}
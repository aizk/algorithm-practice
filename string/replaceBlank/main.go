package main

import (
	"errors"
	"fmt"
)

func main() {
	TestReplaceBlank("", "", ErrEmpty)
	TestReplaceBlank("123", "123", nil)
	TestReplaceBlank("We are happy.", "We%20are%20happy.", nil)
	TestReplaceBlank("     ", "%20%20%20%20%20", nil)
}

func TestReplaceBlank(s string, expected string, expectedErr error) {
	rsp, err := ReplaceBlank(s)
	if rsp != expected || err != expectedErr {
		fmt.Printf("input: %s | expected: %s | output: %s | expectedErr: %v | outputErr: %v \n", s, expected, rsp, expectedErr, err)
	}
}

var ErrEmpty = errors.New("invalid empty param")

func ReplaceBlank(s string) (string, error) {
	var r []byte
	if len(s) == 0 {
		return "", ErrEmpty
	}
	for _, v := range s {
		if v == ' ' {
			r = append(r, '%', '2', '0')
		} else {
			r = append(r, byte(v))
		}
	}
	return string(r), nil
}

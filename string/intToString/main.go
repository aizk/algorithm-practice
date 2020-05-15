package main

import (
	"errors"
	"fmt"
)

func main() {
	TestStrToInt("", 0, 0, ErrSyntax)
}

func TestStrToInt(s string, base int, expected int, expectedErr error) {
	n, err := StrToInt(s, base)
	if n != expected || err != expectedErr {
		fmt.Printf("input: %s | expected: %s | output: %s | expectedErr: %v | outputErr: %v \n", s, expected, n, expectedErr, err)
	}
}

var ErrSyntax = errors.New("invalid syntax")
var ErrRange = errors.New("value out of range")

func StrToInt(s string, base int) (int, error) {
	var n int
	var cutoff, maxVal int

	maxVal = 1 << 63 - 1

	if len(s) < 1 {
		return n, ErrSyntax
	}

	// neg
	neg := false
	if s[0] == '+' {
		s = s[1:]
	} else if s[0] == '-' {
		s = s[1:]
		neg = true
	}
	if len(s) < 1 {
		return n, ErrSyntax
	}

	i := 0
	switch {
	case len(s) < 1:
		return n, ErrSyntax
	case 2 <= base && base <= 36:
		// do nothing
	case base == 0:
		switch {
		case s[0] == '0' && (s[1] == 'x' || s[1] == 'X'):
			// 16 进制
			base = 16
			i = 2
		case s[0] == '0':
			// 8 进制
			base = 8
			i = 1
		}
	default:
		base = 10
	}

	// cutoff
	switch base {
	case 10:
		cutoff = maxVal/10 + 1
	case 16:
		cutoff = maxVal/16 + 1
	default:
		cutoff = maxVal/base + 1
	}

	for ; i < len(s); i++ {
		var v byte
		d := s[i]
		// 读取数字
		switch {
		case '0' <= d && d <= '9':
			v = d - '0'
		case 'a' <= d && d <= 'z':
			v = d - 'a' + 10
		case 'A' <= d && d <= 'Z':
			v = d - 'A' + 10
		case d == ' ':
			continue
		default:
			return n, ErrSyntax
		}
		if int(v) >= base {
			return n, ErrSyntax
		}
		if n >= cutoff {
			n = maxVal
			return n, ErrRange
		}
		n *= base
		n += int(v)
	}

	if neg {
		n = n * -1
	}

	return n, nil
}

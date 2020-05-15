package main

import (
	"strings"
	"errors"
	"fmt"
)

func main() {
	TestVersionCompare("1.2.3a", "1.2.3b", -1, nil)
	TestVersionCompare("1.2.3b", "1.2.3b", 0, nil)
	TestVersionCompare("11.2.3b", "1.2.3b", 1, nil)
	TestVersionCompare("1.2.3b", "11.2.3b", -1, nil)
	TestVersionCompare("", "", 0, ErrSyntax)
	TestVersionCompare("!.y.z", "z.$.x", 0, ErrSyntax)
}

func TestVersionCompare(v1, v2 string, expected int, expectedErr error) {
	n, err := VersionCompare(v1, v2)
	if n != expected || err != expectedErr {
		fmt.Printf("input: %s and %s | expected: %s | output: %s | expectedErr: %v | outputErr: %v \n", v1, v2, expected, n, expectedErr, err)
	}

}

var ErrSyntax = errors.New("invalid syntax")

// if v1 > v2 return 1
// else v1 < v2 return -1
// equal return 0
func VersionCompare(v1, v2 string) (int, error) {
	if len(v1) < 1 || len(v2) < 1 {
		return 0, ErrSyntax
	}
	v1Str := strings.Split(v1, ".")
	v2Str := strings.Split(v2, ".")
	if len(v1Str) <= len(v2Str) {
		for key := range v1Str {
			r, err := compare(v1Str[key], v2Str[key])
			if err != nil {
				return 0, err
			}
			if r != 0 {
				return r, nil
			}
		}
		return 0, nil
	} else {
		for key := range v2Str {
			r, err := compare(v1Str[key], v2Str[key])
			if err != nil {
				return 0, err
			}
			if r != 0 {
				return r, nil
			}
		}
		return 0, nil
	}
}

func compare(x, y string) (int, error) {
	if len(x) == len(y) {
		for k := range x {
			xVal := x[k]
			yVal := y[k]
			if err := checkSymbol(xVal); err != nil {
				return 0, err
			}
			if err := checkSymbol(yVal); err != nil {
				return 0, err
			}
			if xVal > yVal {
				return 1, nil
			} else if xVal < yVal {
				return -1, nil
			}
		}
	} else if len(x) > len(y) {
		return 1, nil
	} else {
		return -1, nil
	}
	return 0, nil
}

func checkSymbol(b byte) error {
	switch {
	case b >= '0' && b <= '9' :
		return nil
	case b >= 'a' && b <= 'z':
		return nil
	}
	return ErrSyntax
}
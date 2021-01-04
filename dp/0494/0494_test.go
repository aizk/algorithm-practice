package _494

import (
	"fmt"
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	input := []int{46,49,5,7,5,21,27,4,4,27,45,24,7,22,25,5,38,14,50,28}
	S := 45

	res := FindTargetSumWays(input, S)
	fmt.Println(res)

	Recursive(input, S)
	fmt.Println("Count: ", Count)
}

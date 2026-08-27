package main

import (
	"cmp"
	"fmt"

	"golang.org/x/exp/constraints"
)

type (
	_ constraints.Float
	_ constraints.Integer
	_ constraints.Signed
	_ constraints.Unsigned
	_ constraints.Complex
	_ constraints.Ordered

	_ cmp.Ordered
)

func Min[T constraints.Ordered](a, b T) T {
	var min T

	if a < b {
		return a
	} else if a > b {
		return b
	}
	return min
}

func main() {
	fmt.Println(Min(10, 100))
}

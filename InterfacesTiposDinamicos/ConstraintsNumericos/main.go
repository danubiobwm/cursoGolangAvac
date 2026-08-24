package main

import "fmt"

type MtInt int

type Number interface {
	~int | ~float64
}

func Sum[T Number](a, b T) T {
	return a + b
}

func main() {

	var (
		x1 = int(10)
		x2 = int(20)
		y1 = float64(10.5)
		y2 = float64(20.5)

		z1 = MtInt(30)
		z2 = MtInt(40)
	)

	fmt.Println("Sum of integers:", Sum(x1, x2))
	fmt.Println("Sum of floats:", Sum(y1, y2))
	fmt.Println("Sum of mixed types:", Sum(y1, y1))
	fmt.Println("Sum of custom type:", Sum(z1, z2))

}

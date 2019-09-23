package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	fmt.Println(s)

	n := 500000000.0

	d := 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}

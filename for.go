package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for {
		fmt.Println("Loop")
		break
	}
	for n := 1; n <= 10; n++ {
		if n%2 != 0 {
			continue
		}
		fmt.Println(n)
	}
}

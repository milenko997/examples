package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("Prazan niz:", a)

	a[4] = 1
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println("Duzina niza je: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Niz: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2 D niz: ", twoD)
}

package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("Paran")
	} else {
		fmt.Println("Neparan")
	}

	if 8%4 == 0 {
		fmt.Println("Brojevi su deljivi")
	} else {
		fmt.Println("Brojevi nisu deljivi")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "je negativan broj")
	} else if num < 10 {
		fmt.Println(num, "je jednocifren broj")
	} else {
		fmt.Println(num, "ima vise od jedne cifre")
	}

}

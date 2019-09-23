package main

import (
	"fmt"
	"time"
)

func main() {
	i := 1
	fmt.Print(i, " ")
	switch i {
	case 1:
		fmt.Println("kec")
	case 2:
		fmt.Println("dvojka")
	case 3:
		fmt.Println("trojka")
	default:
		fmt.Println("Preveliki broj")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Vikend")

	default:
		fmt.Println("Radni dan")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Prepodne")
	default:
		fmt.Println("Posle podne")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Boolean vrednost")
		case int:
			fmt.Println("Integer vrednost")
		default:
			fmt.Println("Nepoznata vrednost", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

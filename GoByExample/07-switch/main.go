package main

import (
	"fmt"
	"time"
)

func main() {
	n := 3
	switch n {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// use comma "," to cover multiple options
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a week day")
	}

	// switch without expression is a shortcut for if-then-elseif
	switch {
	case time.Now().Hour() < 12:
		fmt.Println("It's before noon")
	case time.Now().Hour() > 12:
		fmt.Println("It's afer noon")
	default:
		fmt.Println("It's the noon!")
	}

	// type switch compares type not value
	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am bool")
		case int:
			fmt.Println("I am int")
		default:
			fmt.Printf("I dont know my type - %T\n", i)
		}
	}
	whatAmI(12)
	whatAmI(true)
	whatAmI("no name given")
}

package main

import "fmt"
import "time"

func main() {

	i := 2
	fmt.Println("write", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")

	}

	//comma separated values for multiple expressions, optional default
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend (bout to have me some fun)")
	default:
		fmt.Println("It's a weekday")
	}
}

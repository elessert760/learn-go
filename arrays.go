package main

import "fmt"

func main() {

	var favNums2 [5]float64

	favNums2[0] = 3.14
	favNums2[1] = 2.71828
	favNums2[2] = 1.618
	favNums2[3] = 304.0
	favNums2[4] = 99

	fmt.Println(favNums2[1])

	favNums3 := [5]float64{1, 2, 3.14, 4, 5}

	fmt.Println(favNums3[2])

	for i, value := range favNums3 {
		fmt.Println(value, i+1)
	}

	elem := "The 4th element of the array is "
	i := favNums3[3]
	fmt.Println(elem, i)

	switch i {
	case 1:
		fmt.Println(elem, "one")
	case 2:
		fmt.Println(elem, "two")
	case 3:
		fmt.Println(elem, "three")
	case 4:
		fmt.Println(elem, "four")
	}

}

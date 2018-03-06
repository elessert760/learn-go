package main

import "fmt"

func plus(a int, b int) int{

  return a + b
}

func plusPlus(a, b, c int) int{
  return a + b + c
}



func main()  {

  res := plus(99, 1)
  fmt.Println(res)

  res = plusPlus(101, -11, -22)
  fmt.Println(res)

 }

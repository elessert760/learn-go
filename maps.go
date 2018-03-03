package main

import "fmt"

func main(){

  Grades := make(map[string]float32)

  Grades["Eric"] = 99
  Grades["Timmy"] = 42
  Grades["Jesse"] = 100
  Grades["Finn"] = 99

  fmt.Println(Grades)

  for k, v := range Grades{
    fmt.Println(k, ":", v)
  }

  delete(Grades, "Timmy")
  fmt.Println(Grades)


  n := map[string]int{"foo": 1, "bar": 2, "baz":304}
   fmt.Println("map:", n)

}

package main

import "fmt"
import "strings"

func main() {

  nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
      // use the _ to skip the index, because we only want to sum the values
        sum += num
    }
    fmt.Println("sum:", sum)

    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    for k := range kvs {
        fmt.Println("key:", strings.ToLower(strings.ToUpper(k)))
    }

    for i, c := range "go" {
       fmt.Println(i, c)
   }

}

package main

import (
  "fmt"
  "os"
  "strconv"
)

func main() {
  //get values
  x, _ := strconv.Atoi(os.Args[1])
  y, _ := strconv.Atoi(os.Args[2])
  count, _ := strconv.Atoi(os.Args[3])
  // fmt.Println(x)
  // fmt.Println(y)

  sum := 0
  m := make(map[int]int)

  for index := x; index < count; index += x {
    if index % x == 0 {
      m[index] = index
      // fmt.Println(index)
    }
  }

  for index := y; index < count; index += y {
    if index % y == 0 {
      m[index] = 0
      // fmt.Println(index)
    }
  }

  for key := range m {
    sum += key
  }

  //fmt.Println(m)
  fmt.Println(sum)
}

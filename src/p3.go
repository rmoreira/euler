// Problem 3: Largest Prime factor
//
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import(
  "fmt"
  "os"
  "strconv"
)

func main() {
  max, _ := strconv.Atoi(os.Args[1])
  last := 1
  for index := 2; index <= max; index++ {
    if max%index == 0 {
      max/=index
      //fmt.Println(index)
      last = index
    }
  }
  fmt.Println(last)
}

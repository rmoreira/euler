// Problem 4: Largest palindrome product
// A palindromic number reads the same both ways.
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import(
  "fmt"
  "os"
  "strconv"
  "math"
)

func palindronme(x, y int) bool {
  result := strconv.Itoa(x*y)
  count := len(result)
  answer := false
  // evenORodd := "even"
  // m := make(map[int]int)
  // fmt.Println(math.Pow10(3))
  // fmt.Println("x,y", x,y)
  // fmt.Println("result", result)
  // fmt.Println("count", count)
  // fmt.Println(x,y)
  if count % 2 == 0 { //even
    // evenORodd = "even"
    for index := 0; index < count - 1 - index; index++ {
      // fmt.Println(result[index],result[count - 1 - index])
      if !(result[index] == result[count - 1 - index]) {
        answer = false
        break
      } else {
        answer = true
      }
      // fmt.Println("even", x*y)
    }
  } else { //odd
    // evenORodd = "odd"
    for index := 0; index < count - 2 - index; index++ {
      // fmt.Println(result[index],result[count - 2 - index])
      if !(result[index] == result[count - 1 - index]) {
        answer = false
        break
      } else {
        answer = true
      }
      // fmt.Println("odd", x*y)
    }
  }

  // if answer {
  //   fmt.Println(x,y,x*y,evenORodd)
  // }
  return answer

}
func main() {
  numSize, _ := strconv.Atoi(os.Args[1])
  x := (int(math.Pow10(numSize))) - 1
  y := (int(math.Pow10(numSize))) - 1
  limit := (int(math.Pow10(numSize - 1)))
  largestPal := -1
  num1 := -1
  num2 := -1

  for i,counter := x,0; i > limit && num1 >= num2 ; i-- {
    for j := y; j > limit; j-- {
      if palindronme(i, j){
        multiplication := i*j
        // fmt.Println(multiplication,i,j)
        if multiplication >= largestPal {
          largestPal,num1,num2 = multiplication,i,j
        }
        break
      }
    }
    counter++
  }
  fmt.Println(largestPal,num1,num2)
}

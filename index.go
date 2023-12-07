package main

import (
  "fmt"
)

func add(x, y int) int {
 return x + y
}

func mul(x, y int) int {
  return x * y
}

func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  return arithmetic(arithmetic(a, b), c);
}

func main() {
  fmt.Println(aggregate(2, 3, 4, add)) 
  squareFunc := selfMath(mul)
  fmt.Println(squareFunc(5))
  
  resultFc1 := test(fc2);
  fmt.Println(resultFc1(3, 5, 1))
}

func selfMath(mathFunc func(int, int) int) func (int) int {
    return func(x int) int {
      return mathFunc(x, x)
  }
}

func test(fc1 func (int, int, int) int) func (int, int, int) int {
    return func (x int, y int, z int) int {
        return fc1(x, y, z) - z
  }
}

func fc2 (a int, b int, z int) int {
  return a + b
}

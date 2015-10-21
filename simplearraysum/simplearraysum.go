package main

import "fmt"

func simpleArraySum(size int, array []int) int{
  result := 0
  for _, value := range array {
    result += value
  }
  return result
}

func main() {
  var size, result int
  fmt.Scanf("%d", &size)
  array := make([]int, size) // importante
  for i := 0; i < size; i++{
    fmt.Scanf("%d", &result)
    array[i] = result
  }
  res := simpleArraySum(size, array)
  fmt.Println(res)
}

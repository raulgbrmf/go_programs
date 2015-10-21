package main

import "math"
import "fmt"


func extractDiagonals(size int, matrix [][]int)float64 {
  var diagonal1, diagonal2 int
  diagonal1 = 0
  diagonal2 = 0
  var result float64

  j := 0
  k := size-1
  for i := 0; i < size; i++ {
     // to retrieve the reverse diagonal
    diagonal1 += matrix[i][j]
    diagonal2 += matrix[i][k]
    j++
    k--
    }
  result = float64(diagonal1) - float64(diagonal2)
  result = math.Abs(result)
  return result
}

func main(){
  var size int
    fmt.Scanf("%d", &size)
    matrix := make([][]int, size)
  	for i := 0; i < size; i++ {
  		matrix[i] = make([]int, size)
  		for j := 0; j < size; j++ {
  		 fmt.Scanf("%d", &matrix[i][j])
  		}
    }
    result := extractDiagonals(size, matrix)
    fmt.Println (result)
  }

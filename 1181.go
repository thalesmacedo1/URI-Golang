package main

import "fmt"

func URI1181() {
  var targetLine int
  var operationType string
  var matrix[12][12] float32
  var sum, displayResult float32
  
  fmt.Scanf("%d", &targetLine)
  fmt.Scanf("%s", &operationType)
  
  for i := 0; i < 12; i++ {
      for j := 0; j < 12; j++ {
        fmt.Scanf("%f", &matrix[i][j])
      }
  }

  for i := 0; i < 12; i++ {
    sum += matrix[targetLine][i]
  }
  
  mean := sum / 12.0
  
  if operationType == "S" {
    displayResult = sum
  } else if operationType == "M" {
    displayResult = mean
  }
  
  fmt.Printf("%.1f\n", displayResult)
}

package main

import "fmt"

func URI1180() {
  var times, currentNumber, lower, position int
  fmt.Scanf("%d", &times)
  
  for i := 0; i < times; i++ {
    fmt.Scanf("%d", &currentNumber)
    
    if i == 0 || currentNumber < lower {
      lower = currentNumber
      position = i
    }
  }
  
  fmt.Printf("Menor valor: %d\n", lower)
  fmt.Printf("Posicao: %d\n", position)
}

package main

import "fmt"

func URI1046() {
  var start, end int
  
  fmt.Scanf("%d", &start, &end)
  fmt.Scanf("%d", &end)

  time := end - start
  
  if start == end {
    fmt.Printf("O JOGO DUROU 24 HORA(S)\n")
  } else if start > end{
    time = 24 - (start - end)
    fmt.Printf("O JOGO DUROU %d HORA(S)\n", time)
  } else if start < end{
    fmt.Printf("O JOGO DUROU %d HORA(S)\n", time)
  }
}

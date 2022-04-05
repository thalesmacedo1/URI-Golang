package main

import (
    "fmt"
    "sort"  
)

func URI2428() {
  var dimensions [4]int

  fmt.Scanf("%d", &dimensions[0])
  fmt.Scanf("%d", &dimensions[1])
  fmt.Scanf("%d", &dimensions[2])
  fmt.Scanf("%d", &dimensions[3])

  sort.Ints(dimensions[:])
  
  if dimensions[0] * dimensions[3] == dimensions[1] * dimensions[2] {
    fmt.Printf("S\n")
  } else {
    fmt.Printf("N\n")
  }
}
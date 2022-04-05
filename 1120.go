package main

import (
    "fmt"
)

func URI1120() {
  var failedDigit byte
  var contractValue string
  
  for {
    fmt.Scanf("%b", &failedDigit)
    fmt.Scanf("%s", &contractValue)

    if failedDigit == '0' && contractValue == "0" {
      break
    }

    for i := 0; i<len(contractValue); i++ {
      if failedDigit == contractValue[i] {
        continue
      } else {
        
      }
    }

    
    fmt.Printf("%s\n", contractValue)
  }
}

package main

import "fmt"

func URI1168() {
  var times int
  var numbers string

  leds := 0
  
  fmt.Scanf("%d", &times)
  
  for i := 0; i < times; i++ {
    fmt.Scanf("%s",&numbers)
    for j := 0; j < len(numbers); j++ {
      if numbers[j] == '1' {
        leds += 2;
      } else if numbers[j] == '2' ||
                numbers[j] == '3' ||
                numbers[j] == '5' {
        leds += 5;
      } else if numbers[j] == '6' ||
                numbers[j] == '9' ||
                numbers[j] == '0' {
        leds += 6;
      } else if numbers[j] == '7' {
        leds += 3;
      } else if numbers[j] == '8' {
        leds += 7;
      } else if numbers[j] == '4' {
        leds += 4;
      }
    }
    fmt.Printf("%d leds\n", leds)
    leds = 0
  }
}

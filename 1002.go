package main
 
import (
    "fmt"
)

func URI1002() {
    var R, area float64

    fmt.Scanf("%f", &R)
    area = 3.14159 * (R*R)
    
    fmt.Printf("A=%.4f\n", area)
}
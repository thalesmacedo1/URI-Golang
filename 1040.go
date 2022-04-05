package main

import "fmt"

func URI1040() {
  var gradeA,
      gradeB, 
      gradeC,
      gradeD, mean, finalTestGrade float32

  fmt.Scanf("%f", &gradeA)
  fmt.Scanf("%f", &gradeB)
  fmt.Scanf("%f", &gradeC)
  fmt.Scanf("%f", &gradeD)
  
  mean += (gradeA*2 + gradeB*3 + gradeC*4 + gradeD*1) / 10.0
  
  fmt.Printf("Media: %.1f\n", mean)
  if mean > 7.0 {
    fmt.Printf("Aluno aprovado.\n")
  } else if mean < 5.0 {
    fmt.Printf("Aluno reprovado.\n")    
  } else {
    fmt.Printf("Aluno em exame.\n")
    fmt.Scanf("%f", &finalTestGrade)
    fmt.Printf("Nota do exame: %.1f\n", finalTestGrade)
    if mean >= 5.0 {
      fmt.Printf("Aluno aprovado.\n")
    } else {
      fmt.Printf("Aluno reprovado.\n")
    }
  mean = (mean + finalTestGrade) / 2
  fmt.Printf("Media final: %.1f\n", mean)
  }
}

package main

import "fmt"

func URI1114() {
  var password int

  for password != 2002 {
    fmt.Scanf("%d", &password)
    if password == 2002 {
      fmt.Printf("Acesso Permitido\n")
    } else {
      fmt.Printf("Senha Invalida\n")
    }
  }
}

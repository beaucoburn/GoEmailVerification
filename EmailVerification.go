package main

import (
    "fmt"
    emailverifier "github.com/AfterShip/email-verifier"
)

var (
  verifier = emailverifier.NewVerifier()
)

func main() {
  email := "example@exampledomain.org"

  ret, err := verifier.Verify(email)
  if err != nil {
    fmt.Println("Verify email address failed, error is: ", err)
    return
  }
  if !ret.Syntax.Valid{
    fmt.Println("email validation result", ret)
  }
}

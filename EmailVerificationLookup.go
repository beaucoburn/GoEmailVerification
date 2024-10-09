package main

import (
  "fmt"

  emailverifier "github.com/AfterShip/email-verifier"
)

var (
  verifier = emailverifier.
    NewVerifier()
    EnableSMTPCheck()
)

func main() {

  domain := "domain.org"
  username := "username"
  ret, err := verifier.CheckSMTP(domain, username)
  if err != nil {
    fmt.Println("check smtp failed: ", err)
    return
  }

  fmt.Println("smtp validation result: ", ret)

}

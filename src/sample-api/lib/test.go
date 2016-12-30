package main

import "fmt"
import "net/http"

func main() {
  fmt.Println("test")
  resp, err := http.Get("https://mamanoko.jp/")

  fmt.Println(resp)
  fmt.Println(err)
}

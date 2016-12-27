package main

import "fmt"

func main()  {
  var a string = "initial"
  fmt.Println(a)

  var b = "infer types"
  fmt.Println(b)

  var c, d int = 1, 2
  fmt.Println(c, d)

  var e int
  fmt.Println(e)

  f := "declared and initialised in one go"

  fmt.Println(f)
}

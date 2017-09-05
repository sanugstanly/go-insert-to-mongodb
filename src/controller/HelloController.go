package controller

import(
  "net/http"
  "fmt"
)

func PrintHelloWorld (w http.ResponseWriter, r *http.Request)  {
  w.Write([]byte("Hello World!"))
}


func Test()  {
  fmt.Print("test")
}

package controller

import(
  "net/http"
)

func PrintHelloWorld (w http.ResponseWriter, r *http.Request)  {
  w.Write([]byte("Hello World!"))
}

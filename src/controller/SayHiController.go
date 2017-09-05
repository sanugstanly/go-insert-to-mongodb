package controller

import(
  "net/http"
  "input"
)

func SayHi(w http.ResponseWriter, r *http.Request)  {
  params := input.GetParams(r)
  w.Write([]byte("Hi "+params["name"]))
}

package server

import (
  "net/http"
  "routes"
  "fmt"
)

func Start()  {
  err := http.ListenAndServe(":9080", routes.AppRoutes)
  fmt.Print(err)
}

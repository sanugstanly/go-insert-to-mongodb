package app

import (
  "routes"
  "net/http"
  "fmt"
)

func Start()  {
  router := routes.RenderRoutes()
  err := http.ListenAndServe(":9080", router)
  fmt.Print(err)
}

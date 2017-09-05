package app

import (
  "routes"
  "server"
)

func Start()  {
  routes.RenderRoutes()
  server.Start()
}

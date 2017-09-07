package routes

import (
  "controller"
  "model"
)

var routes = model.Routes{
    model.Route{"hello", "GET", "/", controller.PrintHelloWorld,},
    model.Route{"hi", "GET", "/hi/{name}", controller.SayHi,},
}

package routes

import(
  "controller"
  "github.com/gorilla/mux"
)

var AppRoutes = mux.NewRouter()

type Route struct {
    Url        string `json:"url"`
    Function   string `json:"function"`
}

func RenderRoutes()  {
  AppRoutes.HandleFunc("/", controller.PrintHelloWorld)
  AppRoutes.HandleFunc("/hi/{name}", controller.SayHi)
}

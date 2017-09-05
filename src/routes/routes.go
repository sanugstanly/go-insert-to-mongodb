package routes

import(
  "controller"
  "github.com/gorilla/mux"
  "os"
  "encoding/json"
  "io/ioutil"
  "fmt"
  "net/http"
)

var AppRoutes = mux.NewRouter()

type Route struct {
    Url        string `json:"url"`
    Function   func(w http.ResponseWriter, r *http.Request) `json:"function"`
}

func RenderRoutes()  {
  allroutes := getRoutes()
  for index := 0; index < len(allroutes); index++ {
    // addHandle(allroutes[index].Url, allroutes[index].Function)
    // AppRoutes.HandleFunc("/", controller.PrintHelloWorld)
    controller.Test()
    AppRoutes.HandleFunc(allroutes[index].Url, allroutes[index].Function)
  }
}

// func addhandle(string url, string crtl)  {
//
// }

func getRoutes() []Route {
  raw, err := ioutil.ReadFile("./src/routes/routes.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    var route []Route
    json.Unmarshal(raw, &route)

    return route
}

package main

import (
	"fmt"
	"net/http"
    "log/slog"
    "html/template"
)

type Obj struct {
    ID int
    Name string
}

func main() {
    
    http.HandleFunc("/", Home)

    http.HandleFunc("/test", Test)

    err := http.ListenAndServe(":3000", nil)
    
    if err != nil {
        slog.Debug("error starting server", err)
    }

    fmt.Println("********** Server Ready **********")

}

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in home")
    pageTemplate := template.New("Home")
    pageTemplate, err := pageTemplate.ParseFiles("views/index.html")
    //pageTemplate, err := pageTemplate.Parse(`<p>{{ .ID }}, {{ .Name }}</p>`)
    if err != nil {
        fmt.Println("error")
        slog.Debug("error parsing template", err)
    }

    newObj := Obj{ID: 3, Name:"Test"}

    err = pageTemplate.ExecuteTemplate(w, "Index", newObj)
    if err != nil {
        slog.Debug("error executing template", err)
    }
    fmt.Println("page should have executed")
}

func Test(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "views/index.html")
}



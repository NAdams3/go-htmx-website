package main

import (
	"fmt"
	"net/http"
    "log/slog"
    "html/template"
)

type header struct {
    Title string
}

type footer struct {
    Year string 
}   

type pageData struct {
    Header header
    Content interface{}
    Footer footer
}

var templates map[string]*template.Template 

func main() {

    FilterTest()

    var err error
    
    templateDefs := []string{"page", "home", "contact"}
    templates = make(map[string]*template.Template)

    // parse html
    for i, name := range templateDefs {
        if i == 0 {
            templates[name] = template.New(name)
            templates[name] = getTemplate(*templates[name], "views/parts/"+name+".html")
        }
        if templates[name] == nil {
            templates[name] = getTemplate(*templates["page"], "views/"+name+".html")
        }
    }

    // routes
    http.HandleFunc("/", Home)
    http.HandleFunc("/contact", Contact)

    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("/views/assets"))))

    err = http.ListenAndServe(":3000", nil)
    
    if err != nil {
        slog.Debug("error starting server", err)
        panic(err)
    }

    fmt.Println("********** Server Ready **********")

}

func Render(w http.ResponseWriter, pageTemplate *template.Template, data any) {
    err := pageTemplate.ExecuteTemplate(w, "page", data)
    if err != nil {
        slog.Debug("error executing template", err)
        panic(err)
    }   
}

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Println("in home")

    data := getPageData()

    Render(w, templates["home"], data)
}

func Contact(w http.ResponseWriter, r *http.Request) {
    Render(w, templates["contact"], nil)
}

func getTemplate(t template.Template, filePath string) *template.Template {
    templatePart, err := t.ParseFiles(filePath)
    if err != nil {
        slog.Debug("error parsing template file", err)
        panic(err)
    }
    
    return templatePart  
}

func getPageData() *pageData {
    header := header{Title: "title"}
    content := struct{HTML string}{HTML: `<p>this is my first htmx site.</p>`}
    footer := footer{Year: "2024"}

    data := pageData{Header: header, Content: content, Footer: footer}

    return &data
}


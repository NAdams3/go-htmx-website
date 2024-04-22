package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

type header struct {
	Title string
}

type footer struct {
	Year string
}

type pageData struct {
	Header  header
	Content interface{}
	Footer  footer
}

var templates map[string]*template.Template

func main() {

	templateDefs := []string{"home", "contact"}
	templates = make(map[string]*template.Template)

	// parse html
	for _, name := range templateDefs {
		if templates[name] == nil {
			templates[name] = template.New(name)
			templates[name] = getTemplate(templates[name], "views/parts/page.html", "views/"+name+".html")
		}
	}

	// routes
	http.HandleFunc("/", Home)
	http.HandleFunc("/contact", Contact)
	http.HandleFunc("/contactSubmit", ContactSubmit)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("views/assets/"))))

	err := http.ListenAndServe(":3000", nil)

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

func ContactSubmit(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		slog.Error("error parsing form data", err)
	}

	fmt.Println("name: ", r.Form["name"])
	fmt.Println("email: ", r.Form["email"])
	fmt.Println("message: ", r.Form["message"])

	fmt.Println("in contact submit")

	err = sendEmail(r.Form["name"][0], r.Form["email"][0], "", "")
	if err != nil {
		slog.Error("error sending email", err)
	}

	http.ServeFile(w, r, "views/parts/contact-form.html")
}

func getTemplate(t *template.Template, filePaths ...string) *template.Template {
	templatePart, err := t.ParseFiles(filePaths...)
	if err != nil {
		slog.Debug("error parsing template file", err)
		panic(err)
	}

	return templatePart
}

func getPageData() *pageData {
	header := header{Title: "title"}
	content := struct{ HTML string }{HTML: `<p>this is my first htmx site.</p>`}
	footer := footer{Year: "2024"}

	data := pageData{Header: header, Content: content, Footer: footer}

	return &data
}

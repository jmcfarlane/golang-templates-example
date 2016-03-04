package main

//go:generate go-bindata-assetfs static/... templates/...

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Model of stuff to render a page
type Model struct {
	Title string
	Name  string
}

// Templates with functions available to them
var templates = template.New("").Funcs(templateMap)

// Parse all of the bindata templates
func init() {
	for _, path := range AssetNames() {
		bytes, err := Asset(path)
		if err != nil {
			log.Panicf("Unable to parse: path=%s, err=%s", path, err)
		}
		templates.New(path).Parse(string(bytes))
	}
}

// Render a template given a model
func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Well hello there
func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	model := Model{Name: ps.ByName("name")}
	renderTemplate(w, "templates/hello.html", &model)
}

// The server itself
func main() {
	// mux handler
	router := httprouter.New()

	// Example route that takes one rest style option
	router.GET("/hello/:name", hello)

	// Serve static assets via the "static" directory
	router.ServeFiles("/static/*filepath", assetFS())

	// Serve this program forever
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

//go:generate rice embed-go

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	rice "github.com/GeertJohan/go.rice"
	"github.com/julienschmidt/httprouter"
)

// Model of stuff to render a page
type Model struct {
	Title string
	Name  string
}

// Templates with functions available to them
var (
	templateMap = template.FuncMap{
		"Upper": func(s string) string {
			return strings.ToUpper(s)
		},
	}
	templates   = template.New("").Funcs(templateMap)
	templateBox *rice.Box
)

func newTemplate(path string, _ os.FileInfo, _ error) error {
	if path == "" {
		return nil
	}
	templateString, err := templateBox.String(path)
	if err != nil {
		log.Panicf("Unable to extract: path=%s, err=%s", path, err)
	}
	if _, err = templates.New(filepath.Join("templates", path)).Parse(templateString); err != nil {
		log.Panicf("Unable to parse: path=%s, err=%s", path, err)
	}
	return nil
}

// Render a template given a model
func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	err := templates.ExecuteTemplate(w, tmpl, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func broken(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	renderTemplate(w, "templates/missing.html", nil)
}

// Well hello there
func hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	model := Model{Name: ps.ByName("name")}
	renderTemplate(w, "templates/hello.html", &model)
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	renderTemplate(w, "templates/index.html", nil)
}

func getRouter() *httprouter.Router {
	// Load and parse templates (from binary or disk)
	templateBox = rice.MustFindBox("templates")
	templateBox.Walk("", newTemplate)

	// mux handler
	router := httprouter.New()

	// Index routee
	router.GET("/", index)

	// Example route that takes one rest style option
	router.GET("/hello/:name", hello)

	// Example route that encounters an error
	router.GET("/broken/handler", broken)

	// Serve static assets via the "static" directory
	fs := rice.MustFindBox("static").HTTPBox()
	router.ServeFiles("/static/*filepath", fs)
	return router
}

func main() {
	listen := flag.String("-listen", ":8080", "Interface and port to listen on")
	flag.Parse()
	fmt.Println("Listening on", *listen)
	log.Fatal(http.ListenAndServe(*listen, getRouter()))
}

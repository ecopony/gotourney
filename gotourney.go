package main

import (
	"html/template"
	"net/http"
	"sync"
	"path/filepath"
	"log"
)

func main() {
	http.Handle("/", &templateHandler{filename: "index.html"})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Listen and serve: ", err)
	}
}

type templateHandler struct {
	once     sync.Once
	filename string
	tmpl    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.tmpl.Execute(w, &PageData{Title: "gotourney", Username: "ecopony"})
}

type PageData struct {
	Title    string
	Username string
}

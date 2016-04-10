package main

import (
	"html/template"
	"sync"
	"net/http"
	"fmt"
	"path/filepath"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tmpl    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		fmt.Println("compile template")
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	fmt.Println("serve template")
	t.tmpl.Execute(w, &PageData{Title: "gotourney", Username: "ecopony"})
}

type PageData struct {
	Title    string
	Username string
}

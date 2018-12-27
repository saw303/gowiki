package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const extension = ".wiki"

type Page struct {
	Title string
	Body  string
}

func (p *Page) save() error {
	filename := p.Title + extension
	return ioutil.WriteFile(filename, []byte(p.Body+" "+time.Now().UTC().Format("02.01.2006 15:04:05")+" (UTC)"), 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + extension
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: string(body)}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	page, _ := loadPage("TestPage")
	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintln(w, string(page.Body))
}

func main() {
	p1 := &Page{Title: "TestPage", Body: "This is a sample Page."}
	p1.save()

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

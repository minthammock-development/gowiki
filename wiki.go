package main

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	fileName := p.Title + ".txt"
	return os.WriteFile(fileName, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	fileName := title + ".txt"
	body, E := os.ReadFile(fileName)
	if E != nil {
		return nil, E
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
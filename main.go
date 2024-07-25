package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {

	deck := newDeck()

	deck.shuffle()

	deck.print()

	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", deck)
	if err != nil {
		log.Fatalln(err)
	}

	//h1 := func(w http.ResponseWriter, r *http.Request) {
	//	tmpl.Execute(w, tmpl)
	//}

	//http.HandleFunc("/", h1)

	//log.Fatal(http.ListenAndServe(":8000", nil))
}

//h1 := func(w http.ResponseWriter, r *http.Request) {
//	tmpl := template.Must(template.ParseFiles("index.html"))
//	currentDeck := deck
//	tmpl.Execute(w, currentDeck)
//}

//h1 := func(w http.ResponseWriter, r *http.Request) {
//	tmpl.Execute(w, deck)
//}
//
//http.HandleFunc("/", h1)

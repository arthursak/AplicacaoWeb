package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quant           int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{
			Nome:      "Camiseta",
			Descricao: "Azul, bem bonita",
			Preco:     39,
			Quant:     5},
		{
			Nome:      "Tênis",
			Descricao: "Confortavel",
			Preco:     150,
			Quant:     3},

		{
			Nome:      "Fone",
			Descricao: "Very good",
			Preco:     90,
			Quant:     2},
	}

	temp.ExecuteTemplate(w, "index", produtos)
}

package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

func conectaBD() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=Tutu@0303 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quant, Id       int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}

func index(w http.ResponseWriter, r *http.Request) {

	db := conectaBD()

	selectProd, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProd.Next() {
		var nome, descricao string
		var id, quant int
		var preco float64

		err = selectProd.Scan(&id, &nome, &descricao, &preco, &quant)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Quant = quant
		p.Preco = preco

		produtos = append(produtos, p)

	}

	temp.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
}

package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
} 

var temp = template.Must(template.ParseGlob("templates/*.html"))


func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome:"Camiseta", Descricao: "Azul", Preco:39, Quantidade:5},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Bom", 59, 2},
		{"Prida", "Bomadad", 5911, 22},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
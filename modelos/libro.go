package modelos

import "fmt"

type Libro struct {
	ID         int    `json:"id"`
	Titulo     string `json:"titulo"`
	Autor      string `json:"autor"`
	Categoria  string `json:"categoria"`
	Disponible bool   `json:"disponible"`
}

func (l Libro) Mostrar() {
	fmt.Printf(
		"ID:%d | %s | %s | %s | Disponible:%t\n",
		l.ID,
		l.Titulo,
		l.Autor,
		l.Categoria,
		l.Disponible,
	)
}

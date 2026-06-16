package modelos

import "fmt"

type Libro struct {
	ID         int
	Titulo     string
	Autor      string
	Categoria  string
	Disponible bool
}

// Implementa la interfaz Mostrable
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

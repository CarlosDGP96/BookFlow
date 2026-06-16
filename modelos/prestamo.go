package modelos

import "fmt"

type Prestamo struct {
	ID        int
	IDUsuario int
	IDLibro   int
	Estado    string
}

// Implementa la interfaz Mostrable
func (p Prestamo) Mostrar() {

	fmt.Printf(
		"ID:%d | Usuario:%d | Libro:%d | Estado:%s\n",
		p.ID,
		p.IDUsuario,
		p.IDLibro,
		p.Estado,
	)
}

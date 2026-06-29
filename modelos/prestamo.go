package modelos

import "fmt"

type Prestamo struct {
	ID        int    `json:"id"`
	IDUsuario int    `json:"id_usuario"`
	IDLibro   int    `json:"id_libro"`
	Estado    string `json:"estado"`
}

func (p Prestamo) Mostrar() {
	fmt.Printf(
		"ID:%d | Usuario:%d | Libro:%d | Estado:%s\n",
		p.ID,
		p.IDUsuario,
		p.IDLibro,
		p.Estado,
	)
}

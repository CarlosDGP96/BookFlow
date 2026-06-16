package modelos

import "fmt"

type Usuario struct {
	ID     int
	Nombre string
	Correo string
}

// Implementa la interfaz Mostrable
func (u Usuario) Mostrar() {
	fmt.Printf(
		"ID:%d | Nombre:%s | Correo:%s\n",
		u.ID,
		u.Nombre,
		u.Correo,
	)
}

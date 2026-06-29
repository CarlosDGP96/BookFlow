package modelos

import "fmt"

type Usuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Correo string `json:"correo"`
}

func (u Usuario) Mostrar() {
	fmt.Printf("ID:%d | Nombre:%s | Correo:%s\n", u.ID, u.Nombre, u.Correo)
}

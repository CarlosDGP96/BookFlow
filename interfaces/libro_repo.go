package interfaces

import "bookflow/modelos"

type LibroRepositorio interface {
	Crear(libro modelos.Libro) error
	Listar() []modelos.Libro
	Buscar(id int) (modelos.Libro, error)
	Actualizar(libro modelos.Libro) error
}

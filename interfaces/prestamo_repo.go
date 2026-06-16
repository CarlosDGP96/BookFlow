package interfaces

import "bookflow/modelos"

type PrestamoRepositorio interface {
	Crear(prestamo modelos.Prestamo) error
	Listar() []modelos.Prestamo
}

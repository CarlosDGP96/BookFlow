package interfaces

import "bookflow/modelos"

type UsuarioRepositorio interface {
	Crear(usuario modelos.Usuario) error
	Listar() []modelos.Usuario
	Buscar(id int) (modelos.Usuario, error)
}

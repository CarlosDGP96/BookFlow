package repositorios

import (
	"bookflow/modelos"
	"errors"
)

type UsuarioMemoria struct {
	usuarios []modelos.Usuario
}

func NuevoUsuarioMemoria() *UsuarioMemoria {
	return &UsuarioMemoria{
		usuarios: []modelos.Usuario{},
	}
}

func (u *UsuarioMemoria) Crear(usuario modelos.Usuario) error {
	u.usuarios = append(u.usuarios, usuario)
	return nil
}

func (u *UsuarioMemoria) Listar() []modelos.Usuario {
	return u.usuarios
}

// Buscar localiza un usuario utilizando su ID.
func (u *UsuarioMemoria) Buscar(id int) (modelos.Usuario, error) {

	for _, usuario := range u.usuarios {

		if usuario.ID == id {
			return usuario, nil
		}
	}

	return modelos.Usuario{}, errors.New("usuario no encontrado")
}

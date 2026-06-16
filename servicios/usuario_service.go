package servicios

import (
	"bookflow/interfaces"
	"bookflow/modelos"
	"errors"
)

type UsuarioService struct {
	repo interfaces.UsuarioRepositorio
}

// Constructor del servicio de usuarios.
func NuevoUsuarioService(repo interfaces.UsuarioRepositorio) *UsuarioService {
	return &UsuarioService{
		repo: repo,
	}
}

// CrearUsuario valida la información antes de almacenarla.
func (s *UsuarioService) CrearUsuario(
	id int,
	nombre string,
	correo string,
) error {

	if nombre == "" {
		return errors.New("el nombre no puede estar vacío")
	}

	if correo == "" {
		return errors.New("el correo no puede estar vacío")
	}

	usuario := modelos.Usuario{
		ID:     id,
		Nombre: nombre,
		Correo: correo,
	}

	return s.repo.Crear(usuario)
}

// ListarUsuarios retorna todos los usuarios registrados.
func (s *UsuarioService) ListarUsuarios() []modelos.Usuario {
	return s.repo.Listar()
}

// BuscarUsuario busca un usuario por ID.
func (s *UsuarioService) BuscarUsuario(id int) (modelos.Usuario, error) {
	return s.repo.Buscar(id)
}

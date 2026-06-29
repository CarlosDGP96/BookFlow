package servicios

import (
	"bookflow/interfaces"
	"bookflow/modelos"
	"errors"
)

type UsuarioService struct {
	repo interfaces.UsuarioRepositorio
}

func NuevoUsuarioService(repo interfaces.UsuarioRepositorio) *UsuarioService {
	return &UsuarioService{repo: repo}
}

func (s *UsuarioService) CrearUsuario(id int, nombre string, correo string) error {
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

func (s *UsuarioService) ListarUsuarios() []modelos.Usuario {
	return s.repo.Listar()
}

func (s *UsuarioService) BuscarUsuario(id int) (modelos.Usuario, error) {
	return s.repo.Buscar(id)
}

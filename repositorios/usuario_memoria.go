package repositorios

import (
	"bookflow/modelos"
	"errors"
	"sync"
)

type UsuarioMemoria struct {
	usuarios []modelos.Usuario
	mu       sync.RWMutex
}

func NuevoUsuarioMemoria() *UsuarioMemoria {
	return &UsuarioMemoria{
		usuarios: []modelos.Usuario{},
	}
}

func (u *UsuarioMemoria) Crear(usuario modelos.Usuario) error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.usuarios = append(u.usuarios, usuario)
	return nil
}

func (u *UsuarioMemoria) Listar() []modelos.Usuario {
	u.mu.RLock()
	defer u.mu.RUnlock()

	copia := make([]modelos.Usuario, len(u.usuarios))
	copy(copia, u.usuarios)

	return copia
}

func (u *UsuarioMemoria) Buscar(id int) (modelos.Usuario, error) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	for _, usuario := range u.usuarios {
		if usuario.ID == id {
			return usuario, nil
		}
	}

	return modelos.Usuario{}, errors.New("usuario no encontrado")
}

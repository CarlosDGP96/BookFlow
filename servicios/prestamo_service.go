package servicios

import (
	"bookflow/interfaces"
	"bookflow/modelos"
	"errors"
)

type PrestamoService struct {
	repo interfaces.PrestamoRepositorio
}

// Constructor del servicio de préstamos.
func NuevoPrestamoService(
	repo interfaces.PrestamoRepositorio,
) *PrestamoService {

	return &PrestamoService{
		repo: repo,
	}
}

// CrearPrestamo registra un nuevo préstamo.
func (s *PrestamoService) CrearPrestamo(
	id int,
	idUsuario int,
	idLibro int,
) error {

	if idUsuario <= 0 {
		return errors.New("id de usuario inválido")
	}

	if idLibro <= 0 {
		return errors.New("id de libro inválido")
	}

	prestamo := modelos.Prestamo{
		ID:        id,
		IDUsuario: idUsuario,
		IDLibro:   idLibro,
		Estado:    "Activo",
	}

	return s.repo.Crear(prestamo)
}

// ListarPrestamos retorna todos los préstamos.
func (s *PrestamoService) ListarPrestamos() []modelos.Prestamo {
	return s.repo.Listar()
}

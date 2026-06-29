package servicios

import (
	"bookflow/interfaces"
	"bookflow/modelos"
	"errors"
)

type LibroService struct {
	repo interfaces.LibroRepositorio
}

func NuevoLibroService(repo interfaces.LibroRepositorio) *LibroService {
	return &LibroService{repo: repo}
}

func (s *LibroService) CrearLibro(id int, titulo string, autor string, categoria string) error {
	if titulo == "" {
		return errors.New("el título no puede estar vacío")
	}

	if autor == "" {
		return errors.New("el autor no puede estar vacío")
	}

	libro := modelos.Libro{
		ID:         id,
		Titulo:     titulo,
		Autor:      autor,
		Categoria:  categoria,
		Disponible: true,
	}

	return s.repo.Crear(libro)
}

func (s *LibroService) ListarLibros() []modelos.Libro {
	return s.repo.Listar()
}

func (s *LibroService) BuscarLibro(id int) (modelos.Libro, error) {
	return s.repo.Buscar(id)
}

func (s *LibroService) MarcarNoDisponible(id int) error {
	libro, err := s.repo.Buscar(id)

	if err != nil {
		return err
	}

	libro.Disponible = false

	return s.repo.Actualizar(libro)
}

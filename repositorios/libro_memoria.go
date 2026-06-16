package repositorios

import (
	"bookflow/modelos"
	"errors"
)

type LibroMemoria struct {
	libros []modelos.Libro
}

func NuevoLibroMemoria() *LibroMemoria {
	return &LibroMemoria{
		libros: []modelos.Libro{},
	}
}

func (l *LibroMemoria) Crear(libro modelos.Libro) error {
	l.libros = append(l.libros, libro)
	return nil
}

func (l *LibroMemoria) Listar() []modelos.Libro {
	return l.libros
}

// Buscar localiza un libro utilizando su ID.
func (l *LibroMemoria) Buscar(id int) (modelos.Libro, error) {

	for i := range l.libros {

		if l.libros[i].ID == id {
			return l.libros[i], nil
		}
	}

	return modelos.Libro{}, errors.New("libro no encontrado")
}

// Actualizar reemplaza un libro existente.
func (l *LibroMemoria) Actualizar(libro modelos.Libro) error {

	for i := range l.libros {

		if l.libros[i].ID == libro.ID {

			l.libros[i] = libro
			return nil
		}
	}

	return errors.New("libro no encontrado")
}

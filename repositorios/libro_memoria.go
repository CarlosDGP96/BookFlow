package repositorios

import (
	"bookflow/modelos"
	"errors"
	"sync"
)

type LibroMemoria struct {
	libros []modelos.Libro
	mu     sync.RWMutex
}

func NuevoLibroMemoria() *LibroMemoria {
	return &LibroMemoria{
		libros: []modelos.Libro{},
	}
}

func (l *LibroMemoria) Crear(libro modelos.Libro) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.libros = append(l.libros, libro)
	return nil
}

func (l *LibroMemoria) Listar() []modelos.Libro {
	l.mu.RLock()
	defer l.mu.RUnlock()

	copia := make([]modelos.Libro, len(l.libros))
	copy(copia, l.libros)

	return copia
}

func (l *LibroMemoria) Buscar(id int) (modelos.Libro, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for _, libro := range l.libros {
		if libro.ID == id {
			return libro, nil
		}
	}

	return modelos.Libro{}, errors.New("libro no encontrado")
}

func (l *LibroMemoria) Actualizar(libro modelos.Libro) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	for i := range l.libros {
		if l.libros[i].ID == libro.ID {
			l.libros[i] = libro
			return nil
		}
	}

	return errors.New("libro no encontrado")
}

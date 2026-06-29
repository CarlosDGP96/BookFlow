package repositorios

import (
	"bookflow/modelos"
	"sync"
)

type PrestamoMemoria struct {
	prestamos []modelos.Prestamo
	mu        sync.RWMutex
}

func NuevoPrestamoMemoria() *PrestamoMemoria {
	return &PrestamoMemoria{
		prestamos: []modelos.Prestamo{},
	}
}

func (p *PrestamoMemoria) Crear(prestamo modelos.Prestamo) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.prestamos = append(p.prestamos, prestamo)
	return nil
}

func (p *PrestamoMemoria) Listar() []modelos.Prestamo {
	p.mu.RLock()
	defer p.mu.RUnlock()

	copia := make([]modelos.Prestamo, len(p.prestamos))
	copy(copia, p.prestamos)

	return copia
}

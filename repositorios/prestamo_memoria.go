package repositorios

import (
	"bookflow/modelos"
)

type PrestamoMemoria struct {
	prestamos []modelos.Prestamo
}

func NuevoPrestamoMemoria() *PrestamoMemoria {
	return &PrestamoMemoria{
		prestamos: []modelos.Prestamo{},
	}
}

func (p *PrestamoMemoria) Crear(prestamo modelos.Prestamo) error {
	p.prestamos = append(p.prestamos, prestamo)
	return nil
}

func (p *PrestamoMemoria) Listar() []modelos.Prestamo {
	return p.prestamos
}

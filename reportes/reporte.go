package reportes

type Resumen struct {
	TotalUsuarios  int `json:"total_usuarios"`
	TotalLibros    int `json:"total_libros"`
	TotalPrestamos int `json:"total_prestamos"`
}

func CrearResumen(totalUsuarios int, totalLibros int, totalPrestamos int) Resumen {
	return Resumen{
		TotalUsuarios:  totalUsuarios,
		TotalLibros:    totalLibros,
		TotalPrestamos: totalPrestamos,
	}
}

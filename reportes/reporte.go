package reportes

import "fmt"

func MostrarResumen(
	totalUsuarios int,
	totalLibros int,
	totalPrestamos int,
) {

	fmt.Println("\n===== REPORTE GENERAL =====")

	fmt.Printf("Usuarios registrados: %d\n", totalUsuarios)
	fmt.Printf("Libros registrados: %d\n", totalLibros)
	fmt.Printf("Préstamos registrados: %d\n", totalPrestamos)
}

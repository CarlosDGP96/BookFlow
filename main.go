package main

import (
	"fmt"
	"bookflow/usuarios"
	"bookflow/libros"
	"bookflow/prestamos"
	"bookflow/reportes"
)

func main() {
	fmt.Println("===================================")
	fmt.Println("     SISTEMA BOOKFLOW INICIADO     ")
	fmt.Println("===================================")

	// Inicialización básica de datos (en memoria)
	userService := usuarios.NewServicioUsuario()
	bookService := libros.NewServicioLibro()
	loanService := prestamos.NewServicioPrestamo()

	// Datos de prueba iniciales
	userService.CrearUsuario("Carlos", "carlos@mail.com")
	bookService.CrearLibro("Clean Code", "Robert Martin", "Programación")

	// Ejemplo de préstamo
	loanService.PrestarLibro(1, 1)

	// Reporte simple
	reportes.MostrarResumen(userService, bookService, loanService)
}
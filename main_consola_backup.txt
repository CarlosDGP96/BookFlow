package main

import (
	"bookflow/reportes"
	"bookflow/repositorios"
	"bookflow/servicios"
	"fmt"
)

func main() {

	usuarioRepo := repositorios.NuevoUsuarioMemoria()
	libroRepo := repositorios.NuevoLibroMemoria()
	prestamoRepo := repositorios.NuevoPrestamoMemoria()

	usuarioService := servicios.NuevoUsuarioService(usuarioRepo)
	libroService := servicios.NuevoLibroService(libroRepo)
	prestamoService := servicios.NuevoPrestamoService(prestamoRepo)

	var opcion int

	for {

		fmt.Println("\n===== BOOKFLOW =====")
		fmt.Println("1. Crear usuario")
		fmt.Println("2. Listar usuarios")
		fmt.Println("3. Buscar usuario")
		fmt.Println("4. Registrar libro")
		fmt.Println("5. Listar libros")
		fmt.Println("6. Buscar libro")
		fmt.Println("7. Registrar préstamo")
		fmt.Println("8. Ver préstamos")
		fmt.Println("9. Reporte general")
		fmt.Println("0. Salir")

		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {

		case 1:

			var nombre string
			var correo string

			fmt.Print("Nombre: ")
			fmt.Scanln(&nombre)

			fmt.Print("Correo: ")
			fmt.Scanln(&correo)

			id := len(usuarioService.ListarUsuarios()) + 1

			err := usuarioService.CrearUsuario(
				id,
				nombre,
				correo,
			)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Usuario creado correctamente")
			}

		case 2:

			usuarios := usuarioService.ListarUsuarios()

			fmt.Println("\n===== USUARIOS =====")

			if len(usuarios) == 0 {
				fmt.Println("No existen usuarios")
			}

			for _, usuario := range usuarios {

				fmt.Printf(
					"ID:%d | Nombre:%s | Correo:%s\n",
					usuario.ID,
					usuario.Nombre,
					usuario.Correo,
				)
			}

		case 3:

			var id int

			fmt.Print("Ingrese ID del usuario: ")
			fmt.Scanln(&id)

			usuario, err := usuarioService.BuscarUsuario(id)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				fmt.Println("\nUsuario encontrado")

				fmt.Printf(
					"ID:%d | Nombre:%s | Correo:%s\n",
					usuario.ID,
					usuario.Nombre,
					usuario.Correo,
				)
			}

		case 4:

			var titulo string
			var autor string
			var categoria string

			fmt.Print("Título: ")
			fmt.Scanln(&titulo)

			fmt.Print("Autor: ")
			fmt.Scanln(&autor)

			fmt.Print("Categoría: ")
			fmt.Scanln(&categoria)

			id := len(libroService.ListarLibros()) + 1

			err := libroService.CrearLibro(
				id,
				titulo,
				autor,
				categoria,
			)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Libro registrado correctamente")
			}

		case 5:

			libros := libroService.ListarLibros()

			fmt.Println("\n===== LIBROS =====")

			if len(libros) == 0 {
				fmt.Println("No existen libros")
			}

			for _, libro := range libros {

				estado := "No"

				if libro.Disponible {
					estado = "Sí"
				}

				fmt.Printf(
					"ID:%d | %s | %s | %s | Disponible:%s\n",
					libro.ID,
					libro.Titulo,
					libro.Autor,
					libro.Categoria,
					estado,
				)
			}

		case 6:

			var id int

			fmt.Print("Ingrese ID del libro: ")
			fmt.Scanln(&id)

			libro, err := libroService.BuscarLibro(id)

			if err != nil {
				fmt.Println("Error:", err)
			} else {

				fmt.Println("\nLibro encontrado")

				fmt.Printf(
					"ID:%d | %s | %s | %s | Disponible:%t\n",
					libro.ID,
					libro.Titulo,
					libro.Autor,
					libro.Categoria,
					libro.Disponible,
				)
			}

		case 7:

			var idUsuario int
			var idLibro int

			fmt.Print("ID Usuario: ")
			fmt.Scanln(&idUsuario)

			fmt.Print("ID Libro: ")
			fmt.Scanln(&idLibro)

			_, errUsuario := usuarioService.BuscarUsuario(idUsuario)

			if errUsuario != nil {
				fmt.Println("Error:", errUsuario)
				break
			}

			libro, errLibro := libroService.BuscarLibro(idLibro)

			if errLibro != nil {
				fmt.Println("Error:", errLibro)
				break
			}

			if !libro.Disponible {
				fmt.Println("Error: el libro no está disponible")
				break
			}

			idPrestamo := len(prestamoService.ListarPrestamos()) + 1

			err := prestamoService.CrearPrestamo(
				idPrestamo,
				idUsuario,
				idLibro,
			)

			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			err = libroService.MarcarNoDisponible(idLibro)

			if err != nil {
				fmt.Println("Error:", err)
				break
			}

			fmt.Println("Préstamo registrado correctamente")

		case 8:

			prestamos := prestamoService.ListarPrestamos()

			fmt.Println("\n===== PRÉSTAMOS =====")

			if len(prestamos) == 0 {
				fmt.Println("No existen préstamos")
			}

			for _, prestamo := range prestamos {

				fmt.Printf(
					"ID:%d | Usuario:%d | Libro:%d | Estado:%s\n",
					prestamo.ID,
					prestamo.IDUsuario,
					prestamo.IDLibro,
					prestamo.Estado,
				)
			}

		case 9:

			reportes.MostrarResumen(
				len(usuarioService.ListarUsuarios()),
				len(libroService.ListarLibros()),
				len(prestamoService.ListarPrestamos()),
			)

		case 0:

			fmt.Println("Gracias por utilizar BookFlow")
			return

		default:

			fmt.Println("Opción inválida")
		}
	}
}

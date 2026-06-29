package controladores

import (
	"bookflow/modelos"
	"bookflow/reportes"
	"bookflow/repositorios"
	"bookflow/servicios"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type App struct {
	UsuarioService  *servicios.UsuarioService
	LibroService    *servicios.LibroService
	PrestamoService *servicios.PrestamoService
	mu              sync.Mutex
}

func NuevaApp() *App {
	usuarioRepo := repositorios.NuevoUsuarioMemoria()
	libroRepo := repositorios.NuevoLibroMemoria()
	prestamoRepo := repositorios.NuevoPrestamoMemoria()

	return &App{
		UsuarioService:  servicios.NuevoUsuarioService(usuarioRepo),
		LibroService:    servicios.NuevoLibroService(libroRepo),
		PrestamoService: servicios.NuevoPrestamoService(prestamoRepo),
	}
}

func responderJSON(w http.ResponseWriter, estado int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(estado)
	json.NewEncoder(w).Encode(data)
}

func responderError(w http.ResponseWriter, estado int, mensaje string) {
	responderJSON(w, estado, map[string]string{
		"error": mensaje,
	})
}

func (a *App) Inicio(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	t, err := template.ParseFiles("vistas/index.html")

	if err != nil {
		http.Error(w, "No se pudo cargar la vista", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

func (a *App) UsuariosAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		responderJSON(w, http.StatusOK, a.UsuarioService.ListarUsuarios())
		return
	}

	if r.Method == http.MethodPost {
		var usuario modelos.Usuario

		err := json.NewDecoder(r.Body).Decode(&usuario)

		if err != nil {
			responderError(w, http.StatusBadRequest, "json inválido")
			return
		}

		a.mu.Lock()
		defer a.mu.Unlock()

		id := len(a.UsuarioService.ListarUsuarios()) + 1

		err = a.UsuarioService.CrearUsuario(id, usuario.Nombre, usuario.Correo)

		if err != nil {
			responderError(w, http.StatusBadRequest, err.Error())
			return
		}

		creado, _ := a.UsuarioService.BuscarUsuario(id)
		responderJSON(w, http.StatusCreated, creado)
		return
	}

	responderError(w, http.StatusMethodNotAllowed, "método no permitido")
}

func (a *App) UsuarioPorIDAPI(w http.ResponseWriter, r *http.Request) {
	idTexto := strings.TrimPrefix(r.URL.Path, "/api/usuarios/")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	usuario, err := a.UsuarioService.BuscarUsuario(id)

	if err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, usuario)
}

func (a *App) LibrosAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		responderJSON(w, http.StatusOK, a.LibroService.ListarLibros())
		return
	}

	if r.Method == http.MethodPost {
		var libro modelos.Libro

		err := json.NewDecoder(r.Body).Decode(&libro)

		if err != nil {
			responderError(w, http.StatusBadRequest, "json inválido")
			return
		}

		a.mu.Lock()
		defer a.mu.Unlock()

		id := len(a.LibroService.ListarLibros()) + 1

		err = a.LibroService.CrearLibro(id, libro.Titulo, libro.Autor, libro.Categoria)

		if err != nil {
			responderError(w, http.StatusBadRequest, err.Error())
			return
		}

		creado, _ := a.LibroService.BuscarLibro(id)
		responderJSON(w, http.StatusCreated, creado)
		return
	}

	responderError(w, http.StatusMethodNotAllowed, "método no permitido")
}

func (a *App) LibroPorIDAPI(w http.ResponseWriter, r *http.Request) {
	idTexto := strings.TrimPrefix(r.URL.Path, "/api/libros/")
	id, err := strconv.Atoi(idTexto)

	if err != nil {
		responderError(w, http.StatusBadRequest, "id inválido")
		return
	}

	libro, err := a.LibroService.BuscarLibro(id)

	if err != nil {
		responderError(w, http.StatusNotFound, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, libro)
}

func (a *App) PrestamosAPI(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		responderJSON(w, http.StatusOK, a.PrestamoService.ListarPrestamos())
		return
	}

	if r.Method == http.MethodPost {
		var prestamo modelos.Prestamo

		err := json.NewDecoder(r.Body).Decode(&prestamo)

		if err != nil {
			responderError(w, http.StatusBadRequest, "json inválido")
			return
		}

		a.mu.Lock()
		defer a.mu.Unlock()

		_, errUsuario := a.UsuarioService.BuscarUsuario(prestamo.IDUsuario)

		if errUsuario != nil {
			responderError(w, http.StatusBadRequest, errUsuario.Error())
			return
		}

		libro, errLibro := a.LibroService.BuscarLibro(prestamo.IDLibro)

		if errLibro != nil {
			responderError(w, http.StatusBadRequest, errLibro.Error())
			return
		}

		if !libro.Disponible {
			responderError(w, http.StatusBadRequest, "el libro no está disponible")
			return
		}

		id := len(a.PrestamoService.ListarPrestamos()) + 1

		err = a.PrestamoService.CrearPrestamo(id, prestamo.IDUsuario, prestamo.IDLibro)

		if err != nil {
			responderError(w, http.StatusBadRequest, err.Error())
			return
		}

		err = a.LibroService.MarcarNoDisponible(prestamo.IDLibro)

		if err != nil {
			responderError(w, http.StatusBadRequest, err.Error())
			return
		}

		creado := modelos.Prestamo{
			ID:        id,
			IDUsuario: prestamo.IDUsuario,
			IDLibro:   prestamo.IDLibro,
			Estado:    "Activo",
		}

		responderJSON(w, http.StatusCreated, creado)
		return
	}

	responderError(w, http.StatusMethodNotAllowed, "método no permitido")
}

func (a *App) ReporteAPI(w http.ResponseWriter, r *http.Request) {
	resumen := reportes.CrearResumen(
		len(a.UsuarioService.ListarUsuarios()),
		len(a.LibroService.ListarLibros()),
		len(a.PrestamoService.ListarPrestamos()),
	)

	responderJSON(w, http.StatusOK, resumen)
}

func (a *App) DemoConcurrenciaAPI(w http.ResponseWriter, r *http.Request) {
	canal := make(chan string)

	go func() {
		canal <- "Solicitud procesada mediante goroutine"
	}()

	mensaje := <-canal

	responderJSON(w, http.StatusOK, map[string]string{
		"mensaje": mensaje,
	})
}

package rutas

import (
	"bookflow/controladores"
	"net/http"
)

func ConfigurarRutas(app *controladores.App) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.Inicio)

	mux.HandleFunc("/api/usuarios", app.UsuariosAPI)
	mux.HandleFunc("/api/usuarios/", app.UsuarioPorIDAPI)

	mux.HandleFunc("/api/libros", app.LibrosAPI)
	mux.HandleFunc("/api/libros/", app.LibroPorIDAPI)

	mux.HandleFunc("/api/prestamos", app.PrestamosAPI)

	mux.HandleFunc("/api/reportes", app.ReporteAPI)
	mux.HandleFunc("/api/concurrencia", app.DemoConcurrenciaAPI)

	return mux
}

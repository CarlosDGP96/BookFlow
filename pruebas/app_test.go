package pruebas

import (
	"bookflow/controladores"
	"bookflow/rutas"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrearUsuarioAPI(t *testing.T) {
	app := controladores.NuevaApp()
	router := rutas.ConfigurarRutas(app)

	body := bytes.NewBufferString(`{"nombre":"Carlos","correo":"carlos@mail.com"}`)

	req := httptest.NewRequest(http.MethodPost, "/api/usuarios", body)
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("se esperaba estado 201, se obtuvo %d", res.Code)
	}
}

func TestCrearLibroAPI(t *testing.T) {
	app := controladores.NuevaApp()
	router := rutas.ConfigurarRutas(app)

	body := bytes.NewBufferString(`{"titulo":"CleanCode","autor":"RobertMartin","categoria":"Programacion"}`)

	req := httptest.NewRequest(http.MethodPost, "/api/libros", body)
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("se esperaba estado 201, se obtuvo %d", res.Code)
	}
}

func TestPrestamoLibroNoDisponible(t *testing.T) {
	app := controladores.NuevaApp()
	router := rutas.ConfigurarRutas(app)

	usuario := bytes.NewBufferString(`{"nombre":"Carlos","correo":"carlos@mail.com"}`)
	reqUsuario := httptest.NewRequest(http.MethodPost, "/api/usuarios", usuario)
	reqUsuario.Header.Set("Content-Type", "application/json")
	resUsuario := httptest.NewRecorder()
	router.ServeHTTP(resUsuario, reqUsuario)

	libro := bytes.NewBufferString(`{"titulo":"CleanCode","autor":"RobertMartin","categoria":"Programacion"}`)
	reqLibro := httptest.NewRequest(http.MethodPost, "/api/libros", libro)
	reqLibro.Header.Set("Content-Type", "application/json")
	resLibro := httptest.NewRecorder()
	router.ServeHTTP(resLibro, reqLibro)

	prestamo1 := bytes.NewBufferString(`{"id_usuario":1,"id_libro":1}`)
	reqPrestamo1 := httptest.NewRequest(http.MethodPost, "/api/prestamos", prestamo1)
	reqPrestamo1.Header.Set("Content-Type", "application/json")
	resPrestamo1 := httptest.NewRecorder()
	router.ServeHTTP(resPrestamo1, reqPrestamo1)

	prestamo2 := bytes.NewBufferString(`{"id_usuario":1,"id_libro":1}`)
	reqPrestamo2 := httptest.NewRequest(http.MethodPost, "/api/prestamos", prestamo2)
	reqPrestamo2.Header.Set("Content-Type", "application/json")
	resPrestamo2 := httptest.NewRecorder()
	router.ServeHTTP(resPrestamo2, reqPrestamo2)

	if resPrestamo2.Code != http.StatusBadRequest {
		t.Errorf("se esperaba error 400 por libro no disponible, se obtuvo %d", resPrestamo2.Code)
	}
}

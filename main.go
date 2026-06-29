package main

import (
	"bookflow/controladores"
	"bookflow/rutas"
	"fmt"
	"net/http"
)

func main() {
	app := controladores.NuevaApp()
	router := rutas.ConfigurarRutas(app)

	fmt.Println("Servidor BookFlow iniciado en http://localhost:8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Error al iniciar servidor:", err)
	}
}

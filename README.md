# BookFlow Web

## Datos del Proyecto

**Nombre del proyecto:** BookFlow Web  
**Tipo de aplicación:** Aplicación web MVC con servicios API  
**Lenguaje:** Go  
**Serialización:** JSON  
**Fecha:** Junio 2026  
**Estudiante:** Carlos Guzmán  

## Objetivo

Desarrollar una aplicación web para la gestión de libros electrónicos, usuarios y préstamos, aplicando conceptos de programación orientada a objetos, concurrencia, servicios web y pruebas de software.

## Funcionalidades Principales

- Gestión de usuarios.
- Gestión de libros electrónicos.
- Registro de préstamos.
- Control de disponibilidad de libros.
- Reporte general.
- Servicios web mediante API.
- Serialización de datos mediante JSON.
- Manejo concurrente de solicitudes.
- Pruebas unitarias e integración.

## Arquitectura MVC

El proyecto está organizado en:

- `modelos`: estructuras principales del sistema.
- `vistas`: interfaz web HTML.
- `controladores`: gestión de solicitudes HTTP y APIs.
- `servicios`: lógica de negocio.
- `repositorios`: almacenamiento temporal en memoria.
- `interfaces`: contratos del sistema.
- `rutas`: definición de rutas web.
- `pruebas`: pruebas del sistema.

## Servicios Web Implementados

| Método | Ruta | Funcionalidad |
|---|---|---|
| GET | `/api/usuarios` | Lista usuarios |
| POST | `/api/usuarios` | Crea usuario |
| GET | `/api/usuarios/{id}` | Busca usuario |
| GET | `/api/libros` | Lista libros |
| POST | `/api/libros` | Crea libro |
| GET | `/api/libros/{id}` | Busca libro |
| GET | `/api/prestamos` | Lista préstamos |
| POST | `/api/prestamos` | Crea préstamo |
| GET | `/api/reportes` | Muestra reporte |
| GET | `/api/concurrencia` | Demostración de concurrencia |

## Concurrencia

El servidor web de Go atiende múltiples solicitudes mediante goroutines. Además, los repositorios en memoria utilizan `sync.RWMutex` para proteger el acceso concurrente a los datos.

## Ejecución

```bash
go run .
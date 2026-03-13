# API de Gestión de Videojuegos - Ejercicio 4

Esta es una API RESTful construida con la **librería estándar de Go** (sin frameworks externos). Permite gestionar una colección de videojuegos con persistencia de datos en un archivo JSON local.

## Información del Estudiante
- **Nombre:** Alejandra Saraí Avilés González
- **Carnet:** 24722
- **Puerto Asignado:** 24722

## Tema Elegido
Gestión de una biblioteca de **Videojuegos**. Cada elemento cuenta con: `id`, `title`, `genre`, `platform`, `release_year` y `rating`.

## Endpoints Implementados

| Método | Endpoint | Parámetros | Descripción |
|--------|----------|------------|-------------|
| **GET** | `/api/items` | Ninguno | Lista todos los videojuegos. |
| **GET** | `/api/items?id=1` | Query Param | Filtra un videojuego por su ID. |
| **GET** | `/api/items?genre=RPG` | Query Param | Filtra videojuegos por género (Multi-filtro). |
| **GET** | `/api/items/1` | Path Param | Obtiene un videojuego específico por su ID en la ruta. |
| **POST** | `/api/items` | Body JSON | Crea un nuevo videojuego y lo guarda en el archivo JSON. |
| **PUT** | `/api/items/1` | Path Param + Body | Reemplaza un videojuego existente por uno nuevo. |
| **PATCH** | `/api/items/1` | Path Param + Body | Actualiza parcialmente un videojuego (ej. solo el rating). |
| **DELETE** | `/api/items/1` | Path Param | Elimina permanentemente un videojuego de la lista. |

## Características Técnicas (Criterios de Evaluación)
- **Persistencia Real:** Cualquier cambio (POST, PUT, PATCH, DELETE) se guarda automáticamente en el archivo `data/games.json`.
- **Validación Robusta:** El sistema verifica que los campos obligatorios estén presentes y maneja errores de formato JSON.
- **Manejo de Errores:** Respuestas de error consistentes con formato `{"status": 4xx, "message": "..."}`.
- **Docker:** Archivo `Dockerfile` configurado para desplegar la aplicación en un contenedor aislado.

## Instrucciones para Ejecutar

### 1. Ejecución Local (Go)
- Asegúrate de estar en la raíz del proyecto y ejecuta:
        '''bash
        go run main.go
- El servidor iniciará en: http://localhost:24722

### 2. Ejecución con Docker
- Para construir la imagen y correr el contenedor:
        '''bash
        docker build -t api-videojuegos .
        docker run -p 24722:24722 api-videojuegos

## Ejemplos de JSON para Pruebas

### Crear / Actualizar (POST / PUT)
'''json
{
  "id": 11,
  "title": "Starfield",
  "genre": "RPG",
  "platform": "Xbox/PC",
  "release_year": 2023,
  "rating": 8.5
}

### Actualización Parcial (PATCH)
'''json
{
  "rating": 9.8
}

## Estructura del Repositorio
- /data/games.json: Base de datos en formato JSON (Persistencia).
- main.go: Servidor, rutas y lógica de negocio.
- Dockerfile: Configuración de contenedor.
- README.md: Documentación del proyecto.
# 10 - POST JSON

En esta etapa la API incorpora soporte para crear nuevos recursos utilizando el método `POST` con un body en formato JSON.

Se mantiene el modelo basado en archivo como fuente inicial de datos, pero ahora el servidor permite modificar el estado en memoria.

---

## 🎯 Objetivo de esta etapa

Comprender:

* Cómo leer el body de una petición HTTP
* Cómo usar `json.NewDecoder` para decodificar JSON
* Cómo validar datos enviados por el cliente
* Cómo generar identificadores dinámicamente
* Cómo devolver `201 Created`

---

## 📁 Estructura del proyecto

```
.
├── main.go
├── data/
│   └── teams.json
├── Dockerfile
└── docker-compose.yml
```

La estructura no cambia respecto a la rama anterior.

---

## 🧠 Qué cambió respecto a la rama anterior

Antes:

* Solo existían endpoints `GET`
* Los datos eran únicamente de lectura

Ahora:

* `POST /api/teams` permite crear nuevos equipos
* El servidor decodifica JSON desde el body
* Se valida el input antes de procesarlo
* Se devuelve `201 Created` cuando el recurso es creado correctamente

---

## 🧩 Ejemplo de uso

Crear un nuevo equipo:

```
POST /api/teams
```

Body (JSON):

```json
{
  "name": "Test FC"
}
```

Respuesta esperada:

```json
{
  "id": 21,
  "name": "Test FC"
}
```

Status:

```
201 Created
```

---

## 🔎 Conceptos introducidos

* `json.NewDecoder(r.Body).Decode(&struct)`
* Validación básica de campos requeridos
* Generación automática de ID
* Mutación de datos en memoria (`append`)
* Uso correcto de códigos de estado HTTP

---

## 🐳 Ejecución

El servidor escucha en el puerto 80 dentro del contenedor.

En `docker-compose.yml` se mapea:

```yaml
ports:
  - "8080:80"
```

Probar con Postman o curl:

```bash
curl -X POST http://localhost:8080/api/teams \
  -H "Content-Type: application/json" \
  -d '{"name":"Test FC"}'
```

---

## 📌 Qué estamos aprendiendo realmente

En esta etapa entendemos que:

* Las APIs no solo leen datos, también los crean
* El backend debe validar y procesar el body de las peticiones
* HTTP define semántica clara para creación de recursos (`POST` + `201`)

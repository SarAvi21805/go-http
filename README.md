# Backend con Go — De TCP a APIs RESTful

Este repositorio es un recorrido progresivo para entender cómo funciona un servidor backend en Go desde la base.

El objetivo no es aprender un framework.

El objetivo es entender el problema antes de usar la solución.

---

## 🧠 Enfoque

Comenzamos desde el nivel más bajo posible:

* TCP puro
* Construcción manual de HTTP
* Routing manual
* Uso de la librería estándar
* Separación de archivos
* Servir recursos estáticos
* Generación de vistas
* Construcción de APIs JSON
* Persistencia en archivo
* Manejo de parámetros
* Creación de recursos

Cada rama representa una capa adicional de abstracción.

La idea es poder moverse entre ramas y observar cómo evoluciona el servidor.

---

## 🎯 Qué se busca lograr

Que el estudiante entienda:

* Qué es realmente HTTP
* Qué ocurre cuando el navegador hace una petición
* Qué abstrae `net/http`
* Cómo funciona el routing
* Cómo se sirven archivos
* Cómo se renderizan vistas
* Cómo se construye una API JSON
* Cómo se modelan recursos y operaciones

---

## 🐳 Entorno

Todos los ejemplos están preparados para ejecutarse con Docker y Docker Compose.

Cada rama contiene sus propias instrucciones para levantar el proyecto.

---

## 📚 Ramas del repositorio

**[01-raw-tcp](https://github.com/menene/go-http/tree/01-raw-tcp)**  
Servidor construido directamente sobre TCP. Se construye manualmente la respuesta HTTP para entender cómo funciona el protocolo desde la base.

**[02-http-manual-routing](https://github.com/menene/go-http/tree/02-http-manual-routing)**  
Se parsea manualmente la primera línea del request para extraer método y ruta, implementando routing básico y códigos de estado.

**[03-net-http-basics](https://github.com/menene/go-http/tree/03-net-http-basics)**  
Se introduce la librería estándar `net/http`, eliminando el manejo manual del protocolo y mostrando el valor de la abstracción.

**[04-serve-html-files](https://github.com/menene/go-http/tree/04-serve-html-files)**  
El servidor comienza a servir archivos HTML reales junto con recursos estáticos como CSS e imágenes.

**[05-templates](https://github.com/menene/go-http/tree/05-templates)**  
Se introduce `html/template`, permitiendo generar vistas desde el servidor y reutilizar un layout común.

**[06-posts](https://github.com/menene/go-http/tree/06-posts)**  
Se incorporan formularios HTML y el método POST, permitiendo que el servidor reciba y procese datos enviados por el cliente.

**[07-json-api](https://github.com/menene/go-http/tree/07-json-api)**  
Se elimina la capa de vistas y el servidor pasa a ser una API pura que devuelve JSON utilizando `encoding/json`.

**[08-file-db](https://github.com/menene/go-http/tree/08-file-db)**  
La API comienza a leer datos desde un archivo JSON, simulando una base de datos basada en archivo.

**[09-query-params](https://github.com/menene/go-http/tree/09-query-params)**  
Se agregan parámetros en la URL (`?id=`), permitiendo filtrar resultados y modificar el comportamiento del endpoint según el input recibido.

**[10-post-json](https://github.com/menene/go-http/tree/10-post-json)**  
Se incorpora soporte para `POST` con body en formato JSON, permitiendo crear nuevos recursos, validar datos y devolver `201 Created`.

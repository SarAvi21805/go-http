package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Modelo de datos
type Game struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Genre       string  `json:"genre"`
	Platform    string  `json:"platform"`
	ReleaseYear int     `json:"release_year"`
	Rating      float64 `json:"rating"`
}

// Respuesta de error
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var (
	games     []Game
	mutex     sync.Mutex
	dataFile  = "data/games.json"
	serverPort = "24722"
)

func main() {
	loadData()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/items", handleItems)
	mux.HandleFunc("/api/items/", handleItemByID) // Para Path Parameters

	fmt.Printf("Servidor corriendo en el puerto %s...\n", serverPort)
	log.Fatal(http.ListenAndServe(":"+serverPort, mux))
}

// Persistencia de datos

func loadData() {
	file, _ := os.ReadFile(dataFile)
	json.Unmarshal(file, &games)
}

func saveData() {
	data, _ := json.MarshalIndent(games, "", "  ")
	os.WriteFile(dataFile, data, 0644)
}

// Handlers

func handleItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		getGames(w, r)
	case http.MethodPost:
		postGame(w, r)
	default:
		sendError(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func handleItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	// Extraer ID del Path /api/items/{id}
	segments := strings.Split(r.URL.Path, "/")
	if len(segments) < 4 {
		sendError(w, "ID no proporcionado", http.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(segments[3])

	switch r.Method {
	case http.MethodGet:
		getGameByID(w, id)
	case http.MethodPut:
		putGame(w, r, id)
	case http.MethodDelete:
		deleteGame(w, id)
	default:
		sendError(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// Lógica para el CRUD

func getGames(w http.ResponseWriter, r *http.Request) {
	queryID := r.URL.Query().Get("id")
	genre := r.URL.Query().Get("genre")

	// Filtro por Query Parameter ?id=X
	if queryID != "" {
		id, _ := strconv.Atoi(queryID)
		getGameByID(w, id)
		return
	}

	// Filtro por Género (Multi-filtro)
	filtered := games
	if genre != "" {
		filtered = []Game{}
		for _, g := range games {
			if strings.EqualFold(g.Genre, genre) {
				filtered = append(filtered, g)
			}
		}
	}

	json.NewEncoder(w).Encode(filtered)
}

func getGameByID(w http.ResponseWriter, id int) {
	for _, g := range games {
		if g.ID == id {
			json.NewEncoder(w).Encode(g)
			return
		}
	}
	sendError(w, "Videojuego no encontrado", http.StatusNotFound)
}

func postGame(w http.ResponseWriter, r *http.Request) {
	var newGame Game
	if err := json.NewDecoder(r.Body).Decode(&newGame); err != nil {
		sendError(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validación
	if newGame.ID == 0 || newGame.Title == "" || newGame.Genre == "" {
		sendError(w, "Campos id, title y genre son obligatorios", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	games = append(games, newGame)
	saveData()
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGame)
}

func putGame(w http.ResponseWriter, r *http.Request, id int) {
	var updatedGame Game
	json.NewDecoder(r.Body).Decode(&updatedGame)

	mutex.Lock()
	defer mutex.Unlock()

	for i, g := range games {
		if g.ID == id {
			updatedGame.ID = id // Conservando ID original
			games[i] = updatedGame
			saveData()
			json.NewEncoder(w).Encode(updatedGame)
			return
		}
	}
	sendError(w, "No se pudo actualizar: no existe", http.StatusNotFound)
}

func deleteGame(w http.ResponseWriter, id int) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, g := range games {
		if g.ID == id {
			games = append(games[:i], games[i+1:]...)
			saveData()
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	sendError(w, "No se pudo eliminar: no existe", http.StatusNotFound)
}

// Helpers

func sendError(w http.ResponseWriter, message string, code int) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Status: code, Message: message})
}
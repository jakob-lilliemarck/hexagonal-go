package handlers

import (
	"encoding/json"
	"hexagonal-go/src/lib/driving"
	"net/http"

	"github.com/gorilla/mux"
)

type ContextKey string

const (
	FishServiceKey ContextKey = "FishService"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// GetFishByID handles GET /fish/:id
func GetFishByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	service := r.Context().Value(FishServiceKey).(driving.FishService)
	fish, err := service.Read(id)
	if err != nil {
		http.Error(w, "Fish not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, fish)
}

// GetFishCollection handles GET /fish
func GetFishCollection(w http.ResponseWriter, r *http.Request) {
	service := r.Context().Value(FishServiceKey).(driving.FishService)
	fish := service.ReadCollection()
	respondWithJSON(w, http.StatusOK, fish)
}

// CreateFish handles POST /fish to create a new fish
func CreateFish(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Species string `json:"species"`
		Age     uint32 `json:"age"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	service := r.Context().Value(FishServiceKey).(driving.FishService)
	fish := service.Create(input.Species, input.Age)

	// Respond with the created fish as JSON
	respondWithJSON(w, http.StatusCreated, fish)
}

// CreateFish handles POST /fish to create a new fish
func UpdateFish(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	var input struct {
		Age uint32 `json:"age"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	service := r.Context().Value(FishServiceKey).(driving.FishService)
	fish, err := service.Update(id, input.Age)
	if err != nil {
		http.Error(w, "ERROR", http.StatusBadRequest)
	}

	// Respond with the created fish as JSON
	respondWithJSON(w, http.StatusCreated, fish)
}

// CreateFish handles POST /fish to create a new fish
func DeleteFish(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, ok := params["id"]
	if !ok {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	service := r.Context().Value(FishServiceKey).(driving.FishService)

	service.Delete(id)

	respondWithJSON(w, http.StatusCreated, nil)
}

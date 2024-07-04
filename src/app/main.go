//go:build ignore

package main

import (
	"context"
	"hexagonal-go/src/lib/core"
	"hexagonal-go/src/lib/driven"
	"hexagonal-go/src/lib/driving"
	"hexagonal-go/src/lib/handlers"
	"hexagonal-go/src/lib/utils"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func FishServiceMiddleware(service core.Driving) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, handlers.FishServiceKey, service)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func main() {
	hash := utils.Hash[string, *core.Fish]()
	driven := driven.DrivenAdapter(hash)
	service := driving.FishService{
		Repository: driven,
	}

	r := mux.NewRouter()
	r.Use(FishServiceMiddleware(service))

	r.HandleFunc("/fish/{id}", handlers.GetFishByID).Methods(http.MethodGet)
	r.HandleFunc("/fish", handlers.GetFishCollection).Methods(http.MethodGet)
	r.HandleFunc("/fish", handlers.CreateFish).Methods(http.MethodPost)
	r.HandleFunc("/fish/{id}", handlers.UpdateFish).Methods(http.MethodPut)
	r.HandleFunc("/fish/{id}", handlers.DeleteFish).Methods(http.MethodDelete)

	http.Handle("/", r)

	log.Println("Starting server on http://localhost:9292")
	log.Fatal(http.ListenAndServe(":9292", nil))
}

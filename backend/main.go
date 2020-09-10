package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/maet98/sellerApp/models/buyer"
)

func main() {
	port := "3000"
	r := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // you can add routes here www.example.com
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})

	r.Use(cors.Handler)

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/buyers", func(w http.ResponseWriter, r *http.Request) {
		buyers, err := buyer.FindAll()
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(buyers)
	})

	r.Get("/buyers/{buyerId}", func(w http.ResponseWriter, r *http.Request) {
		BuyerID := chi.URLParam(r, "buyerId")
		buyers, err := buyer.GetById(BuyerID)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(buyers)
	})

	fmt.Println("Serving on port: " + port)
	http.ListenAndServe(":"+port, r)
}

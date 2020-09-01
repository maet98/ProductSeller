package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/maet98/sellerApp/newsfeed"
)

func main() {
	port := "3000"
	feed := newsfeed.New()
	r := chi.NewRouter()
	feed.Add(newsfeed.Item{
		Title: "Hello",
		Post:  "World",
	})

	r.Use(middleware.New

	// Get newsfeed
	r.Get("/newsfeed", func(w http.ResponseWriter, r *http.Request) {
		items := feed.GetAll()
		json.NewEncoder(w).Encode(items)
	})

	// Post newsfeed
	r.Post("/newsfeed", func(w http.ResponseWriter, r *http.Request) {
		request := map[string]string{}
		json.NewDecoder(r.Body).Decode(&request)

		feed.Add(newsfeed.Item{
			Title: request["title"],
			Post:  request["post"],
		})

		w.Write([]byte("Good Job"))
	})

	fmt.Println("Serving on port: " + port)
	http.ListenAndServe(":"+port, r)
}

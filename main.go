package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	config := NewConfig()
	if config == nil {
		log.Fatal("Error creating config")
		return
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(config.filePathRoot)))

	mux.HandleFunc("POST /api/groceries/items", config.handlerAddItem)
	mux.HandleFunc("GET /api/groceries/items", config.handlerGetAllItems)
	mux.HandleFunc("GET /api/groceries/items/{listID}", config.handlerGetAllItemsForList)
	mux.HandleFunc("DELETE /api/groceries/items/{itemID}", config.handlerDeleteItemByID)

	mux.HandleFunc("GET /api/groceries/lists", config.handlerGetAllLists)
	mux.HandleFunc("POST /api/groceries/lists", config.handlerCreateNewList)
	mux.HandleFunc("DELETE /api/groceries/lists/{listID}", config.handlerDeleteListByID)

	wrappedMux := loggingMiddlewareHandler(mux)

	srv := &http.Server{
		Addr:    ":" + config.port,
		Handler: wrappedMux,
	}

	log.Printf("Serving on: http://localhost:%s", config.port)
	log.Fatal(srv.ListenAndServe())
}

func loggingMiddlewareHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s\t%s\tPayload Size: %v", r.Method, r.URL, r.RemoteAddr, r.ContentLength)
		next.ServeHTTP(w, r)
	})
}

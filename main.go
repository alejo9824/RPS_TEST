package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {

	// Crear enrrutador
	router := http.NewServeMux()
	//manejador para archivos estaticos
	fs := http.FileServer(http.Dir("./static"))

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	// Definir rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/NewGame", handlers.NewGame)
	router.HandleFunc("/Game", handlers.Game)
	router.HandleFunc("/Play", handlers.Play)
	router.HandleFunc("/About", handlers.About)

	port := ":8080"
	log.Printf("Servidor escuchando en  http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, router))

}

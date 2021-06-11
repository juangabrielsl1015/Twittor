package main

import (
	"log"

	"github.com/juangabrielsl1015/twittor.git/bd"
	"github.com/juangabrielsl1015/twittor.git/handlers"
)

func main() {
	log.Println("Hola mundo")

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos")
		return
	}

	handlers.Manejadores()

	log.Print("Hay conexión a la base de datos")
}

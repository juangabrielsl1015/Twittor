package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/juangabrielsl1015/twittor.git/middlew"
	"github.com/juangabrielsl1015/twittor.git/routes"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routes.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarperfil", middlew.ChequeoBD(middlew.ValidoJWT(routes.ModificarPerfil))).Methods("PUT")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

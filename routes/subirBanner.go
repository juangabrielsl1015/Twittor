package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/juangabrielsl1015/twittor.git/bd"
	"github.com/juangabrielsl1015/twittor.git/models"
)

/* SubirAvatar sube el avatar al servidor */
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	if err != nil {
		http.Error(w, "Error al obtener la imagen. "+err.Error(), http.StatusBadRequest)
		return
	}

	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen. "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen. "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el banner en la BD. "+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

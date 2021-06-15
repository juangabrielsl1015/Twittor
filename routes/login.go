package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/juangabrielsl1015/twittor.git/bd"
	"github.com/juangabrielsl1015/twittor.git/jwt"
	"github.com/juangabrielsl1015/twittor.git/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o contraseña inválido. "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido. "+err.Error(), http.StatusBadRequest)
		return
	}

	usuario, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		http.Error(w, "El email del usuario es requerido. "+err.Error(), http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GeneroJWT(usuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el token correspondiente. "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Creación de cookie en el navegador
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}

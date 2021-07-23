package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/juangabrielsl1015/twittor.git/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ConsultoRelacion consulta la relación entre dos usuario */
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	bd := MongoCN.Database("twittor")
	col := bd.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	return true, nil
}

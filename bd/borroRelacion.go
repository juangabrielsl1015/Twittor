package bd

import (
	"context"
	"time"

	"github.com/juangabrielsl1015/twittor.git/models"
)

/* BorroRelacion borra la relación en la BD */
func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	bd := MongoCN.Database("twittor")
	col := bd.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}

package models

/* Tweet captura el body, el mensaje que nos llega*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}

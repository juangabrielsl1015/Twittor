package models

import "time"

/* GraboTweet es el formato o estructura que tendrá nuestro Tweet */
type GraboTweet struct {
	Userid  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}

package models

/* RespuesteConsultaRelacion tiene el true o false que se obtiene de consultar la relación entre 2 usuarios */
type RespuesteConsultaRelacion struct {
	Status bool `json:"status"`
}

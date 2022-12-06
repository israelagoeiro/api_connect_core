package mysql

import "time"

type TempoModel struct {
	LigadoParcial   time.Time `json:"ligadoParcial" bson:"ligadoParcial"`
	LigadoDecorrido time.Time `json:"ligadoDecorrido" bson:"ligadoDecorrido"`
}

type FdibModel struct {
	Id           int        `json:"id" bson:"id"`
	IdPeca       int        `json:"idPeca" bson:"idPeca"`
	Nserlum      int        `json:"nserlum" bson:"nserlum"`
	ColetaRede   time.Time  `json:"coletaRede" bson:"coletaRede"`
	Etiqueta     string     `json:"etiqueta" bson:"etiqueta"`
	Status       bool       `json:"status" bson:"status"`
	Tempo        TempoModel `json:"tempo" bson:"tempo"`
	Total        int        `json:"total" bson:"total"`
	TotalNserlum int        `json:"totalNserlum" bson:"totalNserlum"`
	TotalStatus  int        `json:"totalStatus" bson:"totalStatus"`
}

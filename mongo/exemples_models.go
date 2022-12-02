package mongo

import "time"

type TempoModel struct {
	LigadoParcial   time.Time `json:"ligadoParcial" bson:"ligadoParcial"`
	LigadoDecorrido time.Time `json:"ligadoDecorrido" bson:"ligadoDecorrido"`
}

type FdibModel struct {
	IdPeca     string     `json:"idPeca" bson:"idPeca"`
	Nserlum    int        `json:"nserlum" bson:"nserlum"`
	ColetaRede time.Time  `json:"coletaRede" bson:"coletaRede"`
	Etiqueta   string     `json:"etiqueta" bson:"etiqueta"`
	Status     bool       `json:"status" bson:"status"`
	Tempo      TempoModel `json:"tempo" bson:"tempo"`
}

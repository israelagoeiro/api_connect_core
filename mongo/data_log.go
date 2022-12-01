package mongo

import (
	"time"
)

type HistoryModel struct {
	Checksum  string    `json:"cks" bson:"cks"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	Action    string    `json:"action" bson:"action"`
	Owner     string    `json:"owner" bson:"owner"`
}

type InfoModel struct {
	ChangedAt time.Time       `json:"changedAt" bson:"changedAt"`
	CreatedAt time.Time       `json:"createdAt" bson:"createdAt"`
	Checksum  string          `json:"cks" bson:"cks"`
	Owner     string          `json:"owner" bson:"owner"`
	Version   int             `json:"v" bson:"v"`
	History   *[]HistoryModel `json:"history,omitempty" bson:"history,omitempty"`
}

type DataLog struct {
	Action       string
	SaveChange   bool
	SaveHistory  bool
	SaveInfo     bool
	SaveAnalytic bool
	Info         InfoModel
}

type MongoDataLog struct {
	PrepareInsert func(input InsertInput)
	PrepareUpdate func(input UpdateInput)
}

func NewMongoDataLog(dataLog DataLog) MongoDataLog {
	var _info InfoModel
	return MongoDataLog{
		PrepareInsert: func(input InsertInput) {
			//TODO:: Garantir que somente quando houver usuários possa registrar info e history
			if dataLog.SaveInfo || dataLog.SaveHistory {
				//TODO:: Implementar no MongoDataLog:SaveInfo Checksum
				//TODO:: Implementar no MongoDataLog:SaveInfo Owner
				_info = InfoModel{
					ChangedAt: time.Now(),
					CreatedAt: time.Now(),
					Checksum:  "d24a3ee7e829b770ea0f6441bd4294e1f29215dd",
					Owner:     "5f427ecbaa54ea68b0a20b5a:account-public",
					Version:   0,
				}
			}
			if dataLog.SaveHistory {
				if len(dataLog.Action) > 0 {
					//TODO:: Implementar no MongoDataLog:SaveHistory Checksum
					//TODO:: Implementar no MongoDataLog:SaveHistory Owner
					_info.History = &[]HistoryModel{
						{
							Checksum:  "d24a3ee7e829b770ea0f6441bd4294e1f29215dd",
							CreatedAt: time.Now(),
							Action:    dataLog.Action,
							Owner:     "5f427ecbaa54ea68b0a20b5a:account-public",
						},
					}
				} else {
					panic("Error DataLog: A propriedade 'Action' não pode ser vazia.")
				}
			}
		},
		PrepareUpdate: func(input UpdateInput) {
			if dataLog.SaveInfo || dataLog.SaveHistory {
				//TODO:: Implementar no MongoDataLog:SaveInfo Checksum
				//TODO:: Implementar no MongoDataLog:SaveInfo UserId
				input.Set("_info.changedAt", time.Now())
				input.Set("_info.cks", "d24a3ee7e829b770ea0f6441bd4294e1f29215dx")
				input.Set("_info.owner", "5f427ecbaa54ea68b0a20b5a:account-public")
				input.Inc("_info.v", 1)
			}
			if dataLog.SaveHistory {
				if len(dataLog.Action) > 0 {
					//TODO:: Implementar no MongoDataLog:SaveHistory Checksum
					//TODO:: Implementar no MongoDataLog:SaveHistory UserId
					input.AddToSet("_info.history", HistoryModel{
						Checksum:  "d24a3ee7e829b770ea0f6441bd4294e1f29215dd",
						CreatedAt: time.Now(),
						Action:    dataLog.Action,
						Owner:     "5f427ecbaa54ea68b0a20b5a:account-public",
					})
				} else {
					panic("Error DataLog: A propriedade 'Action' não pode ser vazia.")
				}
			}
		},
	}
}

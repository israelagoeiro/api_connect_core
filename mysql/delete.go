package mysql

import "fmt"

//DeleteInterface reune comandos de exclusão
/*type DeleteInterface interface {
	deleteOne() DataResult
	deleteMany() DataResult
}*/

func DeleteOne(param DeleteParams) DataResult {
	param.Filter.Limit(1)
	document := NewDeleteDocument()
	return document.DeleteOne(param)
}
func DeleteMany(param DeleteParams) DataResult {
	document := NewDeleteDocument()
	return document.DeleteMany(param)
}

/*func (param DeleteParams) deleteMany() DataResult {
	document := NewDeleteDocument()
	return document.DeleteMany(param)
}

func (param DeleteParams) deleteOne() DataResult {
	param.Filter.Limit(1)
	document := NewDeleteDocument()
	return document.DeleteOne(param)
}*/

type DeleteParams struct {
	Collection string
	FindParams FindParams
	Filter     MysqlFilter
}

func (u DeleteParams) queryDelete() string {
	dbQuery := fmt.Sprintf("DELETE FROM %s", u.Collection)
	if u.Filter.Values != nil {
		filter := u.Filter.Values()
		if filter != "" {
			dbQuery += filter
		}
	}
	return dbQuery + ";"
}

func (u DeleteParams) PrintQuery() {
	fmt.Println("DeletePrintQuery:", u.queryDelete())
}

type DeleteDocument struct {
	DeleteMany func(param DeleteParams) DataResult
	DeleteOne  func(param DeleteParams) DataResult
}

func NewDeleteDocument() DeleteDocument {
	deleteDocument := DeleteDocument{
		DeleteMany: func(param DeleteParams) (dataResult DataResult) {
			dbQuery := param.queryDelete()
			result, err := Client.Exec(dbQuery)
			if err != nil {
				fmt.Println("Error DeleteDocument:DeleteOne:Query", err.Error())
			}
			rowsAffected, _ := result.RowsAffected()
			dataResult.rowsAffected = int(rowsAffected)
			return dataResult
		},
		DeleteOne: func(param DeleteParams) (dataResult DataResult) {
			//TODO:: Implementar find

			dbQuery := param.queryDelete()
			result, err := Client.Exec(dbQuery)
			if err != nil {
				fmt.Println("Error DeleteDocument:DeleteOne:Query", err.Error())
			}
			rowsAffected, _ := result.RowsAffected()
			dataResult.rowsAffected = int(rowsAffected)
			return dataResult
		},
	}
	return deleteDocument
}

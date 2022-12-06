package mysql

import "fmt"

type UpdateParams struct {
	Collection string
	Filter     MysqlFilter
	FindParams FindParams
	Input      UpdateInput
}

func (u UpdateParams) queryUpdate() string {
	dbQuery := fmt.Sprintf("UPDATE %s %s", u.Collection, u.Input.Values())
	if u.Filter.Values != nil {
		filter := u.Filter.Values()
		if filter != "" {
			dbQuery += filter
		}
	}
	return dbQuery + ";"
}

func (u UpdateParams) PrintQuery() {
	fmt.Println("UpdatePrintQuery:", u.queryUpdate())
}

type UpdateDocument struct {
	UpdateOne  func(param UpdateParams) DataResult
	UpdateMany func(param UpdateParams) DataResult
}

func NewUpdateDocument() UpdateDocument {
	updateDocument := UpdateDocument{
		UpdateMany: func(param UpdateParams) (result DataResult) {
			dbQuery := param.queryUpdate()
			dataResult, err := Client.Exec(dbQuery)
			if err != nil {
				fmt.Println("Error UpdateDocument:UpdateMany:Query", err.Error())
			}
			rowsAffected, _ := dataResult.RowsAffected()
			result.rowsAffected = int(rowsAffected)
			return result
		},
		UpdateOne: func(param UpdateParams) (result DataResult) {
			//TODO:: Implementar find
			dbQuery := param.queryUpdate()
			dataResult, err := Client.Exec(dbQuery)
			if err != nil {
				fmt.Println("Error UpdateDocument:UpdateOne:Query", err.Error())
			}
			rowsAffected, _ := dataResult.RowsAffected()
			result.rowsAffected = int(rowsAffected)
			return result
		},
	}

	return updateDocument
}

package mysql

import (
	"fmt"
)

type InsertParams struct {
	Collection string
	Input      InsertInput
	FindParams FindParams
}

func (u InsertParams) queryInsert() string {
	fields, values := u.Input.Values()
	dbQuery := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", u.Collection, fields, values)
	return dbQuery
}

func (u InsertParams) PrintQuery() {
	fmt.Println("InsertPrintQuery:", u.queryInsert())
}

type DocumentInsert struct {
	InsertMany  func(param InsertParams) DataResult
	InsertOne   func(param InsertParams) DataResult
	InsertQuery func(query string) DataResult
}

func NewDocumentInsert() DocumentInsert {
	updateDocument := DocumentInsert{
		InsertMany: func(param InsertParams) (dataResult DataResult) {
			field, values := param.Input.Values()
			dbQuery := fmt.Sprintf("INSERT INTO %s%s VALUES%s;", param.Collection, field, values)
			result, err := Client.Exec(dbQuery)

			if err != nil {
				fmt.Println("Error DocumentInsert:InsertOne:Query", err.Error())
			}

			lastInsertId, _ := result.LastInsertId()
			rowsAffected, _ := result.RowsAffected()

			dataResult.lastInsertId = int(lastInsertId)
			dataResult.rowsAffected = int(rowsAffected)

			return dataResult
		},
		InsertOne: func(param InsertParams) (dataResult DataResult) {
			dbQuery := param.queryInsert()
			result, err := Client.Exec(dbQuery)
			if err != nil {
				fmt.Println("Error DocumentInsert:InsertOne:Query", err.Error())
			}

			lastInsertId, _ := result.LastInsertId()
			rowsAffected, _ := result.RowsAffected()

			dataResult.lastInsertId = int(lastInsertId)
			dataResult.rowsAffected = int(rowsAffected)

			return dataResult
		},
		InsertQuery: func(query string) (dataResult DataResult) {
			result, err := Client.Exec(query)
			if err != nil {
				fmt.Println("Error DocumentInsert:InsertOne:Query", err.Error())
			}

			lastInsertId, _ := result.LastInsertId()
			rowsAffected, _ := result.RowsAffected()

			dataResult.lastInsertId = int(lastInsertId)
			dataResult.rowsAffected = int(rowsAffected)

			return dataResult
		},
	}

	return updateDocument
}

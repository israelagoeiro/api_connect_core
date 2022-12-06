package mysql

import (
	"fmt"
)

type FindParams struct {
	Collection string
	Filter     MysqlFilter
	Fields     Fields
}

func (u FindParams) querySelect() string {
	fields := u.Fields.Values(u.Collection)
	dbQuery := fmt.Sprintf("SELECT %s FROM %s", fields, u.Collection)
	if u.Filter.Values != nil {
		filter := u.Filter.Values()
		if filter != "" {
			dbQuery += filter
		}
	}
	return dbQuery + ";"
}

func (u FindParams) PrintQuery() {
	fmt.Println("FindPrintQuery:", u.querySelect())
}

type FindDocument struct {
	Find      func(param FindParams) FindDataResult
	FindQuery func(query string) FindDataResult
	FindOne   func(param FindParams) FindDataResult
}

func NewFindDocument() FindDocument {
	updateDocument := FindDocument{
		Find: func(param FindParams) FindDataResult {
			dbQuery := param.querySelect()
			rows, err := Client.Query(dbQuery)
			if err != nil {
				fmt.Println("Error Find.Query", err.Error())
			}
			return FindDataResult{
				rows: rows,
			}
		},
		FindQuery: func(query string) FindDataResult {
			rows, err := Client.Query(query)
			if err != nil {
				fmt.Println("Error Find.Query", err.Error())
			}
			return FindDataResult{
				rows: rows,
			}
		},
		FindOne: func(param FindParams) FindDataResult {
			if param.Filter.Limit != nil {
				param.Filter.Limit(1)
			}
			dbQuery := param.querySelect()
			rows, err := Client.Query(dbQuery)
			if err != nil {
				fmt.Println("Error FindOne.Query", err.Error())
			}
			return FindDataResult{
				rows: rows,
			}
		},
	}
	return updateDocument
}

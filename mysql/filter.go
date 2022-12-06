package mysql

import (
	"fmt"
)

type MysqlFilter struct {
	Where     func(condition string)
	And       func(condition string)
	Or        func(condition string)
	OrderBy   func(condition string)
	GroupBy   func(condition string)
	Having    func(condition string)
	Limit     func(values ...int)
	InnerJoin func(collection string, condition string)
	LeftJoin  func(collection string, condition string)
	RightJoin func(collection string, condition string)
	Values    func() string
}

func NewFilter() MysqlFilter {
	where := ""
	groupBy := ""
	having := ""
	orderBy := ""
	innerJoin := ""
	limitAndOffset := ""
	var listAnd []string
	var listOr []string

	return MysqlFilter{
		Where: func(condition string) {
			where = condition
		},
		And: func(condition string) {
			listAnd = append(listAnd, condition)
		},
		Or: func(condition string) {
			listOr = append(listOr, condition)
		},
		OrderBy: func(order string) {
			orderBy = order
		},
		GroupBy: func(fields string) {
			groupBy = fields
		},
		Having: func(fields string) {
			having = fields
		},
		InnerJoin: func(collection string, condition string) {
			innerJoin = fmt.Sprintf(" INNER JOIN %s ON %s", collection, condition)
		},
		LeftJoin: func(collection string, condition string) {
			innerJoin = fmt.Sprintf(" LEFT JOIN %s ON %s", collection, condition)
		},
		RightJoin: func(collection string, condition string) {
			innerJoin = fmt.Sprintf(" RIGHT JOIN %s ON %s", collection, condition)
		},
		Limit: func(values ...int) {
			if len(values) == 1 {
				limitAndOffset = fmt.Sprintf(" LIMIT %d", values[0])
			} else if len(values) == 2 {
				limitAndOffset = fmt.Sprintf(" LIMIT %d OFFSET %d", values[0], values[1])
			}
		},
		Values: func() string {
			conditions := ""
			if where != "" {
				conditions = fmt.Sprintf(" WHERE (%s", where)
				if conditions != "" && len(listAnd) > 0 {
					for _, and := range listAnd {
						conditions += fmt.Sprintf(" AND %s", and)
					}
				}
				conditions += ")"
			}

			if conditions != "" && len(listOr) > 0 {
				for _, or := range listOr {
					conditions += fmt.Sprintf(" OR (%s)", or)
				}
			}

			if groupBy != "" {
				conditions += fmt.Sprintf(" GROUP BY %s", groupBy)
			}

			if having != "" {
				conditions += fmt.Sprintf(" HAVING %s", having)
			}

			if innerJoin != "" {
				conditions += innerJoin
			}

			if orderBy != "" {
				conditions += fmt.Sprintf(" ORDER BY %s", orderBy)
			}

			if limitAndOffset != "" {
				conditions += limitAndOffset
			}

			return conditions
		},
	}
}

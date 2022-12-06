package mysql

import (
	"fmt"
	"sort"
	"strings"
)

type UpdateInput struct {
	Map     func(values map[string]any)
	ListMap func(values []map[string]any)
	Data    func(field string, values any)
	Values  func() string
}

func NewUpdateInput() UpdateInput {
	var listData []string
	return UpdateInput{
		Map: func(value map[string]any) {
			for k, v := range value {
				fieldValue := fmt.Sprintf("%s='%v'", k, v)
				listData = append(listData, fieldValue)
			}
		},
		Data: func(field string, value any) {
			fieldValue := fmt.Sprintf("%s='%v'", field, value)
			listData = append(listData, fieldValue)
		},
		Values: func() string {
			sort.Strings(listData)
			return "SET " + strings.Join(listData, ",")
		},
	}
}

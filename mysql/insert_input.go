package mysql

import (
	"fmt"
	"sort"
	"strings"
)

type InsertInput struct {
	Map     func(values map[string]any)
	ListMap func(values []map[string]any)
	Data    func(field string, values any)
	Values  func() (string, string)
}

func NewInsertInput() InsertInput {
	var listMap []map[string]any
	listData := map[string]any{}

	createFields := func(listValues []map[string]any) string {
		listFields := make([]string, 0, len(listValues[0]))
		var keys []string
		for k := range listValues[0] {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			listFields = append(listFields, k)
		}
		return "(" + strings.Join(listFields, ",") + ")"
	}

	createValues := func(listValues []map[string]any) string {
		result := []string{}
		for _, item := range listValues {
			listValues := make([]string, 0, len(item))
			var keys []string
			for k := range item {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			for _, k := range keys {
				listValues = append(listValues, fmt.Sprintf("%v", item[k]))
			}
			values := "('" + strings.Join(listValues, "','") + "')"
			result = append(result, values)
		}

		return strings.Join(result, ",")
	}

	return InsertInput{
		Map: func(value map[string]any) {
			values := map[string]any{}
			for k, v := range value {
				values[k] = fmt.Sprintf("%v", v)
			}
			listMap = append(listMap, values)
		},
		Data: func(field string, value any) {
			listData[field] = fmt.Sprintf("%v", value)
		},
		Values: func() (string, string) {
			var listValues []map[string]any
			if len(listMap) > 0 {
				listValues = listMap
			}

			if len(listData) > 0 {
				listValues = append(listValues, listData)
			}
			fields := createFields(listValues)
			values := createValues(listValues)

			return fields, values
		},
	}
}

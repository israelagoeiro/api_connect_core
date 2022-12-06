package mysql

import (
	"strings"
)

type Fields struct {
	All    func()
	Add    func(fields ...string)
	Values func(collection string) string
}

func NewFields() Fields {
	var listFields []string
	all := false

	return Fields{
		All: func() {
			all = true
		},
		Add: func(fields ...string) {
			listFields = fields
		},
		Values: func(collection string) string {
			if all {
				return "*"
			}

			var fields []string
			if len(listFields) > 0 {
				for _, item := range listFields {
					fields = append(fields, item)
				}
			}
			return strings.Join(fields, ",")
		},
	}
}

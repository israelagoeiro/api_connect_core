package mysql

//UpdateInterface reune comandos de exclusão
type UpdateInterface interface {
	mongoUpdateOne() DataResult
	mongoUpdateMany() DataResult
}

func (param UpdateParams) mongoUpdateOne() DataResult {
	document := NewUpdateDocument()
	return document.UpdateOne(param)
}
func UpdateOne(param UpdateInterface) DataResult {
	return param.mongoUpdateOne()
}

func (param UpdateParams) mongoUpdateMany() DataResult {
	document := NewUpdateDocument()
	return document.UpdateMany(param)
}
func UpdateMany(param UpdateInterface) DataResult {
	return param.mongoUpdateMany()
}

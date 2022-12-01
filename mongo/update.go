package mongo

//UpdateInterface reune comandos de exclusão
type UpdateInterface interface {
	mongoFindOneAndUpdate() DataResult
	mongoUpdateOne() DataResult
	mongoUpdateMany() DataResult
}

func (param UpdateParams) mongoFindOneAndUpdate() DataResult {
	document := NewUpdateDocument(param)
	return document.FindOneAndUpdate()
}

func FindOneAndUpdate(param UpdateInterface) DataResult {
	return param.mongoFindOneAndUpdate()
}

func (param UpdateParams) mongoUpdateOne() DataResult {
	document := NewUpdateDocument(param)
	return document.UpdateOne()
}
func UpdateOne(param UpdateInterface) DataResult {
	return param.mongoUpdateOne()
}

func (param UpdateParams) mongoUpdateMany() DataResult {
	document := NewUpdateDocument(param)
	return document.UpdateMany()
}
func UpdateMany(param UpdateInterface) DataResult {
	return param.mongoUpdateMany()
}

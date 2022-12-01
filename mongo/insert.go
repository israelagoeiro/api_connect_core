package mongo

//InsertInterface reune comandos de inserção
type InsertInterface interface {
	mongoInsertOne() DataResult
	mongoInsertMany() DataResult
}

func (param InsertParams) mongoInsertOne() DataResult {
	document := NewDocumentInsert(param)
	return document.InsertOne()
}

func InsertOne(param InsertInterface) DataResult {
	return param.mongoInsertOne()
}

func (param InsertParams) mongoInsertMany() DataResult {
	document := NewDocumentInsert(param)
	return document.InsertOne()
}

func InsertMany(param InsertInterface) DataResult {
	return param.mongoInsertMany()
}

package mysql

//InsertInterface reune comandos de inserção
type InsertInterface interface {
	mysqlInsertOne() DataResult
	mysqlInsertMany() DataResult
}

func (param InsertParams) mysqlInsertOne() DataResult {
	document := NewDocumentInsert()
	return document.InsertOne(param)
}

func InsertOne(param InsertInterface) DataResult {
	return param.mysqlInsertOne()
}

func (param InsertParams) mysqlInsertMany() DataResult {
	document := NewDocumentInsert()
	return document.InsertMany(param)
}

func InsertMany(param InsertInterface) DataResult {
	return param.mysqlInsertMany()
}

func InsertQuery(query string) DataResult {
	document := NewDocumentInsert()
	return document.InsertQuery(query)
}

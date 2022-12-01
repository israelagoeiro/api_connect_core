package mongo

//DeleteInterface reune comandos de exclusão
type DeleteInterface interface {
	deleteOne() DataResult
	deleteMany() DataResult
}

//InsertInterface reune comandos de inserção
type InsertInterface interface {
	_mongoInsertOne() DataResult
	_mongoInsertMany() DataResult
}

//UpdateInterface reune comandos de exclusão
type UpdateInterface interface {
	_mongoFindOneAndUpdate() DataResult
	_mongoUpdateOne() DataResult
	_mongoUpdateMany() DataResult
}

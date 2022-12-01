package mongo

//FindInterface reune comandos de consulta
type FindInterface interface {
	findOne() DataResult
	find() DataResult
}

//Find executa um comando find e retorna um 'DataResult' contendo um único documento da coleção.
//
//O parâmetro do 'filter' deve ser um 'MongoFilter' contendo operadores de consulta e pode ser usado para selecionar o
//documento a ser retornado. Não pode ser nulo. Se o 'filter' não corresponder a nenhum documento, será retornado um
//'DataResult' com um erro definido como ErrNoDocuments. Se 'filter' corresponder a vários documentos, o primeiro
//documento da lista um será selecionado do conjunto correspondente e retornado.
//
//O parâmetro 'options' deve ser um 'FindOptions' e pode ser usado para especificar opções para esta operação
//(consulte a documentação options.FindOneOptions).
//
//Para obter mais informações sobre o comando, consulte https://www.mongodb.com/docs/manual/reference/command/find/ .
func Find(param FindParams) DataResult {
	return param.find()
}
func (param FindParams) find() DataResult {
	document := NewDocumentFind(param)
	return document.Find()
}

//FindOne executa um comando find e retorna um 'DataResult' contendo um único documento da coleção.
//
//O parâmetro do 'filter' deve ser um 'MongoFilter' contendo operadores de consulta e pode ser usado para selecionar o
//documento a ser retornado. Não pode ser nulo. Se o 'filter' não corresponder a nenhum documento, será retornado um
//'DataResult' com um erro definido como ErrNoDocuments. Se 'filter' corresponder a vários documentos, o primeiro
//documento da lista um será selecionado do conjunto correspondente e retornado.
//
//O parâmetro 'options' deve ser um 'FindOptions' e pode ser usado para especificar opções para esta operação
//(consulte a documentação options.FindOneOptions).
//
//Para obter mais informações sobre o comando, consulte https://www.mongodb.com/docs/manual/reference/command/find/.
func FindOne(param FindParams) DataResult {
	return param.findOne()
}
func (param FindParams) findOne() DataResult {
	document := NewDocumentFind(param)
	return document.FindOne()
}

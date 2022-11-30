package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var MongoClient *mongo.Client

//DeleteInterface reune comandos de exclusão
type DeleteInterface interface {
	deleteOne() DataResult
}

//FindInterface reune comandos de consulta
type FindInterface interface {
	findOne() DataResult
	find() DataResult
}

//InsertInterface reune comandos de inserção
type InsertInterface interface {
	insertOne() DataResult
}

//UpdateInterface reune comandos de exclusão
type UpdateInterface interface {
	findOneAndUpdate() DataResult
}

func DeleteOne(param DeleteInterface) DataResult {
	return param.deleteOne()
	//return param.DeleteMany()
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
func Find(param MongoFindParams) DataResult {
	return param.find()
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
func FindOne(param MongoFindParams) DataResult {
	return param.findOne()
}

func UpdateOne(param UpdateInterface) DataResult {
	return param.findOneAndUpdate()
}

func InsertOne(param InsertInterface) DataResult {
	return param.insertOne()
	//return param.InsertMany()
}

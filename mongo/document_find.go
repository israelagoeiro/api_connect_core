package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type FindOptions struct {
	Sort bson.D
}

type FindParams struct {
	Collection string
	Connection string
	Database   string
	Filter     MongoFilter
	Fields     []string
	Options    FindOptions
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

type DocumentFind struct {
	Find    func() DataResult
	FindOne func() DataResult
}

func NewDocumentFind(param FindParams) DocumentFind {
	apiFields := NewFields(param.Fields)
	apiDocumentUpdate := DocumentFind{
		Find: func() DataResult {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			collection := GetCollection(param.Database, param.Collection)
			opts := options.Find()

			if param.Options.Sort != nil {
				opts.SetSort(param.Options.Sort)
			}

			if param.Fields != nil {
				opts.SetProjection(apiFields.Values())
			}

			filter := bson.D{}
			if param.Filter.Values != nil {
				filter = param.Filter.Values()
			}
			result, _ := collection.Find(ctx, filter, opts)

			return DataResult{
				_all:         result.All,
				_ctx:         ctx,
				_cancel:      cancel,
				_mongoResult: result,
				_debug: func() {
					var dataResult []any
					if err := result.All(ctx, &dataResult); err != nil {
						fmt.Println("Error Find:Debug()", err.Error())
					}
					err := result.Close(ctx)
					if err != nil {
						fmt.Println("Error Find:Debug()", err.Error())
						cancel()
					}
					fmt.Println("Find:Debug()", "Collection:"+param.Collection, dataResult)
				},
			}
		},
		FindOne: func() DataResult {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			collection := GetCollection(param.Database, param.Collection)
			opts := options.FindOne()

			if param.Options.Sort != nil {
				opts.SetSort(param.Options.Sort)
			}

			if param.Fields != nil {
				opts.SetProjection(apiFields.Values())
			}

			filter := bson.D{}
			if param.Filter.Values != nil {
				filter = param.Filter.Values()
			}
			result := collection.FindOne(ctx, filter, opts)

			return DataResult{
				_decode: result.Decode,
				_debug: func() {
					var dataResult any
					err := result.Decode(&dataResult)
					if err != nil {
						fmt.Println("Error FindOne:Debug()", err.Error())
						return
					}
					fmt.Println("FindOne:Debug()", "Collection:"+param.Collection, dataResult)
				},
			}
		},
	}
	return apiDocumentUpdate
}

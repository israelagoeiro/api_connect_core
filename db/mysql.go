package db

type MysqlUpdateParams struct {
	Action         string
	Table          string
	Connection     string
	Database       string
	Query          any
	ReturnOriginal bool
}

func (m MysqlUpdateParams) findOneAndUpdate() string {
	//	util.PrintStruct(m)
	return ""
}

type MysqlInsertParams struct {
	Collection string
	Connection string
	Database   string
	DataLog    DataLog
	//Input      MongoInput
}

func (param MysqlInsertParams) insertOne() DataResult {
	//document := NewMongoDocumentInsert(param)
	//return document.InsertOne()
	return DataResult{}
}

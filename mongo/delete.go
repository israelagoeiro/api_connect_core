package mongo

//DeleteInterface reune comandos de exclusão
type DeleteInterface interface {
	deleteOne() DataResult
	deleteMany() DataResult
}

func DeleteOne(param DeleteInterface) DataResult {
	return param.deleteOne()
}
func DeleteMany(param DeleteInterface) DataResult {
	return param.deleteMany()
}

func (param DeleteParams) deleteMany() DataResult {
	document := NewDeleteDocument(param)
	return document.DeleteMany()
}

func (param DeleteParams) deleteOne() DataResult {
	document := NewDeleteDocument(param)
	return document.DeleteOne()
}

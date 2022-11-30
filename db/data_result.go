package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type DataResult struct {
	_result        func() interface{}
	_decode        func(v interface{}) error
	_all           func(ctx context.Context, results interface{}) error
	_ctx           context.Context
	_debug         func()
	_matchedCount  int64 // The number of documents matched by the filter.
	_modifiedCount int64 // The number of documents modified by the operation.
	_upsertedCount int64 // The number of documents upserted by the operation.
	_upsertedID    interface{}
	_id            interface{}
	_mongoResult   *mongo.Cursor
	_cancel        context.CancelFunc
}

func (u DataResult) Model(model interface{}) {
	if u._decode != nil {
		err := u._decode(model)
		if err != nil {
			panic("UpdateResult::Decode - " + err.Error())
		}
	}
}
func (u DataResult) Models(models interface{}) {
	if u._all != nil {
		err := u._all(u._ctx, models)
		if err != nil {
			fmt.Println("Error Find:Models()", err.Error())
		}

		err = u._mongoResult.Close(u._ctx)
		if err != nil {
			fmt.Println("Error Find:Models()", err.Error())
			u._cancel()
		}
	}
}
func (u DataResult) Print() {
	u._debug()
}
func (u DataResult) API() []string {
	u._debug()
	return []string{""}
}

func (u DataResult) Result() interface{} {
	return u._result()
}
func (u DataResult) TotalAfected() int {
	return 0
}
func (u DataResult) Id() {
	return
}
func (u DataResult) Ids() {
}

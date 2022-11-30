package db

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/test"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestEq(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("gato", Eq("verde"))
	filter.Add("bola", Eq("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$eq", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$eq", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidEq(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestEqField(t *testing.T) {
	got := EqField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$eq", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidEq(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGt(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("gato", Gt("verde"))
	filter.Add("bola", Gt("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$gt", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$gt", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidGt(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestGtField(t *testing.T) {
	got := GtField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$gt", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidGt(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestGte(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("gato", Gte("verde"))
	filter.Add("bola", Gte("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$gte", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$gte", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidGte(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestGteField(t *testing.T) {
	got := GteField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$gte", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidGte(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLt(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("gato", Lt("verde"))
	filter.Add("bola", Lt("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$lt", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$lt", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidLt(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestLtField(t *testing.T) {
	got := LtField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$lt", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidLt(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestLte(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("gato", Lte("verde"))
	filter.Add("bola", Lte("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$lte", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$lte", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidLte(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestLteField(t *testing.T) {
	got := LteField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$lte", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidLte(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNe(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("gato", Ne("verde"))
	filter.Add("bola", Ne("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$ne", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$ne", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidNe(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestNeField(t *testing.T) {
	got := NeField("tags", "B")
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$ne", Value: "B"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidNe(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestIn(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("qty", In([]any{20, 30, 41}))
	filter.Add("type", In([]any{"casa", "azul", "verde"}))

	got := filter.Values()
	want := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$in", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$in", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidIn(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestInField(t *testing.T) {
	got := InField("tags", []any{20, 30, 41})
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$in", Value: []any{20, 30, 41}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidIn(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNin(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("qty", Nin([]any{20, 30, 41}))
	filter.Add("type", Nin([]any{"casa", "azul", "verde"}))

	got := filter.Values()
	want := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$nin", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$nin", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidNin(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestNinField(t *testing.T) {
	got := NinField("tags", []any{20, 30, 41})
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$nin", Value: []any{20, 30, 41}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidIn(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestExists(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("gato", Exists(false))
	filter.Add("bola", Exists(true))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$exists", Value: false}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$exists", Value: true}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidExists(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestExistsField(t *testing.T) {
	got := ExistsField("tags", false)
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$exists", Value: false}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidExists(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestType(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("gato", Type("verde"))
	filter.Add("bola", Type("azul"))

	got := filter.Values()
	want := bson.D{bson.E{Key: "gato", Value: bson.D{bson.E{Key: "$type", Value: "verde"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$type", Value: "azul"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidType(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestTypeField(t *testing.T) {
	got := TypeField("type", "car")
	want := bson.D{bson.E{Key: "type", Value: bson.D{bson.E{Key: "$type", Value: "car"}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidType(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestDivide(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("qty", Divide([]any{20, 30, 41}))
	filter.Add("type", Divide([]any{"casa", "azul", "verde"}))

	got := filter.Values()
	want := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$divide", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$divide", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidDivide(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestDivideField(t *testing.T) {
	got := DivideField("tags", []any{20, 30, 41})
	want := bson.D{bson.E{Key: "tags", Value: bson.D{bson.E{Key: "$divide", Value: []any{20, 30, 41}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidDivide(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestExpr(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("qty", Expr([]any{20, 30, 41}))
	filter.Add("type", Expr([]any{"casa", "azul", "verde"}))

	got := filter.Values()
	want := bson.D{bson.E{Key: "qty", Value: bson.D{bson.E{Key: "$expr", Value: []any{20, 30, 41}}}}, bson.E{Key: "type", Value: bson.D{bson.E{Key: "$expr", Value: []any{"casa", "azul", "verde"}}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidExpr(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

/*func TestExprField(t *testing.T) {
	got := ExprField("tags", []any{20, 30, 41})
	want := bson.D{bson.E{Key: "$expr", Value: bson.E{Key: "tags", Value: []any{20, 30, 41}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidExpr(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}*/

func TestNot(t *testing.T) {
	var filter = NewMongoFilter()
	filter.Add("item", Not("/^p.*/"))
	filter.Add("bola", Not(Regex("/^p.*/")))

	got := filter.Values()
	want := bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$not", Value: "/^p.*/"}}}, bson.E{Key: "bola", Value: bson.D{bson.E{Key: "$not", Value: Regex("/^p.*/")}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidNot(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestNotField(t *testing.T) {
	got := NotField("item", Gt(1.99))
	want := bson.D{bson.E{Key: "item", Value: bson.D{bson.E{Key: "$not", Value: Gt(1.99)}}}}
	fmt.Println("got", got)
	fmt.Println("want", want)

	if !test.IsValidNot(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

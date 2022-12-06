package mysql_test

import (
	"fmt"
	"github.com/israelagoeiro/api_connect_core/mysql"
	"testing"
)

func TestWhere(t *testing.T) {
	filter := mysql.NewFilter()
	filter.Where("last_name = 'Johnson'")

	expected := " WHERE (last_name = 'Johnson')"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestAnd(t *testing.T) {
	filter := mysql.NewFilter()
	filter.Where("last_name = 'Johnson'")
	filter.And("contact_id > 3000")

	expected := " WHERE (last_name = 'Johnson' AND contact_id > 3000)"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestOr(t *testing.T) {
	filter := mysql.NewFilter()
	filter.Where("last_name = 'Johnson'")
	filter.And("contact_id > 3000")
	filter.Or("contact_id > 3000")
	filter.Or("state = 'Florida'")

	expected := " WHERE (last_name = 'Johnson' AND contact_id > 3000) OR (contact_id > 3000) OR (state = 'Florida')"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestGroupBy(t *testing.T) {
	filter := mysql.NewFilter()
	filter.Where("last_name = 'Johnson'")
	filter.And("contact_id > 3000")
	filter.Or("contact_id > 3000")
	filter.Or("state = 'Florida'")
	filter.OrderBy("last_name = 'Johnson'")
	filter.GroupBy("last_name,contact_id,state")

	expected := " WHERE (last_name = 'Johnson' AND contact_id > 3000) OR (contact_id > 3000) OR (state = 'Florida') GROUP BY last_name,contact_id,state ORDER BY last_name = 'Johnson'"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestHaving(t *testing.T) {
	filter := mysql.NewFilter()
	filter.Where("last_name = 'Johnson'")
	filter.And("contact_id > 3000")
	filter.Or("contact_id > 3000")
	filter.Or("state = 'Florida'")
	filter.GroupBy("last_name,contact_id,state")
	filter.Having("MAX(salary) > 25000")

	expected := " WHERE (last_name = 'Johnson' AND contact_id > 3000) OR (contact_id > 3000) OR (state = 'Florida') GROUP BY last_name,contact_id,state HAVING MAX(salary) > 25000"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestOrderBy(t *testing.T) {
	filter := mysql.NewFilter()
	filter.Where("last_name = 'Johnson'")
	filter.And("contact_id > 3000")
	filter.Or("contact_id > 3000")
	filter.Or("state = 'Florida'")
	filter.GroupBy("last_name,contact_id,state")
	filter.Having("MAX(salary) > 25000")
	filter.OrderBy("last_name ASC, contact_id DESC")

	expected := " WHERE (last_name = 'Johnson' AND contact_id > 3000) OR (contact_id > 3000) OR (state = 'Florida') GROUP BY last_name,contact_id,state HAVING MAX(salary) > 25000 ORDER BY last_name ASC, contact_id DESC"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestLimitAndOffset(t *testing.T) {
	filter := mysql.NewFilter()
	filter.Where("last_name = 'Johnson'")
	filter.And("contact_id > 3000")
	filter.Or("contact_id > 3000")
	filter.Or("state = 'Florida'")
	filter.GroupBy("last_name,contact_id,state")
	filter.Having("MAX(salary) > 25000")
	filter.OrderBy("last_name ASC, contact_id DESC")
	filter.Limit(1, 0)

	expected := " WHERE (last_name = 'Johnson' AND contact_id > 3000) OR (contact_id > 3000) OR (state = 'Florida') GROUP BY last_name,contact_id,state HAVING MAX(salary) > 25000 ORDER BY last_name ASC, contact_id DESC LIMIT 1 OFFSET 0"
	obtained := filter.Values()

	if expected != obtained {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestInnerJoin(t *testing.T) {
	filter := mysql.NewFilter()
	filter.InnerJoin("orders", "suppliers.supplier_id = orders.supplier_id")

	expected := " INNER JOIN orders ON suppliers.supplier_id = orders.supplier_id"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestLeftJoin(t *testing.T) {
	filter := mysql.NewFilter()
	filter.LeftJoin("orders", "suppliers.supplier_id = orders.supplier_id")

	expected := " LEFT JOIN orders ON suppliers.supplier_id = orders.supplier_id"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

func TestRightJoin(t *testing.T) {
	filter := mysql.NewFilter()
	filter.RightJoin("orders", "suppliers.supplier_id = orders.supplier_id")

	expected := " RIGHT JOIN orders ON suppliers.supplier_id = orders.supplier_id"
	obtained := filter.Values()
	fmt.Println("expected", expected)
	fmt.Println("obtained", obtained)

	if len(expected) != len(obtained) {
		fmt.Println("expected", expected)
		fmt.Println("obtained", obtained)
		t.Errorf("expected %q, obtained %q", expected, obtained)
	}
}

package util

import (
	"fmt"
	"reflect"
)

func PrintStruct(m any) {
	v := reflect.ValueOf(m)
	typeOfS := v.Type()
	fmt.Println("//----------------------------")

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s:%v,\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
	fmt.Println("----------------------------//")
}

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsStr(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

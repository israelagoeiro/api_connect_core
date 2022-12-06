package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"reflect"
)

func LoadEnv(filename string) {
	// load .env file
	err := godotenv.Load(filename)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

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

func ToMap(values interface{}) map[string]any {
	result := map[string]any{}
	fmt.Println(values)
	/*for i := 0; i < reflect.ValueOf(values).Len(); i++ {
		field := fmt.Sprintf("%v", reflect.ValueOf(values).Index(i).Field(0))
		value := fmt.Sprintf("%v", reflect.ValueOf(values).Index(i).Field(1))
		result[field] = value
	}*/
	return result
}

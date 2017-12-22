package main

import (
	"fmt"
	"reflect"

	"github.com/wangming1993/share/golang/reflect/utils"
)

func main() {
	//callFunc()
	callPackagePublicFunc()
}

func callPackagePublicFunc() {
	var result []string

	scans(&result)

	fmt.Println(result)
}

func callFunc() {
	funcValue := reflect.ValueOf(utils.JSON)
	funcValue.Call(nil)

	fmt.Println(funcValue.Kind() == reflect.Func)

	method := reflect.Method{
		Name:    "GetJsonOutput",
		PkgPath: "github.com/wangming1993/share/golang/reflect/utils",
	}

	reflect.ValueOf(&method).Call(nil)
}

func mike() {
	fmt.Println("I'm Mike!")
}

func scans(result interface{}) {
	switch reflect.ValueOf(result).Kind() {
	case reflect.Ptr:
		re := makeSlice(reflect.TypeOf(result).Elem())

		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(re))
	}
}

func makeSlice(elemType reflect.Type) interface{} {
	if elemType.Kind() == reflect.Slice {
		elemType = elemType.Elem()
	}

	sliceType := reflect.SliceOf(elemType)

	v := reflect.MakeSlice(sliceType, 0, 0)

	for i := 0; i < 10; i++ {
		switch elemType.Kind() {
		case reflect.Int:
			v = reflect.Append(v, reflect.ValueOf(i))
		case reflect.String:
			v = reflect.Append(v, reflect.ValueOf(fmt.Sprintf("%d", i)))
		}

	}

	//fmt.Println(v)
	return v.Interface()
}

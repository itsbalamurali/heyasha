package engine

import (
	"reflect"
	"errors"
)

//http://stackoverflow.com/questions/18017979/golang-pointer-to-function-from-string-functions-name
func Call(fn map[string]interface{}, command string, params ... interface{})(result []reflect.Value, err error)  {
	f := reflect.ValueOf(fn[command])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
/*
// /Functions map
funcs := map[string]interface{} {
	"command": function,
	"command2": function2,
	//so on and so forth
}

*/
package core

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Run(structName, methodName string, input map[string]any) (any, error) {
	//1. Find registered struct by name
	instance, ok := registry[structName]
	if !ok {
		return nil, fmt.Errorf("struct %s not found", structName)
	}
	//2. Create a fresh pointer of the same type
	t := reflect.TypeOf(instance)
	v := reflect.New(t.Elem())

	//3. Fill with input
	inputBytes, _ := json.Marshal(input)
	if err := json.Unmarshal(inputBytes, v.Interface()); err != nil {
		return nil, fmt.Errorf("failed to unmarshal input %s : %w", structName, err)
	}

	//4. Find method by name
	method := reflect.ValueOf(v.Interface()).MethodByName(methodName)
	if !method.IsValid() {
		return nil, fmt.Errorf("method %s not found in struct %s", methodName, structName)
	}

	//5. Call method (only supports methods with no args)
	if method.Type().NumIn() != 0 || method.Type().NumOut() != 1 {
		return nil, fmt.Errorf("method %s in struct %s must have no arguments and return exactly one value", methodName, structName)
	}
	results := method.Call(nil)

	//6. Handle return (support 1 or 2 results)
	if len(results) == 1 {
		return results[0].Interface(), nil
	}
	if len(results) == 2 {
		if err, ok := results[1].Interface().(error); ok && err != nil {
			return nil, fmt.Errorf("method %s in struct %s returned error: %w", methodName, structName, err)
		}
		return results[0].Interface(), nil
	}
	return nil, fmt.Errorf("method %s in struct %s returned unexpected number of results", methodName, structName)

}

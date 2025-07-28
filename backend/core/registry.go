package core

import (
	"doc-kit/models"
	"fmt"
	"reflect"
)

var registry = make(map[string]any)

func RegisterStruct(name string, instance any) {
	if _, exists := registry[name]; exists {
		panic("struct already registered: " + name)
	}
	registry[name] = instance
}

func buildStructMeta(name string, instance any) models.StructMeta {
	t := reflect.TypeOf(instance)
	if t.Kind() != reflect.Pointer {
		t = t.Elem()
	}

	var fields []models.FieldMeta
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fields = append(fields, models.FieldMeta{
			Name:    field.Name,
			Type:    field.Type.String(),
			Tag:     string(field.Tag),
			Comment: "", // Comments are not directly accessible in Go, you might need to use a tool like go/doc

		})
	}
	var methods []models.MethodMeta
	typ := reflect.TypeOf(instance)
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		methods = append(methods, models.MethodMeta{
			Name: method.Name,
			Doc:  "", // Method documentation is not directly accessible in Go, you might need to use a tool like go/doc
		})
	}
	return models.StructMeta{
		Name:    name,
		Doc:     fmt.Sprintf("%s struct", name),
		Fields:  fields,
		Methods: methods,
	}
}

func GetAllStructs() []models.StructMeta {
	var structs []models.StructMeta
	for name, instance := range registry {
		meta := buildStructMeta(name, instance)
		structs = append(structs, meta)
	}
	return structs
}

func GetStructByName(name string) (models.StructMeta, bool) {
	if instance, ok := registry[name]; ok {
		return buildStructMeta(name, instance), true
	}
	return models.StructMeta{}, false
}

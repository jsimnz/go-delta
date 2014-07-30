package delta

import (
	"errors"
	"reflect"
)

var (
	ErrNotStruct = errors.New("The given type is not a struct")
	WarnDiffType = errors.New("Warning given types are not of the same type")
)

type fields map[string]interface{}

type structInfo struct {
	sInterface interface{}
	sValue     reflect.Value
	sType      reflect.Type
	fields     fields
}

func getStructInfo(val interface{}) (*structInfo, error) {
	if !isStruct(val) {
		return nil, ErrNotStruct
	}

	info := &structInfo{
		sInterface: val,
		sValue:     reflect.ValueOf(val),
		sType:      reflect.TypeOf(val),
		fields:     make(fields),
	}

	numFields := info.sValue.NumField()
	for i := 0; i < numFields; i++ {
		field := info.sType.Field(i)
		info.fields[field.Name] = info.sValue.Field(i).Interface()
	}

	return info, nil
}

func isStruct(val interface{}) bool {
	return reflect.ValueOf(val).Kind() == reflect.Struct
}

// Find the delta between 2 structs
func Struct(s1, s2 interface{}) (map[string]interface{}, error) {
	var err error

	s1info, err = getStructInfo(s1)
	if err != nil {
		return nil, err
	}
	s2info, err = getStructInfo(s2)
	if err != nil {
		return nil, err
	}
}

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
	Fields fields

	sInterface interface{}
	sValue     reflect.Value
	sType      reflect.Type
}

func getStructInfo(val interface{}) (*structInfo, error) {
	if !isStruct(val) {
		return nil, ErrNotStruct
	}

	info := &structInfo{
		sInterface: val,
		sValue:     reflect.ValueOf(val),
		sType:      reflect.TypeOf(val),
		Fields:     make(fields),
	}

	numFields := info.sValue.NumField()
	for i := 0; i < numFields; i++ {
		field := info.sType.Field(i)
		info.Fields[field.Name] = info.sValue.Field(i).Interface()
	}

	return info, nil
}

func isStruct(val interface{}) bool {
	return reflect.ValueOf(val).Kind() == reflect.Struct
}

// Find the delta between 2 structs
func Struct(base, compare interface{}) (map[string]interface{}, error) {
	var err error

	baseInfo, err := getStructInfo(base)
	if err != nil {
		return nil, err
	}
	compareInfo, err := getStructInfo(compare)
	if err != nil {
		return nil, err
	}

	diffFields := make(fields)
	for fieldName := range compareInfo.Fields {
		comparefieldVal := compareInfo.Fields[fieldName]
		basefieldVal, exists := baseInfo.Fields[fieldName]
		if exists {
			if !equal(comparefieldVal, basefieldVal) {
				diffFields[fieldName] = comparefieldVal
			}
		} else {
			diffFields[fieldName] = comparefieldVal
		}
	}

	return diffFields, nil
}

// Returns true if both parameters have the same type and value
// false otherwise
func equal(v1, v2 interface{}) bool {
	if reflect.TypeOf(v1).Kind() != reflect.TypeOf(v2).Kind() {
		return false
	} else if !reflect.DeepEqual(v1, v2) {
		return false
	}
	return true
}

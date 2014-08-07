package delta

import (
	"fmt"
	"testing"
)

type structDeltaSimple struct {
	Field1 string
	Field2 string
}

type DeltaEmbeddedStruct struct {
	Test []EmbeddedStruct
}

type EmbeddedStruct struct {
	DeeperTest string
}

func TestStructSimpleDeltaEqual(t *testing.T) {
	s1 := structDeltaSimple{"value1", "value2"}
	s2 := structDeltaSimple{"value1", "value2"}

	diff, err := Struct(s1, s2)
	if err != nil {
		t.Errorf("Struct returned an error where one shouldn't be: %v", err)
	}
	if len(diff) > 0 {
		t.Errorf("Should be no differences between structs, got: %v", diff)
	}
}

func TestStructDeltaNotEqual(t *testing.T) {
	s1 := structDeltaSimple{"value1", "value2"}
	s2 := structDeltaSimple{"value1", "value3"}

	diff, err := Struct(s1, s2)
	if err != nil {
		t.Errorf("Struct returned an error where one shouldn't be: %v", err)
	}
	if len(diff) != 1 {
		t.Errorf("Should be no differences between structs, got: %v", diff)
	}
	if diff["Field2"] != "value3" {
		t.Errorf("Struct reported wrong diff for struct, should be 'value3', got: '%v'", diff["Field2"])
	}
}

func TestEmbeddedStruct(t *testing.T) {
	s1 := DeltaEmbeddedStruct{}
	s2 := DeltaEmbeddedStruct{Test: []EmbeddedStruct{EmbeddedStruct{DeeperTest: "HELLO WORLD"}}}

	diff, err := Struct(s1, s2)
	if err != nil {
		t.Errorf("Struct returned an error: %v", err)
	}
	fmt.Println(diff)
}

package register

import (
	"github.com/fourhundredfour/dumper/tree"
	"testing"
)

func TestRegister_GenerateMapHash (t *testing.T) {
	mockTree := tree.Tree{Table: "mock", Column: "mock_column", Value: "fourhundredfour"}
	expectedValue := mockTree.Table + "_" + mockTree.Column + "_" + mockTree.Value
	instance := Register{}
	if expectedValue != instance.generateMapHash(&mockTree) {
		t.Fatal("Map Hash is not as expected")
	}
}

func TestRegister_Add(t *testing.T) {
	mockTree := tree.Tree{Table: "mock", Column: "mock_column", Value: "fourhundredfour"}
	instance := New()
	instance.Add(&mockTree)

	if _, ok := instance.trees[instance.generateMapHash(&mockTree)]; !ok {
		t.Fatal("Tree was not added to the register")
	}
}

func TestRegister_Exists(t *testing.T) {
	mockTree := tree.Tree{Table: "mock", Column: "mock_column", Value: "fourhundredfour"}
	secondMockTree := tree.Tree{Table: "mock", Column: "mock_column", Value: "fourhundredfour2"}
	instance := New()
	instance.Add(&mockTree)

	if ok := instance.Exists(&mockTree); !ok {
		t.Fatal("Tree was not added to the register")
	}
	if ok := instance.Exists(&secondMockTree); ok {
		t.Fatal("Tree was not added but was found in the register")
	}
}
package main

import (
	"testing"
)

func TestPowerSet_Put(t *testing.T) {
	var result = PowerSet[string]{}

	result.Put("1")
	if len(result.values) != 1 {
		t.Error("Error")
	}

	result.Put("2")
	if len(result.values) != 2 {
		t.Error("Error")
	}

	result.Put("3")
	if len(result.values) != 3 {
		t.Error("Error")
	}
	result.Put("3")
	if len(result.values) != 3 {
		t.Error("Error")
	}
	result.Put("4")
	if len(result.values) != 4 {
		t.Error("Error")
	}

	/*var resultString = PowerSet[string]{}
	for i := 0; i < 10000; i++ {
		var strConvert = fmt.Sprint(i)
		resultString.Put(strConvert)
	}
	if len(resultString.values) != 10000 {
		t.Error("Error")
	}*/

}

func TestPowerSet_Size(t *testing.T) {
	var result = PowerSet[string]{}

	result.Put("1")
	if result.Size() != 1 {
		t.Error("Error")
	}

	result.Put("2")
	if result.Size() != 2 {
		t.Error("Error")
	}

	result.Put("3")
	if result.Size() != 3 {
		t.Error("Error")
	}
	result.Put("3")
	if result.Size() != 3 {
		t.Error("Error")
	}
	result.Put("4")
	if result.Size() != 4 {
		t.Error("Error")
	}
}

func TestPowerSet_Get(t *testing.T) {
	var result = PowerSet[string]{}

	result.Put("1")
	result.Put("2")
	result.Put("3")

	var answer = result.Get("2")
	if answer != true {
		t.Error("Error")
	}

	answer = result.Get("1")
	if answer != true {
		t.Error("Error")
	}

	answer = result.Get("100")
	if answer != false {
		t.Error("Error")
	}
}

func TestPowerSet_Remove(t *testing.T) {
	var result = PowerSet[string]{}

	result.Put("1")
	result.Put("2")
	result.Put("3")

	var answer = result.Remove("2")
	if answer != true {
		t.Error("Error")
	}

	answer = result.Remove("3")
	if answer != true {
		t.Error("Error")
	}

	answer = result.Remove("3")
	if answer != false {
		t.Error("Error")
	}
	answer = result.Remove("2")
	if answer != false {
		t.Error("Error")
	}

	answer = result.Remove("1")
	if answer != true {
		t.Error("Error")
	}
	if result.Size() != 0 {
		t.Error("Error")
	}

}

func TestPowerSet_Intersection(t *testing.T) {
	var set1 = PowerSet[string]{}
	var set2 = PowerSet[string]{}

	set1.Put("1")
	set1.Put("2")
	set1.Put("3")
	set1.Put("4")
	set1.Put("5")
	set1.Put("6")

	set2.Put("2")
	set2.Put("4")
	set2.Put("6")
	set2.Put("8")
	set2.Put("10")

	var answer = set1.Intersection(set2)
	if answer.Size() != 3 {
		t.Error("Error")
	}

	if answer.values[0] != "2" {
		t.Error("Error")
	}

	if answer.values[1] != "4" {
		t.Error("Error")
	}

	if answer.values[2] != "6" {
		t.Error("Error")
	}

}

func TestPowerSet_Union(t *testing.T) {
	var set1 = PowerSet[string]{}
	var set2 = PowerSet[string]{}

	set1.Put("1")
	set1.Put("2")
	set1.Put("3")
	set1.Put("4")
	set1.Put("5")
	set1.Put("6")

	set2.Put("2")
	set2.Put("4")
	set2.Put("6")
	set2.Put("8")
	set2.Put("10")

	var answer = set1.Union(set2)
	if answer.Size() != 8 {
		t.Error("Error")
	}

	if answer.values[0] != "1" {
		t.Error("Error")
	}

	if answer.values[7] != "10" {
		t.Error("Error")
	}

}

func TestPowerSet_Difference(t *testing.T) {
	var set1 = PowerSet[string]{}
	var set2 = PowerSet[string]{}

	set1.Put("1")
	set1.Put("2")
	set1.Put("3")
	set1.Put("4")
	set1.Put("5")
	set1.Put("6")

	set2.Put("2")
	set2.Put("4")
	set2.Put("6")
	set2.Put("8")
	set2.Put("10")

	var answer = set1.Difference(set2)
	if answer.Size() != 3 {
		t.Error("Error")
	}

	if answer.values[0] != "1" {
		t.Error("Error")
	}

	if answer.values[1] != "3" {
		t.Error("Error")
	}

	if answer.values[2] != "5" {
		t.Error("Error")
	}
}

func TestPowerSet_IsSubset(t *testing.T) {
	var set1 = PowerSet[string]{}
	var set2 = PowerSet[string]{}

	set1.Put("1")
	set1.Put("2")
	set1.Put("3")
	set1.Put("4")
	set1.Put("5")
	set1.Put("6")

	set2.Put("2")
	set2.Put("4")
	set2.Put("6")

	var answer = set1.IsSubset(set2)
	if answer != true {
		t.Error("Error")
	}

	set2.Put("8")

	answer = set1.IsSubset(set2)
	if answer != false {
		t.Error("Error")
	}

}

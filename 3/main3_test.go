package main

import (
	"reflect"
	"testing"
)

func TestAddAndGet(t *testing.T) {
	sim := StringIntMap{data: make(map[string]int)}

	sim.Add("apple", 10)

	val, ok := sim.Get("apple")
	if !ok {
		t.Errorf("Get() ok = %v; want true (key 'apple' should exist)", ok)
	}
	if val != 10 {
		t.Errorf("Get() value = %d; want %d", val, 10)
	}
}

func TestRemove(t *testing.T) {
	sim := StringIntMap{data: map[string]int{"banana": 5}}

	sim.Remove("banana")

	if sim.Exists("banana") {
		t.Errorf("Exists() = true; want false after removing key 'banana'")
	}
}

func TestExists(t *testing.T) {
	sim := StringIntMap{data: map[string]int{"cherry": 15}}

	if !sim.Exists("cherry") {
		t.Errorf("Exists() = false; want true for existing key 'cherry'")
	}

	if sim.Exists("unknown") {
		t.Errorf("Exists() = true; want false for non-existent key")
	}
}

func TestCopy(t *testing.T) {
	originalData := map[string]int{"a": 1, "b": 2}
	sim := StringIntMap{data: originalData}

	copiedMap := sim.Copy()

	if !reflect.DeepEqual(originalData, copiedMap) {
		t.Errorf("Copy() mismatch: got %v, want %v", copiedMap, originalData)
	}

	sim.Add("c", 3)
	if _, ok := copiedMap["c"]; ok {
		t.Errorf("Copy() is a reference: reflected changes from original map")
	}
}

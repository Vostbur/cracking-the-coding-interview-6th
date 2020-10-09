// linked_list_test.py
package main

import (
	"reflect"
	"testing"
)

func TestAddMultiple(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 3, 4, 5}
	ll.AddMultiple(data)
	if actual := ll.GetSlice(); !reflect.DeepEqual(actual, data) {
		t.Errorf("Expected: %v. Actual: %v", data, actual)
	}
}

func TestAddToHeadMultiple(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 3, 4, 5}
	ll.AddToHeadMultiple(data)
	if actual := ll.GetSlice(); !reflect.DeepEqual(actual, data) {
		t.Errorf("Expected: %v. Actual: %v", data, actual)
	}
}

func TestDel(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 3, 4, 5}
	dataAfterDelete := []int{2, 4}
	ll.AddMultiple(data)
	if err := ll.Del(1); err != nil {
		t.Errorf(err.Error())
	}
	if err := ll.Del(5); err != nil {
		t.Errorf(err.Error())
	}
	if err := ll.Del(3); err != nil {
		t.Errorf(err.Error())
	}
	if actual := ll.GetSlice(); !reflect.DeepEqual(actual, dataAfterDelete) {
		t.Errorf("Expected: %v. Actual: %v", dataAfterDelete, actual)
	}
}

func TestDelByIndex(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 3, 4, 5}
	dataAfterDelete := []int{2, 4}
	ll.AddMultiple(data)
	if err := ll.DelByIndex(0); err != nil {
		t.Errorf(err.Error())
	}
	if err := ll.DelByIndex(3); err != nil {
		t.Errorf(err.Error())
	}
	if err := ll.DelByIndex(1); err != nil {
		t.Errorf(err.Error())
	}
	if actual := ll.GetSlice(); !reflect.DeepEqual(actual, dataAfterDelete) {
		t.Errorf("Expected: %v. Actual: %v", dataAfterDelete, actual)
	}
}

func TestGetReverseSlice(t *testing.T) {
	var ll LinkedList
	data := []int{1, 2, 3, 4, 5}
	dataReverse := []int{5, 4, 3, 2, 1}
	ll.AddMultiple(data)
	if actual := ll.GetReverseSlice(); !reflect.DeepEqual(actual, dataReverse) {
		t.Errorf("Expected: %v. Actual: %v", dataReverse, actual)
	}
}

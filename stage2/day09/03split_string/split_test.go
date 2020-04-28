package split_string

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("babcbef", "b")
	want := []string{"", "a", "c", "ef"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect %v,but got %v", want, got)
	}
}

func TestSplit2(t *testing.T) {
	got := Split("aabbcc", "bb")
	want := []string{"aa", "cc"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("expect %v,but got %v", want, got)
	}
}

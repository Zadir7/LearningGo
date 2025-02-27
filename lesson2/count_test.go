package lesson2

import "testing"

func TestCountReturnsZeroOnEmptyString(t *testing.T) {
	str := ""
	expected := 0

	if cnt := Count(str, ' '); cnt != expected {
		t.Fatalf("bad count, expected 0, actual %d", cnt)
	}
}

func TestCountReturnsZero(t *testing.T) {
	str := "aab ccc"
	expected := 3

	if cnt := Count(str, 'c'); cnt != expected {
		t.Fatalf("bad count, expected %d, actual %d", expected, cnt)
	}
}

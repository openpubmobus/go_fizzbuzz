package main

import "testing"

func TestHello(t *testing.T) {
	actual := Hello(1)
	expected := "1"

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}

package main

import "testing"

func TestRegularNumbersReturnThemselves(t *testing.T) {
	actual := FizzBuzz(1)
	expected := "1"

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}

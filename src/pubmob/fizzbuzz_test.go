package main

import "testing"

func TestRegularNumbersReturnThemselves(t *testing.T) {
	actual := FizzBuzz(1)
	expected := "1"

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}

func TestAnotherTest(t *testing.T) {
	actual := FizzBuzz(2)
	expected := "2"

	if actual != expected {
		t.Errorf("actual %q expected %q", actual, expected)
	}
}

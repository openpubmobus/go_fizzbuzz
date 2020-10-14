package main

import (
	"testing"

	. "github.com/stretchr/testify/assert"
)

func TestRegularNumbersReturnThemselves(t *testing.T) {
	Equal(t, "1", FizzBuzz(1))
	Equal(t, "2", FizzBuzz(2))
}

func TestNumbersDivisibleBy3(t *testing.T) {
	Equal(t, "Fizz", FizzBuzz(3))
	Equal(t, "Fizz", FizzBuzz(27))
}

func TestNumbersDivisibleBy5(t *testing.T) {
	Equal(t, "Buzz", FizzBuzz(5))
	Equal(t, "Buzz", FizzBuzz(10))
}

func TestNumbersDivisibleBy3and5(t *testing.T) {
	Equal(t, "FizzBuzz", FizzBuzz(15))
	Equal(t, "FizzBuzz", FizzBuzz(30))
}

func TestOneThroughFifteen(t *testing.T) {
	Equal(t, "1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz",
		TempName(15))
}

func TempName(limit int) string {
	s := FizzBuzz(1)
	for i := 2; i <= limit; i++ {
		s += " " + FizzBuzz(i)
	}
	return s
}

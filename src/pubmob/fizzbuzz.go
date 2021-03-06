package pubmob

import (
	"fmt"
	"rsc.io/quote"
	"strconv"
)

const Fizz = "Fizz"
const Buzz = "Buzz"

func FizzBuzz(number int) string {

	fmt.Println(quote.Hello())
	divisibleBy3 := number % 3 == 0
	divisibleBy5 := number % 5 == 0

	if divisibleBy3 && divisibleBy5 {
		return Fizz + Buzz
	}
	if divisibleBy3 {
		return Fizz
	}
	if divisibleBy5 {
		return Buzz
	}

	return strconv.Itoa(number)
}

package main


import "github.com/pubmob"

// TempName is being exported
func TempName(limit int) string {
	s := pubmob.FizzBuzz(1)
	for i := 2; i <= limit; i++ {
		s += " " + pubmob.FizzBuzz(i)
	}
	return s
}

func main() {
	TempName(20)
}

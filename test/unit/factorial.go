package factorial

import (
	"fmt"
)

type myError struct {
	err string
}

func (myError *myError) Error() string {
	return fmt.Sprintf("Error:%s", myError.err)
}

func newError(info string) (err *myError) {
	return &myError{info}
}

// Factorial function
func Factorial(n int) (facVal uint64, err error) {
	if n < 0 {
		err = newError("Factorial of negative number doesn't exist.")
		return
	}

	facVal = 1
	for i := 1; i <= n; i++ {
		facVal *= uint64(i)
	}
	return

}

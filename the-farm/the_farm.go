package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows int
	msg  string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.cows, err.msg)
}

// TODO: define the 'DivideFood' function
func DivideFood(t FodderCalculator, nCows int) (float64, error) {
	var err error
	amount, err := t.FodderAmount(nCows)
	if err != nil {
		return 0, err
	}
	fac, err := t.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount / float64(nCows) * fac, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(t FodderCalculator, nCows int) (float64, error) {
	if nCows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(t, nCows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows: cows,
			msg:  "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			cows: cows,
			msg:  "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

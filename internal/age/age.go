package age

import (
	"errors"
	"fmt"
)

type age int

func FromInt(i int) (err error, a age) {
	if i < 0 || i > 120 {
		err = errors.New("Age must be between 0 and 120")
		return err, a
	}

	return err, age(i)
}

func (a age) HowManyCandles() string {

	return fmt.Sprintf("You are %d years old, so you need %d candles", a, a)
}

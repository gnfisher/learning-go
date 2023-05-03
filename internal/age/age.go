package age

import (
	"errors"
)

type age int

func FromInt(i int) (err error, a age) {
	if i > 120 {
		err = errors.New("age is too high")
		return err, a
	}

	return err, age(i)
}

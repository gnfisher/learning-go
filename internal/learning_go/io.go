package learning_go

import (
	"io"
	"unicode"
)

func CountChars(r io.Reader) (map[string]int, error) {
	// 16 byte buffer size
	buf := make([]byte, 16)
	// Where we store the counts
	out := map[string]int{}

	for {
		n, err := r.Read(buf)

		// return if it is end of file
		if err == io.EOF {
			return out, nil
		}

		// Return the error if its anything else
		if err != nil {
			return nil, err
		}

		// Each byte into the map
		for _, b := range buf[:n] {
			if unicode.IsLetter(rune(b)) {
				out[string(b)]++
			}
		}
	}
}

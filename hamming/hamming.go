package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("input strings not equal")
	}

	var d int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}

	return d, nil
}

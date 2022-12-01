package hlp

import "log"

func Must[T interface{}](a T, err error) T {
	if err != nil {
		log.Fatal(err)
	}

	return a
}

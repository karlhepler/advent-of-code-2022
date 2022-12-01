package hlp

import "log"

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Must2[T interface{}](a T, err error) T {
	if err != nil {
		log.Fatal(err)
	}

	return a
}

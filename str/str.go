package str

import "strconv"

func AtoiSlice(a []string) ([]int, error) {
	n := len(a)
	out := make([]int, n)
	var err error

	for i := 0; i < n; i++ {
		out[i], err = strconv.Atoi(a[i])
		if err != nil {
			return nil, err
		}
	}

	return out, nil
}

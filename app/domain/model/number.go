package model

import "strconv"

type Number struct {
	Value string
}

func (n *Number) ToInt() (int, error) {
	number, err := strconv.Atoi(n.Value)
	if err != nil {
		return 0, err
	}

	return number, nil
}

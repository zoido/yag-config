package yag

import "strconv"

type float32Value struct {
	dest *float32
}

func (iv *float32Value) Set(val string) error {
	num, err := strconv.ParseFloat(val, 32)
	if err != nil {
		return err
	}

	*iv.dest = float32(num)
	return nil
}

func (iv *float32Value) String() string {
	return strconv.FormatFloat(float64(*iv.dest), 'G', -1, 32)
}

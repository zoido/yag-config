package yag

import "strconv"

type intValue struct {
	dest *int
}

func (iv *intValue) Set(val string) error {
	num, err := strconv.Atoi(val)
	if err != nil {
		return err
	}

	*iv.dest = num
	return nil
}

func (iv *intValue) String() string {
	return strconv.Itoa(*iv.dest)
}

type int32Value struct {
	dest *int32
}

func (iv *int32Value) Set(val string) error {
	num, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		return err
	}

	*iv.dest = int32(num)
	return nil
}

func (iv *int32Value) String() string {
	return strconv.FormatInt(int64(*iv.dest), 10)
}

type int64Value struct {
	dest *int64
}

func (iv *int64Value) Set(val string) error {
	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return err
	}

	*iv.dest = num
	return nil
}

func (iv *int64Value) String() string {
	return strconv.FormatInt(*iv.dest, 10)
}

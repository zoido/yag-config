package yag

import "strconv"

type uintValue struct {
	dest *uint
}

func (iv *uintValue) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 0)
	if err != nil {
		return err
	}

	*iv.dest = uint(num)
	return nil
}

func (iv *uintValue) String() string {
	return strconv.FormatUint(uint64(*iv.dest), 10)
}

type uint8Value struct {
	dest *uint8
}

func (iv *uint8Value) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return err
	}

	*iv.dest = uint8(num)
	return nil
}

func (iv *uint8Value) String() string {
	return strconv.FormatUint(uint64(*iv.dest), 10)
}

type uint16Value struct {
	dest *uint16
}

func (iv *uint16Value) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		return err
	}

	*iv.dest = uint16(num)
	return nil
}

func (iv *uint16Value) String() string {
	return strconv.FormatUint(uint64(*iv.dest), 10)
}

type uint32Value struct {
	dest *uint32
}

func (iv *uint32Value) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return err
	}

	*iv.dest = uint32(num)
	return nil
}

func (iv *uint32Value) String() string {
	return strconv.FormatUint(uint64(*iv.dest), 10)
}

type uint64Value struct {
	dest *uint64
}

func (iv *uint64Value) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}

	*iv.dest = num
	return nil
}

func (iv *uint64Value) String() string {
	return strconv.FormatUint(*iv.dest, 10)
}

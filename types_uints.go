package yag

import "strconv"

type uintValue struct {
	dest *uint
}

func (uiv *uintValue) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 0)
	if err != nil {
		return err
	}

	*uiv.dest = uint(num)
	return nil
}

func (uiv *uintValue) String() string {
	return strconv.FormatUint(uint64(*uiv.dest), 10)
}

type uint8Value struct {
	dest *uint8
}

func (uiv *uint8Value) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		return err
	}

	*uiv.dest = uint8(num)
	return nil
}

func (uiv *uint8Value) String() string {
	return strconv.FormatUint(uint64(*uiv.dest), 10)
}

type uint16Value struct {
	dest *uint16
}

func (uiv *uint16Value) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		return err
	}

	*uiv.dest = uint16(num)
	return nil
}

func (uiv *uint16Value) String() string {
	return strconv.FormatUint(uint64(*uiv.dest), 10)
}

type uint32Value struct {
	dest *uint32
}

func (uiv *uint32Value) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return err
	}

	*uiv.dest = uint32(num)
	return nil
}

func (uiv *uint32Value) String() string {
	return strconv.FormatUint(uint64(*uiv.dest), 10)
}

type uint64Value struct {
	dest *uint64
}

func (uiv *uint64Value) Set(val string) error {
	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return err
	}

	*uiv.dest = num
	return nil
}

func (uiv *uint64Value) String() string {
	return strconv.FormatUint(*uiv.dest, 10)
}

package args

import (
	"github.com/zoido/yag-config/value"
)

// Ints returns new instance of Parser implementation that parses int values.
func Ints(i *[]int) Parser {
	return &intsParser{dest: i}
}

type intsParser struct {
	dest *[]int
}

func (ip *intsParser) Parse(values []string) (int, error) {
	var n int
	outs := make([]int, len(values))
	for i, v := range values {
		n = i + 1
		o := value.Int(&outs[i])
		if err := o.Set(v); err != nil {
			return i, err
		}
	}
	*ip.dest = outs
	return n, nil
}

// Int8s returns new instance of Parser implementation that parses int8 values.
func Int8s(i8 *[]int8) Parser {
	return &int8sParser{dest: i8}
}

type int8sParser struct {
	dest *[]int8
}

func (ip *int8sParser) Parse(values []string) (int, error) {
	var n int
	outs := make([]int8, len(values))
	for i, v := range values {
		n = i + 1
		o := value.Int8(&outs[i])
		if err := o.Set(v); err != nil {
			return i, err
		}
	}
	*ip.dest = outs
	return n, nil
}

// Int16s returns new instance of Parser implementation that parses int16 values.
func Int16s(i16 *[]int16) Parser {
	return &int16sParser{dest: i16}
}

type int16sParser struct {
	dest *[]int16
}

func (ip *int16sParser) Parse(values []string) (int, error) {
	var n int
	outs := make([]int16, len(values))
	for i, v := range values {
		n = i + 1
		o := value.Int16(&outs[i])
		if err := o.Set(v); err != nil {
			return i, err
		}
	}
	*ip.dest = outs
	return n, nil
}

// Int32s returns new instance of Parser implementation that parses int32 values.
func Int32s(i32 *[]int32) Parser {
	return &int32sParser{dest: i32}
}

type int32sParser struct {
	dest *[]int32
}

func (ip *int32sParser) Parse(values []string) (int, error) {
	var n int
	outs := make([]int32, len(values))
	for i, v := range values {
		n = i + 1
		o := value.Int32(&outs[i])
		if err := o.Set(v); err != nil {
			return i, err
		}
	}
	*ip.dest = outs
	return n, nil
}

// Int64s returns new instance of Parser implementation that parses int64 values.
func Int64s(i64 *[]int64) Parser {
	return &int64sParser{dest: i64}
}

type int64sParser struct {
	dest *[]int64
}

func (ip *int64sParser) Parse(values []string) (int, error) {
	var n int
	outs := make([]int64, len(values))
	for i, v := range values {
		n = i + 1
		o := value.Int64(&outs[i])
		if err := o.Set(v); err != nil {
			return i, err
		}
	}
	*ip.dest = outs
	return n, nil
}

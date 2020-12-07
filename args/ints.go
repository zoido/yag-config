package args

import (
	"fmt"

	"github.com/zoido/yag-config/value"
)

// Ints returns new instance of Parser implementation that parses int values.
func Ints(i *[]int) Parser {
	return &intsParser{dest: i}
}

type intsParser struct {
	dest *[]int
}

func (sp *intsParser) Parse(values []string) error {
	outs := make([]int, len(values))
	for i, v := range values {
		o := value.Int(&outs[i])
		err := o.Set(v)
		if err != nil {
			return fmt.Errorf("parsing int argument on position %d", i+1)
		}
	}
	*sp.dest = outs
	return nil
}

// Int8s returns new instance of Parser implementation that parses int8 values.
func Int8s(i8 *[]int8) Parser {
	return &int8sParser{dest: i8}
}

type int8sParser struct {
	dest *[]int8
}

func (sp *int8sParser) Parse(values []string) error {
	outs := make([]int8, len(values))
	for i, v := range values {
		o := value.Int8(&outs[i])
		err := o.Set(v)
		if err != nil {
			return fmt.Errorf("parsing int8 argument on position %d", i+1)
		}
	}
	*sp.dest = outs
	return nil
}

// Int16s returns new instance of Parser implementation that parses int16 values.
func Int16s(i16 *[]int16) Parser {
	return &int16sParser{dest: i16}
}

type int16sParser struct {
	dest *[]int16
}

func (sp *int16sParser) Parse(values []string) error {
	outs := make([]int16, len(values))
	for i, v := range values {
		o := value.Int16(&outs[i])
		err := o.Set(v)
		if err != nil {
			return fmt.Errorf("parsing int16 argument on position %d", i+1)
		}
	}
	*sp.dest = outs
	return nil
}

// Int32s returns new instance of Parser implementation that parses int32 values.
func Int32s(i32 *[]int32) Parser {
	return &int32sParser{dest: i32}
}

type int32sParser struct {
	dest *[]int32
}

func (sp *int32sParser) Parse(values []string) error {
	outs := make([]int32, len(values))
	for i, v := range values {
		o := value.Int32(&outs[i])
		err := o.Set(v)
		if err != nil {
			return fmt.Errorf("parsing int32 argument on position %d", i+1)
		}
	}
	*sp.dest = outs
	return nil
}

// Int64s returns new instance of Parser implementation that parses int64 values.
func Int64s(i64 *[]int64) Parser {
	return &int64sParser{dest: i64}
}

type int64sParser struct {
	dest *[]int64
}

func (sp *int64sParser) Parse(values []string) error {
	outs := make([]int64, len(values))
	for i, v := range values {
		o := value.Int64(&outs[i])
		err := o.Set(v)
		if err != nil {
			return fmt.Errorf("parsing int64 argument on position %d", i+1)
		}
	}
	*sp.dest = outs
	return nil
}

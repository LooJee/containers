package set

import "fmt"

type InvalidDataTypeErr struct {
	DataType string
}

func (e *InvalidDataTypeErr) Error() string {
	return fmt.Sprintf("not support data type : %s", e.DataType)
}

type UnsuitableTypeErr struct {
	Want string
	Got  string
}

func (e *UnsuitableTypeErr) Error() string {
	return fmt.Sprintf("data type not suitable, want: %s, got: %s", e.Want, e.Got)
}

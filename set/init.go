package set

import "reflect"

var (
	validKind = make(map[reflect.Kind]struct{})
)

func init() {
	for i := reflect.Int; i <= reflect.Float64; i++ {
		validKind[i] = struct{}{}
	}

	validKind[reflect.String] = struct{}{}
}

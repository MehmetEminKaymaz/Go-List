package Lists

import "reflect"

func(q Query) ToSlice() (slice interface{}){
	return reflect.Indirect(q.v)
}

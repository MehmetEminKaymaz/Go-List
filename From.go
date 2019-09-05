package Lists

import (
	"reflect"
)

type Query struct {
	v reflect.Value
}


func From(i interface{}) Query{
	val:=reflect.ValueOf(i)
    if val.Kind()!=reflect.Slice{
    	return Query{}
	}
	return Query{
		v:val,
	}
}

func(q *Query) Len() int{
	return q.v.Len()
}





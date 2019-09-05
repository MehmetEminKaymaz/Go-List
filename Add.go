package Lists

import "reflect"

func(q *Query) Add(x interface{}) {
	q.v=reflect.Append(q.v,reflect.ValueOf(x))
}

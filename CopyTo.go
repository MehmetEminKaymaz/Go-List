package Lists


func(q *Query) CopyTo(slice interface{}) {
	reflect.Copy(reflect.Indirect(reflect.ValueOf(slice)),q.v)

}
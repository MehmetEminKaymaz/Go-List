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















/*func(q *Query) OfType(typ interface{}) Query{
	mytyp:=reflect.TypeOf(typ)

	newSlice:=reflect.MakeSlice(reflect.SliceOf(mytyp),0,1)

	for i:=0;i<q.v.Len();i++{
		if mytyp==q.v.Index(i).Type(){
			newSlice=reflect.Append(newSlice,q.v.Index(i))
		}

	}

	return Query{
		v:newSlice,
	}
}*/



package Lists

import "reflect"

func(q *Query) Exist(slice interface{})  (state bool){
	slicev:=reflect.ValueOf(slice)
	state =true
	for i:=0;i<slicev.Len();i++{
		for k:=0; k<q.v.Len();k++{
			if q.v.Index(k).Interface()!=slicev.Index(i).Interface(){
				state=false
				return
			}
		}
	}
	return

}

package Lists

func(q *Query) Any(look func(i interface{}) bool ) (state bool){
	state=false
	for i:=0 ;i<q.v.Len();i++{
		if look(q.v.Index(i).Interface()){
			state=true
			return
		}

	}

	return
}

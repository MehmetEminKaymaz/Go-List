![Image of Yaktocat](https://raw.githubusercontent.com/MehmetEminKaymaz/Go-List/master/gopher%20ben.jpeg)

# Go-List
List implementation over slices



### Installation

```
  go get github.com/MehmetEminKaymaz/Go-List
  
```

### Quickstart

```Go
var b []int
	myList:=Lists.From(b)
	myList.Add(1)
	myList.Add(2)
	myList.Add(3)
	fmt.Println(myList.ToSlice())
```

Output
[1 2 3]

```Go
fmt.Println(myList.Where(func(x interface{}) bool {
		return x.(int)<3
	}).ToSlice())

	fmt.Println(myList.Skip(1).Take(2).ToSlice())
```


```Go
var a []Human
	myList:=Lists.From(a)
	myList.Add(Human{Name:"Mehmet",Age:23})
	myList.Add(Human{Name:"Samet",Age:30})
	myList.Reverse()
	fmt.Println(myList.ToSlice())
```
Output : [{Samet 30} {Mehmet 23}]

```Go
myList.Foreach(func(x interface{}) (y interface{}) {
		a:=(x.(Human).Age)+1

		y=Human{Age:a,Name:x.(Human).Name}
		return
	})
	fmt.Println(myList.ToSlice())
  
```

Output : [{Mehmet 24} {Samet 31}]
  

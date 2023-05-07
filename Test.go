package main

import "fmt"

type Order int

const (
	HitsAsc Order = iota
	HitsDesc
	CreatedAtAsc
	CreatedAtDesc
)

func (o Order) String() string {
	switch o {
	case HitsAsc:
		return "hits asc"
	case HitsDesc:
		return "hits desc"
	case CreatedAtAsc:
		return "created_at asc"
	case CreatedAtDesc:
		return "created_at desc"
	default:
		panic(fmt.Sprintf("unsupported order: %d", o))
	}
}

func main() {
	fmt.Println(HitsAsc)
	fmt.Println(HitsDesc)
	fmt.Println(CreatedAtAsc)
	fmt.Println(CreatedAtDesc)
	num := 1
	fmt.Printf("Order is %v", Order(num))

	fmt.Printf("Type is :%T", HitsDesc)
}

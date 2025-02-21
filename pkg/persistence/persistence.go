package persistence

type Order string

func (o Order) String() string {
	return string(o)
}

const (
	OrderASC  = Order("asc")
	OrderDESC = Order("desc")
)

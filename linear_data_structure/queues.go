// Package linear_data_structure
// Time    : 2021/5/11 9:42 上午
// Author  : xushiyin
// contact : yuqingxushiyin@gmail.com
package linear_data_structure

// Order class
type Order struct {
	Priority     int
	Quantity     int
	Product      string
	CustomerName string
}

type Queue []*Order

// New method initializes with Order with priority, quantity, product, customerName
func (order *Order) New(priority int, quantity int, product string,
	customerName string) {
	order.Priority = priority
	order.Quantity = quantity
	order.Product = product
	order.CustomerName = customerName
}

// Add method adds the order to the queue
func (queue *Queue) Add(order *Order) {
	if len(*queue) == 0 {
		*queue = append(*queue, order)
	} else {
		var appended bool
		appended = false
		var i int
		var addedOrder *Order
		for i, addedOrder = range *queue {
			if order.Priority > addedOrder.Priority {
				*queue = append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				appended = true
				break
			}
		}
		if !appended {
			*queue = append(*queue, order)
		}
	}
}

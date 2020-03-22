package problem0729

// Node represents a event.
//
type Node struct {
	start, end  int
	left, right *Node
}

func (n *Node) insert(newNode *Node) bool {
	if newNode.start >= n.end {
		if n.right == nil {
			n.right = newNode
			return true
		}
		return n.right.insert(newNode)
	} else if newNode.end <= n.start {
		if n.left == nil {
			n.left = newNode
			return true
		}
		return n.left.insert(newNode)
	} else {
		return false
	}
}

// MyCalendar is a calendar object.
//
type MyCalendar struct {
	root *Node
}

// Constructor constructs a new calendar.
//
func Constructor() MyCalendar {
	var carlendar = MyCalendar{
		root: nil,
	}
	return carlendar
}

// Book a new event.
func (n *MyCalendar) Book(start int, end int) bool {
	node := Node{
		start: start,
		end:   end,
	}
	if n.root == nil {
		n.root = &node
		return true
	}
	return n.root.insert(&node)
}

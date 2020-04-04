/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

 type Node struct {
	 Val int
	 Next *Node
	 Random *Node
 }

 func List(head *Node) *Node {
    if head == nil {
		return nil
	}
	// Create a map to store the node.
	m := make(map[*Node]*Node);
	var cur *Node
	// Loop 1: copy the original Node Values.
	cur = head
	for cur != nil {
		node := Node {
			Val: cur.Val,
		}
		m[cur] := &node
		cur = cur.Next
	}

	// Loop 2
	cur = head 
	for cur != nil {
		m[cur].Next = m[cur.Next]
		m[cur].Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]

}
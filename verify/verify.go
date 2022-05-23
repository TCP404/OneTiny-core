package verify

type verifyNode struct {
	Handler
	Next *verifyNode
}

type verifyChain struct {
	head *verifyNode
	tail *verifyNode
}

func (c *verifyChain) AddToHead(h Handler) Verifier {
	node := &verifyNode{
		Handler: h,
		Next:    c.head,
	}
	if c.head == nil {
		c.head, c.tail = node, node
		return c
	}
	c.head = node
	return c
}

func (c *verifyChain) AddToTail(h Handler) Verifier {
	node := &verifyNode{
		Handler: h,
		Next:    nil,
	}
	if c.head == nil || c.tail == nil {
		c.head, c.tail = node, node
		return c
	}
	c.tail.Next = node
	return c
}

func (c *verifyChain) Iterator() error {
	curr := c.head
	for curr != nil {
		if err := curr.Handle(); err != nil {
			return err
		}
		curr = curr.Next
	}
	return nil
}

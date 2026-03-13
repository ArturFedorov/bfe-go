package flatten_nested_list_iterator

type NestedInteger struct {
	val  *int
	list []*NestedInteger
}

func (n *NestedInteger) IsInteger() bool { return n.val != nil }
func (n *NestedInteger) GetInteger() int {
	if n.val != nil {
		return *n.val
	}
	return 0
}
func (n *NestedInteger) GetList() []*NestedInteger { return n.list }

type NestedIterator struct{}

func Constructor(nestedList []*NestedInteger) *NestedIterator { return nil }
func (it *NestedIterator) Next() int                          { return 0 }
func (it *NestedIterator) HasNext() bool                      { return false }

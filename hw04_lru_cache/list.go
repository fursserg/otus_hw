package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
	key   Key
}

type list struct {
	length int
	first  *ListItem
	last   *ListItem
}

func (l *list) Len() int {
	return l.length
}

func (l *list) Front() *ListItem {
	return l.first
}

func (l *list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(value interface{}) *ListItem {
	newFirst := &ListItem{
		Value: value,
		Prev:  nil,
		Next:  l.first,
	}

	l.init(newFirst)

	if l.Len() > 0 {
		l.first.Prev = newFirst
	}

	l.first = newFirst

	l.length++

	return newFirst
}

func (l *list) PushBack(value interface{}) *ListItem {
	newLast := &ListItem{
		Value: value,
		Prev:  l.last,
		Next:  nil,
	}

	l.init(newLast)

	if l.Len() > 0 {
		l.last.Next = newLast
	}

	l.last = newLast

	l.length++

	return newLast
}

func (l *list) Remove(item *ListItem) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		l.first = item.Next
	}

	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		l.last = item.Prev
	}

	l.length--
}

func (l *list) MoveToFront(item *ListItem) {
	if l.first == item {
		return
	}

	item.Prev.Next = item.Next

	if item != l.last {
		item.Next.Prev = item.Prev
	} else {
		l.last = item.Prev
	}

	item.Next = l.first
	item.Prev = nil

	l.first.Prev = item
	l.first = item
}

func (l *list) init(v *ListItem) {
	if l.Len() == 0 {
		l.first = v
		l.last = v
	}
}

func (li *ListItem) GetKey() Key {
	return li.key
}

func (li *ListItem) SetKey(key Key) {
	li.key = key
}

func (li *ListItem) GetValue() interface{} {
	return li.Value
}

func NewList() List {
	return &list{}
}

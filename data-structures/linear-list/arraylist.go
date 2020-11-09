package linear_list

import "errors"

var initcapacity = 10
var defaultCapacity = 10
// 1. empty linear-list
type ArrayList struct {
	size     int
	capacity int
	element  []interface{}
}

// 自定义接口实现
//type Array interface {
//	add(index int, ele interface{})
//}


func emptylist() *ArrayList{
	return &ArrayList{}
}
// 2. get by index
func (l *ArrayList) get(index int) (interface{}, error) {
	err := l.checkIndex(index)
	if err != nil {
		return nil, err
	}
	return l.get(index)
}

func (l *ArrayList) checkIndex(index int) error {
	if index < 0 || index >= l.size {
		return errors.New("out of index")
	}
	return nil
}

// 3. find By first
func (l *ArrayList) findFirst(obj interface{}) (int, error) {
	if obj == nil {
		return -1, errors.New("element cannot be null")
	}
	for pos, ele := range l.element {
		if ele == obj {
			return pos, nil
		}
	}
	return -1, nil
}

func (l *ArrayList) resize() {
	if l.size == 0 {
		l.capacity = defaultCapacity
	} else {
		l.capacity = l.capacity << 1
	}

	new := make([]interface{}, l.size)
	copy(new, l.element)
	//
	l.element = new
}

func (l *ArrayList) add(index int, ele interface{}) {
	if index == l.size {
		l.resize()
	}
	l.element[index] = ele
	l.size++
}

func (l *ArrayList) remove(index int) {
	if index != 0 {
		copy(l.element[index:], l.element[index+1:l.size])
		l.element[l.size-1] = nil
		l.size--
	}

}

func (l *ArrayList) len() int {
	return l.size
}

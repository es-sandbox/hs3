package main

import "container/list"

type threadSafeList struct {
	list *list.List
}

func newThreadSafeList() *threadSafeList {
	return &threadSafeList{
		list: &list.List{},
	}
}

func (tsl *threadSafeList) pushBack(rawCommand string) {
	tsl.list.PushBack(rawCommand)
}

func (tsl *threadSafeList) popFront() string {
	elem := tsl.list.Front()
	if elem == nil {
		return ""
	}

	value, ok := elem.Value.(string)
	if !ok {
		return ""
	}

	tsl.list.Remove(elem)

	return value
}
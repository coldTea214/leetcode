package main

import "fmt"

type LRUCache struct {
	capacity int
	data     *DoublyLinkedList
	dataLoc  map[int]*DoublyLinkedListNode
}

type DoublyLinkedList struct {
	root *DoublyLinkedListNode
}

type DoublyLinkedListNode struct {
	key, value int
	prev, next *DoublyLinkedListNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     NewDoublyLinkedList(),
		dataLoc:  make(map[int]*DoublyLinkedListNode),
	}
}

func (lru *LRUCache) Get(key int) int {
	node, ok := lru.dataLoc[key]
	if !ok {
		return -1
	}
	lru.data.MoveToFront(node)
	return node.value
}

func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.dataLoc[key]; ok {
		lru.data.MoveToFront(node)
		node.value = value
	} else {
		if len(lru.dataLoc) == lru.capacity {
			back := lru.data.Back()
			lru.data.Remove(back)
			delete(lru.dataLoc, back.key)
		}
		node := &DoublyLinkedListNode{key: key, value: value}
		lru.data.PushFront(node)
		lru.dataLoc[key] = node
	}
}

func NewDoublyLinkedList() *DoublyLinkedList {
	list := &DoublyLinkedList{
		root: &DoublyLinkedListNode{},
	}
	list.root.next = list.root
	list.root.prev = list.root
	return list
}

func (list *DoublyLinkedList) MoveToFront(node *DoublyLinkedListNode) {
	if list.root.next == node {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = list.root
	node.next = list.root.next
	node.prev.next = node
	node.next.prev = node
}

func (list *DoublyLinkedList) Back() *DoublyLinkedListNode {
	if list.root.prev == list.root {
		return nil
	}
	return list.root.prev
}

func (list *DoublyLinkedList) Remove(node *DoublyLinkedListNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (list *DoublyLinkedList) PushFront(node *DoublyLinkedListNode) {
	node.prev = list.root
	node.next = list.root.next
	node.prev.next = node
	node.next.prev = node
}

func main() {
	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)
	fmt.Println(lru.Get(1)) // 返回1
	lru.Put(3, 3)
	fmt.Println(lru.Get(2)) // 返回-1
	lru.Put(4, 4)
	fmt.Println(lru.Get(1)) // 返回-1
	fmt.Println(lru.Get(3)) // 返回3
	fmt.Println(lru.Get(4)) // 返回4
}

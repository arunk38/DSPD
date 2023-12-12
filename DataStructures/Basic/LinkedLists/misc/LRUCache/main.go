package main

import "fmt"

type Node struct {
	Prev, Next *Node
	Key, Value int
}

type LRUCache struct {
	Head       *Node
	Tail       *Node
	mapPointer map[int]*Node
	cacheCap   int
}

func newNode(key, val int) *Node {
	return &Node{
		Key:   key,
		Value: val,
	}
}

func newLRUCache(head, tail *Node, cap int) LRUCache {
	return LRUCache{
		Head:       head,
		Tail:       tail,
		mapPointer: make(map[int]*Node),
		cacheCap:   cap,
	}
}

func Constructor(capacity int) LRUCache {
	head, tail := newNode(0, 0), newNode(0, 0)

	head.Next = tail
	tail.Prev = head
	return newLRUCache(head, tail, capacity)
}

func (this *LRUCache) remove(node *Node) {
	delete(this.mapPointer, node.Key)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) insert(node *Node) {
	this.mapPointer[node.Key] = node
	next := this.Head.Next
	this.Head.Next = node
	node.Prev = this.Head
	next.Prev = node
	node.Next = next
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.mapPointer[key]; ok {
		this.remove(n)
		this.insert(n)
		return n.Value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.mapPointer[key]; ok {
		this.remove(this.mapPointer[key])
	}

	if len(this.mapPointer) == this.cacheCap {
		this.remove(this.Tail.Prev)
	}

	this.insert(newNode(key, value))
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)           // cache is {1=1}
	lRUCache.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // return 1
	lRUCache.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	fmt.Println(lRUCache.Get(3)) // return 3
	fmt.Println(lRUCache.Get(4)) // return 4
}

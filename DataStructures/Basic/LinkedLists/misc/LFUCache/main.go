package main

import "fmt"

/*
 * Implement the LFUCache class:
 *
 *  - LFUCache(int capacity) Initializes the object with the capacity of the data structure.
 *  - int get(int key) Gets the value of the key if the key exists in the cache. Otherwise, returns -1.
 *  - void put(int key, int value) Update the value of the key if present, or inserts the key if not already present.
 *    When the cache reaches its capacity, it should invalidate and remove the least frequently used key before inserting a new item.
 *	  For this problem, when there is a tie (i.e., two or more keys with the same frequency), the least recently used key would be invalidated.
 *
 * To determine the least frequently used key, a use counter is maintained for each key in the cache. The key with the smallest use counter
 * is the least frequently used key.
 * When a key is first inserted into the cache, its use counter is set to 1 (due to the put operation). The use counter for a key in the cache
 * is incremented either a get or put operation is called on it.
 *
 * The functions get and put must each run in O(1) average time complexity.
 *
 * link: https://leetcode.com/problems/lfu-cache/
 */

// FreqNode represents a node in the frequency list
type FreqNode struct {
	freq       int
	nodes      []*Node
	next, prev *FreqNode
}

// Node represents a node in LFU cache.
type Node struct {
	key, value int
	freqNode   *FreqNode
}

// LFUCache represents the LFU cache structure
type LFUCache struct {
	capacity int
	cache    map[int]*Node
	freq     map[int]*FreqNode
	minFreq  int
}

// Constructor initializes the LFU cache with given capacity
func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		freq:     make(map[int]*FreqNode),
	}
}

func (this *LFUCache) updateFrequency(node *Node) {
	freqNode := node.freqNode

	// Remove the current node from its freqlist
	for i := 0; i <= (len(freqNode.nodes) - 1); i++ {
		if freqNode.nodes[i].key == node.key {
			freqNode.nodes = append(freqNode.nodes[:i], freqNode.nodes[i+1:]...)
		}
	}

	// remove the freqNode if there are not more nodes in it
	if len(freqNode.nodes) == 0 {
		this.removeFreqNode(freqNode)
		if this.minFreq == freqNode.freq {
			this.minFreq++
		}
	}

	// add to next freqNode list, or create a new list
	newNodeFreq := freqNode.freq + 1
	if nextFreqNode, ok := this.freq[newNodeFreq]; ok {
		nextFreqNode.nodes = append(nextFreqNode.nodes, node)
		node.freqNode = nextFreqNode
	} else {
		nextFreqNode := &FreqNode{freq: newNodeFreq, nodes: []*Node{node}}
		this.freq[newNodeFreq] = nextFreqNode

		if nextFreqNode.next == nil {
			nextFreqNode.next = this.freq[this.minFreq].next
			this.freq[this.minFreq].next = nextFreqNode
		} else {
			nextFreqNode.next.prev = nextFreqNode
		}

		node.freqNode = nextFreqNode
	}
}

func (this *LFUCache) addToFrequencyList(node *Node) {
	freq := 1
	if node.freqNode != nil {
		freq = node.freqNode.freq + 1
	}

	if freqNode, ok := this.freq[freq]; ok {
		freqNode.nodes = append(freqNode.nodes, node)
		node.freqNode = freqNode
	} else {
		freqNode := &FreqNode{freq: freq, nodes: []*Node{node}}
		this.freq[freq] = freqNode

		if prevFreqNode, ok := this.freq[freq-1]; ok {
			freqNode.prev = prevFreqNode
			freqNode.next = prevFreqNode.next
			prevFreqNode.next = freqNode
		} else {
			freqNode.next = this.freq[freq].next
			this.freq[freq].next = freqNode
		}

		if freqNode.next != nil {
			freqNode.next.prev = freqNode
		}

		node.freqNode = freqNode
		this.minFreq = freq
	}
}

func (this *LFUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// key is present, increment its freq by 1
		this.updateFrequency(node)
		return node.value
	}

	return -1
}

func (this *LFUCache) removeFreqNode(freqNode *FreqNode) {
	if freqNode.prev != nil {
		freqNode.prev.next = freqNode.next
	} else {
		this.freq[this.minFreq].next = freqNode.next
	}

	if freqNode.next != nil {
		freqNode.next.prev = freqNode.prev
	}

	delete(this.freq, freqNode.freq)
}

func (this *LFUCache) evict() {
	freqNode := this.freq[this.minFreq]
	// 0 index node is the LRU key, evict to make room for next node
	evitNode := freqNode.nodes[0]

	// iterate the items in the freq node, and delete them from the cache
	delete(this.cache, evitNode.key)

	this.removeNodeFromFreqNode(evitNode)
}

func (this *LFUCache) removeNodeFromFreqNode(node *Node) {
	freqNode := node.freqNode

	freqNode.nodes = freqNode.nodes[1:]

	if len(freqNode.nodes) == 0 {
		this.removeFreqNode(freqNode)
		if this.minFreq == freqNode.freq {
			this.minFreq++
		}
	}

}

func (this *LFUCache) Put(key, value int) {
	if this.capacity == 0 {
		return
	}

	if node, ok := this.cache[key]; ok {
		node.value = value
		this.updateFrequency(node)
	} else {
		if len(this.cache) >= this.capacity {
			this.evict()
		}

		newNode := &Node{key: key, value: value}
		this.cache[key] = newNode
		this.addToFrequencyList(newNode)
	}
}

func main() {
	lfucache := Constructor(2)
	lfucache.Put(3, 1)           // cache [3,_], cnt(3) = 1
	lfucache.Put(2, 1)           // cache [3,2], cnt{3} = 1, cnt(2) = 1
	lfucache.Put(2, 2)           // cache [3,2], cnt{3} = 1, cnt(2) = 2
	lfucache.Put(4, 4)           // cache [4,2], cnt{4} = 1, cnt(2) = 2
	fmt.Println(lfucache.Get(2)) // return 2

}

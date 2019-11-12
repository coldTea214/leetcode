import "container/list"

type LRUCache struct {
	cap     int
	data    *list.List
	dataLoc map[int]*list.Element
}

type ElementData struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:     capacity,
		data:    list.New(),
		dataLoc: make(map[int]*list.Element, capacity),
	}
}

func (cache *LRUCache) Get(key int) int {
	if dataPtr, ok := cache.dataLoc[key]; ok {
		val := dataPtr.Value.(*ElementData).value
		cache.data.MoveToFront(dataPtr)
		return val
	}
	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if dataPtr, ok := cache.dataLoc[key]; ok {
		cache.data.MoveToFront(dataPtr)
		dataPtr.Value.(*ElementData).value = value
	} else {
		if cache.data.Len() == cache.cap {
			back := cache.data.Remove(cache.data.Back())
			delete(cache.dataLoc, back.(*ElementData).key)
		}

		data := &ElementData{key: key, value: value}
		dataPtr := cache.data.PushFront(data)
		cache.dataLoc[key] = dataPtr
	}
}

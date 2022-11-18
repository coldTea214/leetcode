// trie
type Trie struct {
	isWord bool
	// 这里是通用前缀树写法，本题限定了输入，用 26 位数组更好
	// 具体可参考 212
	children map[rune]*Trie
}

func Constructor() Trie {
	return Trie{isWord: false, children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, ch := range word {
		if child, ok := cur.children[ch]; ok {
			cur = child
		} else {
			newChild := &Trie{children: make(map[rune]*Trie)}
			cur.children[ch] = newChild
			cur = newChild
		}
	}
	cur.isWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, ch := range word {
		child, ok := cur.children[ch]
		if !ok {
			return false
		}
		cur = child
	}
	return cur.isWord
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, ch := range prefix {
		child, ok := cur.children[ch]
		if !ok {
			return false
		}
		cur = child
	}
	return true
}

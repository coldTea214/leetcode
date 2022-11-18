// trie
type WordDictionary struct {
	children map[rune]*WordDictionary
	isWord   bool
}

func Constructor() WordDictionary {
	return WordDictionary{children: make(map[rune]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, ch := range word {
		if child, ok := cur.children[ch]; ok {
			cur = child
		} else {
			newChild := &WordDictionary{children: make(map[rune]*WordDictionary)}
			cur.children[ch] = newChild
			cur = newChild
		}
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	cur := this
	for i, ch := range word {
		if rune(ch) == '.' {
			isMatched := false
			for _, v := range cur.children {
				if v.Search(word[i+1:]) {
					isMatched = true
				}
			}
			return isMatched
		} else if _, ok := cur.children[rune(ch)]; !ok {
			return false
		}
		cur = cur.children[rune(ch)]
	}
	return len(cur.children) == 0 || cur.isWord
}

type Queue struct {
	words []string
}

func (q *Queue) isEmpty() bool {
	return len(q.words) == 0
}

func (q *Queue) length() int {
	return len(q.words)
}

func (q *Queue) enqueue(word string) {
	q.words = append(q.words, word)
}

func (q *Queue) dequeue() string {
	word := q.words[0]
	q.words = q.words[1:]
	return word
}

func ladderLength(beginWord string, endWord string, words []string) int {
	wordExist := make(map[string]bool, len(words))
	for _, word := range words {
		wordExist[word] = true
	}
	if !wordExist[endWord] {
		return 0
	}

	var queue Queue
	queue.enqueue(beginWord)
	wordEnqueued := map[string]bool{beginWord: true}
	length := 2
	for !queue.isEmpty() {
		qLen := queue.length()
		for i := 0; i < qLen; i++ {
			front := queue.dequeue()
			for i := 0; i < len(front); i++ {
				bytes := []byte(front)
				for j := 0; j < 26; j++ {
					bytes[i] = 'a' + byte(j)
					word := string(bytes)
					if word == endWord {
						return length
					}
					if wordExist[word] && !wordEnqueued[word] {
						queue.enqueue(word)
						wordEnqueued[word] = true
					}
				}
			}
		}
		length++
	}
	return 0
}

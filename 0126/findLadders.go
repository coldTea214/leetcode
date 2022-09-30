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

func findLadders(beginWord string, endWord string, words []string) [][]string {
	wordExist := make(map[string]bool, len(words))
	for _, word := range words {
		wordExist[word] = true
	}

	nextWord := make(map[string][]string)
	ladderLength := func() int {
		if !wordExist[endWord] {
			return 0
		}

		var queue Queue
		queue.enqueue(beginWord)
		wordEnqueued := map[string]bool{beginWord: true}
		length := 2
		quit := false
		for !queue.isEmpty() {
			qLen := queue.length()
			// wordsInLevelN 记录出现在第 N 层的 word
			// 既要处理好：abc -> abe,adc -> ade 这种同长度同终点两路径的情况
			// 也要防止 ade 在第 3 层入队两次，会超时
			wordsInLevelN := make(map[string]bool)
			for i := 0; i < qLen; i++ {
				front := queue.dequeue()
				for i := 0; i < len(front); i++ {
					bytes := []byte(front)
					for j := 0; j < 26; j++ {
						bytes[i] = 'a' + byte(j)
						word := string(bytes)
						if word == endWord {
							quit = true
						}
						if wordExist[word] {
							if word != front && !wordEnqueued[word] {
								nextWord[front] = append(nextWord[front], word)
								wordsInLevelN[word] = true
							}
						}
					}
				}
			}
			for word := range wordsInLevelN {
				queue.enqueue(word)
				wordEnqueued[word] = true
			}

			if quit {
				return length
			}
			length++
		}
		return 0
	}
	length := ladderLength()
	if length == 0 {
		return [][]string{}
	}

	var ladders [][]string
	var doFindLadders func(string, []string)
	doFindLadders = func(beginWord string, ladder []string) {
		if len(ladder) == length {
			if ladder[len(ladder)-1] == endWord {
				deepcopy := make([]string, len(ladder))
				copy(deepcopy, ladder)
				ladders = append(ladders, deepcopy)
			}
			return
		}

		for _, word := range nextWord[beginWord] {
			doFindLadders(word, append(ladder, word))
		}
	}
	doFindLadders(beginWord, []string{beginWord})
	return ladders
}

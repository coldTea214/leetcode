// trie
func findWords(board [][]byte, words []string) []string {
	m := len(board)
	if m == 0 {
		return nil
	}
	n := len(board[0])
	if n == 0 {
		return nil
	}

	var res []string
	trie := buildTrie(words)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, trie, &res)
		}
	}
	return res
}

type TrieNode struct {
	children [26]*TrieNode
	// 非标准 trie
	word string
}

func buildTrie(words []string) *TrieNode {
	root := &TrieNode{}
	for _, word := range words {
		cur := root
		for _, c := range word {
			cidx := int(c - 'a')
			if cur.children[cidx] == nil {
				cur.children[cidx] = &TrieNode{}
			}
			cur = cur.children[cidx]
		}
		cur.word = word
	}
	return root
}

func dfs(board [][]byte, i int, j int, trie *TrieNode, res *[]string) {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || trie.children[int(c-'a')] == nil {
		return
	}

	trie = trie.children[int(c-'a')]
	if trie.word != "" {
		*res = append(*res, trie.word)
		trie.word = ""
	}

	board[i][j] = '#'
	dfs(board, i-1, j, trie, res)
	dfs(board, i+1, j, trie, res)
	dfs(board, i, j-1, trie, res)
	dfs(board, i, j+1, trie, res)
	board[i][j] = c
}

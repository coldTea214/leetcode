import "strconv"

type ValidWordAbbr struct {
	abbrExist map[string]bool
	abbr2Word map[string]string
}

func Constructor(dictionary []string) ValidWordAbbr {
	abbrExist := make(map[string]bool)
	abbr2Word := make(map[string]string)
	for _, word := range dictionary {
		abbr := getAbbr(word)
		if !abbrExist[abbr] {
			abbrExist[abbr] = true
			abbr2Word[abbr] = word
		} else if word != abbr2Word[abbr] {
			delete(abbr2Word, abbr)
		}
	}
	return ValidWordAbbr{abbrExist, abbr2Word}
}

func (vwa *ValidWordAbbr) IsUnique(word string) bool {
	abbr := getAbbr(word)
	if !vwa.abbrExist[abbr] {
		return true
	}
	return word == vwa.abbr2Word[abbr]
}

func getAbbr(word string) string {
	if len(word) <= 2 {
		return word
	} else {
		return string(word[0]) + strconv.Itoa(len(word)-2) + string(word[len(word)-1])
	}
}


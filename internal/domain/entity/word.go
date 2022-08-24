package entity

import "errors"

// Word 単語を表すEntity
type Word struct {
	ID       string
	WordName string
	Meaning  string
}

type UserWords map[string][]*Word

// isWordValid 単語が入力されているかを判定する
func (w *Word) isWordValid() bool {
	return len(w.WordName) != 0
}

// NewWord 新たな単語を生成する
func NewWord(
	ID string,
	WordName string,
	Meaning string,
) (*Word, error) {
	word := &Word{
		ID:       ID,
		WordName: WordName,
		Meaning:  Meaning,
	}

	ok := word.isWordValid()
	if !ok {
		return nil, errors.New("the number of answer is not existed in questions")
	}

	return word, nil
}

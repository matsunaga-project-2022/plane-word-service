package entity

import "errors"

// Word 単語を表すEntity
type Word struct {
	ID       uint32
	UserID   uint32
	WordName string
	Meaning  string
}

// isUserValid ユーザが存在するかを判定する
func (w *Word) isUserValid(userID uint32) bool {
	return w.UserID == userID
}

// isWordValid 単語が入力されているかを判定する
func (w *Word) isWordValid() bool {
	return len(w.WordName) != 0
}

// NewWord 新たな単語を生成する
func NewWord(
	ID uint32,
	UserID uint32,
	WordName string,
	Meaning string,
) (*Word, error) {
	word := &Word{
		ID:       ID,
		UserID:   UserID,
		WordName: WordName,
		Meaning:  Meaning,
	}

	ok := word.isUserValid(4)
	if !ok {
		return nil, errors.New("question choice length is not four")
	}

	ok = word.isWordValid()
	if !ok {
		return nil, errors.New("the number of answer is not existed in questions")
	}

	return word, nil
}

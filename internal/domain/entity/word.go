package entity

import "errors"

// Word 単語を表すEntity
type Word struct {
	ID       uint32
	UserID   uint32
	WordName string
	Meaning  string
}

// isWordValid 単語が入力されているか判定？
func (w *Word) isWordValid() bool {
	return len(w.WordName) != 0
}

// isUserValid ユーザが存在するか判定？
func (w *Word) isUserValid(userID uint32) bool {
	return w.UserID == userID
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

	ok := word.isWordValid()
	if !ok {
		return nil, errors.New("question choice length is not four")
	}

	ok = word.isUserValid(1)
	if !ok {
		return nil, errors.New("the number of answer is not existed in questions")
	}

	return word, nil
}

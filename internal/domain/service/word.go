package service

import (
	"context"
	"github.com/matsunaga-project-2022/shanks/internal/domain/entity"
	"github.com/matsunaga-project-2022/shanks/internal/domain/repository"
	"github.com/matsunaga-project-2022/shanks/internal/pkg/uuid"
)

type Word struct {
	repository repository.Word
}

func NewWord(repository repository.Word) *Word {
	return &Word{
		repository: repository,
	}
}

func (w Word) CreateWord(ctx context.Context, wordName string, meaning string, userID string) error {
	id := uuid.New()
	word := &entity.Word{
		ID:       id,
		WordName: wordName,
		Meaning:  meaning,
	}

	if err := w.repository.CreateWord(ctx, word, userID); err != nil {
		return err
	}
	return nil
}

func (w Word) UpdateWord(ctx context.Context, id string, meaning string) error {
	word, err := w.repository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	word.Meaning = meaning

	if err = w.repository.UpdateWord(ctx, word); err != nil {
		return err
	}
	return nil
}

func (w Word) DeleteWord(ctx context.Context, id string) error {
	if err := w.repository.DeleteWord(ctx, id); err != nil {
		return err
	}
	return nil
}

func (w Word) ListWordByUserID(ctx context.Context, userID string) ([]*entity.Word, error) {
	word, err := w.repository.ListWordByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return word, nil
}

func (w Word) GetWordByID(ctx context.Context, id string) (*entity.Word, error) {
	word, err := w.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return word, nil
}

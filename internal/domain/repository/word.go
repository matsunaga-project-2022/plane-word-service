package repository

import (
	"context"
	"github.com/matsunaga-project-2022/shanks/internal/domain/entity"
)

type Word interface {
	CreateWord(ctx context.Context, word *entity.Word, userID string) error
	UpdateWord(ctx context.Context, word *entity.Word) error
	DeleteWord(ctx context.Context, id string) error
	ListWordByUserID(ctx context.Context, userID string) ([]*entity.Word, error)
}

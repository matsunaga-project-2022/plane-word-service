package repository

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/matsunaga-project-2022/shanks/internal/domain/entity"
	"log"
)

type WordRepositoryMysqlImp struct {
	db *sqlx.DB
}

func NewWordRepositoryMysqlImp(db *sqlx.DB) Word {
	return &WordRepositoryMysqlImp{
		db: db,
	}
}

func (wr *WordRepositoryMysqlImp) ListWordByUserID(ctx context.Context, userID string) ([]*entity.Word, error) {
	var words []*entity.Word
	log.Println(userID)
	err := wr.db.SelectContext(ctx, &words, `
		SELECT id, word_name, meaning FROM words
		WHERE user_id = ?
	`, userID)

	if err == sql.ErrNoRows {
		log.Println("Data not found.")
		return nil, nil
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return words, nil
}

func (wr *WordRepositoryMysqlImp) GetByID(ctx context.Context, id string) (*entity.Word, error) {
	word := &entity.Word{}
	log.Println(id)
	err := wr.db.GetContext(ctx, word, `
		SELECT id, word_name, meaning FROM words
		WHERE id = ?
	`, id)

	if err == sql.ErrNoRows {
		log.Println("Data not found.")
		return nil, nil
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return word, nil
}

func (wr *WordRepositoryMysqlImp) CreateWord(ctx context.Context, word *entity.Word, userID string) (err error) {
	tx := wr.db.MustBeginTx(ctx, nil)
	_, err = tx.ExecContext(ctx, `
		INSERT INTO words (id, user_id, word_name, meaning)
		VALUES (?,?,?,?)
	`, word.ID, userID, word.WordName, word.Meaning)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}
	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (wr *WordRepositoryMysqlImp) UpdateWord(ctx context.Context, word *entity.Word) (err error) {
	tx := wr.db.MustBeginTx(ctx, nil)
	_, err = tx.NamedExecContext(ctx, `
		UPDATE words SET meaning = :meaning
		WHERE id = :id
	`, word)

	if err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

func (wr *WordRepositoryMysqlImp) DeleteWord(ctx context.Context, id string) (err error) {
	tx := wr.db.MustBeginTx(ctx, nil)
	result := tx.MustExecContext(ctx, `
		DELETE FROM words
		WHERE id = ?
	`, id)

	if _, err = result.LastInsertId(); err != nil {
		tx.Rollback()
		return
	}

	if err = tx.Commit(); err != nil {
		return
	}

	return
}

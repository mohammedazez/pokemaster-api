package transactor

import (
	"context"

	"gorm.io/gorm"
)

type (
	Transactor struct {
		db *gorm.DB
	}

	TxKey struct{}
)

func NewTransactor(db *gorm.DB) *Transactor {
	return &Transactor{
		db,
	}
}

func (t *Transactor) WithinTransaction(tFunc func(txCtx context.Context) error) error {
	ctx := context.Background()
	tx := t.db.Begin()

	// run callback
	err := tFunc(injectTx(ctx, tx))
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func injectTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, TxKey{}, tx)
}

func ExtractTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(TxKey{}).(*gorm.DB); ok {
		return tx
	}
	return nil
}

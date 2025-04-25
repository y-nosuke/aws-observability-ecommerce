package repository

import (
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository は全てのリポジトリの基本インターフェース
type Repository interface {
	WithTx(*sql.Tx) Repository
	DB() boil.ContextExecutor
}

// RepositoryBase は基本リポジトリの共通実装
type RepositoryBase struct {
	executor boil.ContextExecutor
}

// NewRepositoryBase は新しいRepositoryBaseを作成します
func NewRepositoryBase(executor boil.ContextExecutor) RepositoryBase {
	return RepositoryBase{
		executor: executor,
	}
}

// WithTx はトランザクションを設定したリポジトリを返します
func (r RepositoryBase) WithTx(tx *sql.Tx) RepositoryBase {
	r.executor = tx
	return r
}

// DB はデータベースエグゼキュータを返します
func (r RepositoryBase) DB() boil.ContextExecutor {
	return r.executor
}

// RunInTransaction はトランザクション内で関数を実行します
func RunInTransaction(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(); err != nil {
				panic(err)
			}
			panic(p)
		}
	}()

	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	return tx.Commit()
}

package repository

import (
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Repository はリポジトリの基本インターフェース
type Repository interface {
	DB() boil.ContextExecutor
	WithTx(tx *sql.Tx) Repository
}

// RepositoryBase はリポジトリの基本実装
type RepositoryBase struct {
	db boil.ContextExecutor
}

// NewRepositoryBase は新しいRepositoryBaseを作成します
func NewRepositoryBase(executor boil.ContextExecutor) RepositoryBase {
	return RepositoryBase{
		db: executor,
	}
}

// DB はデータベース接続を返します
func (r RepositoryBase) DB() boil.ContextExecutor {
	return r.db
}

// WithTx はトランザクションを設定したRepositoryBaseを返します
func (r RepositoryBase) WithTx(tx *sql.Tx) RepositoryBase {
	return RepositoryBase{
		db: tx,
	}
}

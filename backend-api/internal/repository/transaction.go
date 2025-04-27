package repository

import (
	"context"
	"database/sql"
	"fmt"
)

// RunInTransaction はトランザクション内で関数を実行します
func RunInTransaction(ctx context.Context, db *sql.DB, fn func(*sql.Tx) error) error {
	// トランザクションを開始
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// 関数の実行
	if err := fn(tx); err != nil {
		// エラーの場合はロールバック
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("failed to rollback transaction: %w (original error: %v)", rbErr, err)
		}
		return err
	}

	// トランザクションのコミット
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

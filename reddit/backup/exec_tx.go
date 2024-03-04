package db

import (
	"context"
	"fmt"
)

// ExecTx executes a function within a database
func (repo *SQLRepo) execTx(ctx context.Context, fn func(*Queries) error) error {

	// txOptions := pgx.TxOptions{
	// 	IsoLevel:   pgx.Serializable,
	// 	AccessMode: pgx.ReadWrite,
	// }

	fmt.Printf("\n repo execTx %v\n", repo)
	fmt.Printf("\n context execTx %v\n", ctx)
	// fmt.Printf("\n txOptions execTx %v\n", txOptions)

	tx, err := repo.connPool.Begin(ctx)
	fmt.Printf("\n tx execTx %v\n", tx)
	fmt.Printf("\n repo execTx %v\n", repo)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}

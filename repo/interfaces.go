package repo

import (
	"context"
)

type (
	App interface {
		Migrate()
		Debug()

		// transaction
		Begin() App
		Commit() error
		Rollback() error
	}
)

type txKey struct{}

// injectTx injects transaction to context
func InjectTx(ctx context.Context, tx App) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

// extractTx extracts transaction from context
func ExtractTx(ctx context.Context) App {
	if tx, ok := ctx.Value(txKey{}).(App); ok {
		return tx
	}
	return nil
}

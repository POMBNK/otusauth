package tx

import (
	"context"

	"github.com/POMBNK/gateway/pkg/client/postgres"
)

type Tx interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) error) error
}

type tx struct {
	t *postgres.Db
}

func (t *tx) WithTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.t.WithTx(ctx, fn)
}

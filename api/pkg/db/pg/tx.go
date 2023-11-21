package pg

import (
	"context"
	"time"

	"github.com/cenkalti/backoff/v4"
	pkgerrors "github.com/pkg/errors"
)

// Tx starts a transaction
func Tx(ctx context.Context, dbconn BeginnerExecutor, callback func(ContextExecutor) error) error {
	return TxWithBackOff(ctx, ExponentialBackOff(3, time.Minute), dbconn, callback)
}

// TxWithBackOff starts a transaction with the provided backoff policy
func TxWithBackOff(ctx context.Context, b backoff.BackOff, dbconn BeginnerExecutor, callback func(ContextExecutor) error) error {
	if b == nil {
		b = &backoff.StopBackOff{}
	}

	tx, err := beginTx(ctx, dbconn, b)
	if err != nil {
		return err
	}

	tx = &instrumentedTx{
		Transactor: tx,
		info:       dbconn.InstanceInfo(),
		ctx:        ctx,
	}

	var committed bool
	defer func() {
		if committed {
			return
		}
		// better to be assured we always clean up even in (unlikely) case where `callback(tx)` invoked `runtime.Goexit`
		// this defer might encounter sql.ErrTxDone if tx was committed; but err is discarded
		_ = tx.Rollback()
	}()

	// errors from here on are code errors and should not retried
	if err = callback(tx); err != nil {
		return err
	}
	if err = tx.Commit(); err != nil {
		return pkgerrors.WithStack(err)
	}

	committed = true
	return nil
}

func beginTx(ctx context.Context, dbconn BeginnerExecutor, b backoff.BackOff) (Transactor, error) {
	var tryCount int
	var tx Transactor
	if err := backoff.Retry(func() error {
		tryCount++
		var err error

		tx, err = dbconn.BeginTx(ctx, nil)

		return pkgerrors.WithStack(err)
	}, backoff.WithContext(b, ctx)); err != nil {
		return nil, err
	}
	return tx, nil
}

package repository

import (
	"context"

	"code.in.spdigital.sg/sp-digital/athena/blob/azstorage"
	"code.in.spdigital.sg/sp-digital/athena/db/pg"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/blobstorage"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/cardpayment"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/egiro"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/flins"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/order"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/outgoingevent"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/paynow"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/recurring"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/sap"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/unidollar"
	"code.in.spdigital.sg/sp-digital/njord/api/internal/repository/utilitiesaccount"
	"github.com/cenkalti/backoff/v4"
)

// Registry interface provide specification for
// available domain objects in repository, Registry acts as
// Object Factory for all domain object instantiation.
type Registry interface {
	// CheckPGConnection will check if the DB connection is alive or not
	CheckPGConnection(context.Context) error
	// GetOrder retrieves order repository
	GetOrder() order.Repository
	// GetPayNow retrieves paynow repository
	GetPayNow() paynow.Repository
	// GetSAP retrieves sap repository
	GetSAP() sap.Repository
	// GetUtilitiesAccount return new object for utilities account domain
	GetUtilitiesAccount() utilitiesaccount.Repository
	// GetCardPayment return new object for adhoc payment and recurring payment
	GetCardPayment() cardpayment.Repository
	// GetFlins return news object for flins
	GetFlins() flins.Repository
	// DoInTx allows to combine multiple repository operations in a single tx
	DoInTx(ctx context.Context, txFunc TxFunc, overrideBackoffPolicy backoff.BackOff) error
	// GetBlobStorage returns Azure Storage Repository
	GetBlobStorage() blobstorage.Repository
	// GetOutgoingEvent returns events repository
	GetOutgoingEvent() outgoingevent.Repository
	// GetRecurring returns recurring repository
	GetRecurring() recurring.Repository
	// GetUniDollarPayment return new object for uni dollar payment
	GetUniDollarPayment() unidollar.Repository
	// GetEGiro returns egiro repository
	GetEGiro() egiro.Repository
}

// TxFunc func to allow multiple repository commits in single transaction.
type TxFunc func(ctx context.Context, repo Registry) error

// New instantiates Registry implementation
func New(pgConn pg.BeginnerExecutor) Registry {
	return &impl{
		pgConn: pgConn,
	}
}

// NewWithBlobStorage instantiates Registry implementation with azblob
func NewWithBlobStorage(pgConn pg.BeginnerExecutor, blobClient *azstorage.Client, storageContainerName string) Registry {
	return &impl{
		pgConn:  pgConn,
		storage: blobstorage.New(blobClient, storageContainerName),
	}
}

type impl struct {
	pgConn pg.BeginnerExecutor // the connection used for all db operations
	tx     pg.ContextExecutor  // the transaction created from pgConn that will be override the use of all pgConn usage
	// NOTE: A registry instance can either have a pgConn OR a tx. Cannot have both
	blobClient *azstorage.Client
	storage    blobstorage.Repository
}

// GetBlobStorage returns the Azure Storage repo
func (i impl) GetBlobStorage() blobstorage.Repository {
	return i.storage
}

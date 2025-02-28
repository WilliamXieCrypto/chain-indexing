package handlers

import (
	applogger "github.com/WilliamXieCrypto/chain-indexing/external/logger"
	"github.com/valyala/fasthttp"

	"github.com/WilliamXieCrypto/chain-indexing/appinterface/projection/view"
	"github.com/WilliamXieCrypto/chain-indexing/appinterface/rdb"
	"github.com/WilliamXieCrypto/chain-indexing/infrastructure/httpapi"
	account_transaction_view "github.com/WilliamXieCrypto/chain-indexing/projection/account_transaction/view"
)

type AccountTransactions struct {
	logger applogger.Logger

	accountTransactionsView *account_transaction_view.AccountTransactions
}

func NewAccountTransactions(logger applogger.Logger, rdbHandle *rdb.Handle) *AccountTransactions {
	return &AccountTransactions{
		logger.WithFields(applogger.LogFields{
			"module": "AccountTransactionsHandler",
		}),

		account_transaction_view.NewAccountTransactions(rdbHandle),
	}
}

func (handler *AccountTransactions) ListByAccount(ctx *fasthttp.RequestCtx) {
	pagination, err := httpapi.ParsePagination(ctx)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	account, accountOk := URLValueGuard(ctx, handler.logger, "account")
	if !accountOk {
		return
	}

	queryArgs := ctx.QueryArgs()

	idOrder := view.ORDER_ASC
	if queryArgs.Has("order") {
		if string(queryArgs.Peek("order")) == "height.desc" {
			idOrder = view.ORDER_DESC
		}
	}
	memo := ""
	if queryArgs.Has("memo") {
		memo = string(queryArgs.Peek("memo"))
	}

	filter := account_transaction_view.AccountTransactionsListFilter{
		Account: account,
		Memo:    memo,
	}

	blocks, paginationResult, err := handler.accountTransactionsView.List(
		filter, account_transaction_view.AccountTransactionsListOrder{Id: idOrder}, pagination,
	)
	if err != nil {
		handler.logger.Errorf("error listing account transactions: %v", err)
		httpapi.InternalServerError(ctx)
		return
	}

	httpapi.SuccessWithPagination(ctx, blocks, paginationResult)
}

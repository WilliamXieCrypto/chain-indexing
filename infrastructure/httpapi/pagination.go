package httpapi

import (
	"strconv"

	"github.com/valyala/fasthttp"

	pagination_interface "github.com/WilliamXieCrypto/chain-indexing/appinterface/pagination"
)

func ParsePagination(ctx *fasthttp.RequestCtx) (*pagination_interface.Pagination, error) {
	var err error

	var pagination pagination_interface.PaginationType
	var page, limit int64

	queryArgs := NewQueryArgs(ctx.QueryArgs())

	pagination = queryArgs.Get("pagination")
	if pagination == "" {
		pagination = pagination_interface.PAGINATION_OFFSET
	}
	if pagination != pagination_interface.PAGINATION_OFFSET {
		return nil, ErrInvalidPagination
	}

	pageQuery := queryArgs.Get("page")
	if pageQuery == "" {
		page = int64(1)
	} else {
		page, err = strconv.ParseInt(pageQuery, 10, 64)
		if err != nil {
			return nil, ErrInvalidPage
		}
		if page == 0 {
			return nil, ErrInvalidPage
		}
	}

	var defaultLimit int64 = int64(20)

	limitQuery := queryArgs.Get("limit")
	if limitQuery == "" {
		limit = defaultLimit
	} else {
		limit, err = strconv.ParseInt(limitQuery, 10, 64)
		if err != nil {
			return nil, ErrInvalidPage
		}
		if limit <= 0 {
			limit = defaultLimit
		}
	}

	return pagination_interface.NewOffsetPagination(page, limit), nil
}

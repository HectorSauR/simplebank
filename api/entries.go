package api

import (
	"database/sql"
	"net/http"

	db "github.com/HectorSauR/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createEntryRequest struct {
	AccountID int64 `uri:"id" binding:"required"`
	Amount    int64 `json:"amount" binding:"required,min:1"`
}

func (server *Server) createEntry(ctx *gin.Context) {
	var req createEntryRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateEntryParams{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	}

	entry, err := server.store.CreateEntry(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

type getAccountEntriesRequest struct {
	AccountID int64 `uri:"id" binding:"required"`
	EntryID   int64 `uri:"entryId" binding:"required"`
}

func (server *Server) getAccountEntries(ctx *gin.Context) {
	var req getAccountEntriesRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.GetEntryParams{
		ID:        req.EntryID,
		AccountID: req.AccountID,
	}

	entry, err := server.store.GetEntry(ctx, args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, entry)
}

type listEntryAccount struct {
	AccountID int64 `uri:"id" binding:"required"`
}

func (server *Server) listEntry(ctx *gin.Context) {
	var req listEntryAccount
	var pagination Pagination

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindQuery(&pagination); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountParams{
		Limit:  pagination.PageSize - 1,
		Offset: (pagination.PageID - 1) * pagination.PageSize,
	}

	accounts, err := server.store.ListAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

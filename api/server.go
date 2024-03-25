package api

import (
	db "github.com/HectorSauR/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	store  db.Store
	router *gin.Engine
}

type Pagination struct {
	PageID   int32 `form:"page_id" binding:"required,min=0"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	//accounnts
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.PUT("/accounts/:id", server.updateAccount)

	//entries
	router.POST("/accounts/:id/entries", server.createEntry)
	router.GET("/accounts/:id/entries/:entryId", server.getAccountEntries)
	router.GET("/accounts/:id/entries", server.listEntry)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

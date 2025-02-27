package rest

import (
	"database/sql"
	"net/http"

	db "plbooking_go_structure1/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

// get a ticket request
type getTicketRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *HttpServer) getTicketbyID(ctx *gin.Context) {
	var req getTicketRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ticket, err := server.Pgdbc.GetTicketByID(ctx, req.ID)
	// If error(Internal Error)
	if err != nil {
		// if id not in db(SQL Query Error)
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, ticket)
}

// list tickets request
type listTicketsRequest struct {
	MatchID   int32 `form:"match_id" binding:"required"`
	PageIndex int32 `form:"page_index" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *HttpServer) listTickets(ctx *gin.Context) {
	var req listTicketsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	arg := db.ListTicketsWithDetailsParams{
		MatchID: req.MatchID,
		Limit:   req.PageSize,
		Offset:  (req.PageIndex - 1) * req.PageSize,
	}

	tickets, err := server.Pgdbc.ListTicketsWithDetails(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, tickets)
}

// list minimum price of tickets by match
func (server *HttpServer) listMinPriceOfTickets(ctx *gin.Context) {
	ticks, err := server.Pgdbc.ListMinPriceOfTicketsByMatch(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i := 0; i < len(ticks); i++ {
		ticks[i].MinPrice = string(ticks[i].MinPrice.([]byte))
	}
	ctx.JSON(http.StatusOK, ticks)
}

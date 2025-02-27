package rest

import (
	"net/http"

	db "plbooking_go_structure1/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

// list users request
type getSeatRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

// list stadiums request
func (server *HttpServer) getSeatByID(ctx *gin.Context) {
	var req getSeatRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	seat, err := server.Pgdbc.GetSeatByID(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, seat)
}

// list seats request
type listSeatsParams struct {
	StadiumID int32 `form:"stadium_id" binding:"required,min=1"`
	PageIndex int32 `form:"page_index" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=50"`
}

// list stadiums request
func (server *HttpServer) listSeats(ctx *gin.Context) {
	var req listSeatsParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	arg := db.ListSeatsParams{
		StadiumID: req.StadiumID,
		Limit:     req.PageSize,
		Offset:    (req.PageIndex - 1) * req.PageSize,
	}

	seats, err := server.Pgdbc.ListSeats(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, seats)
}

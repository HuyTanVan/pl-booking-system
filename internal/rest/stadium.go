package rest

import (
	"database/sql"
	"net/http"

	db "plbooking_go_structure1/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

// get a stadium request
type getStadiumRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *HttpServer) getStadiumByID(ctx *gin.Context) {
	var req getStadiumRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	sta, err := server.Pgdbc.GetStadiumByID(ctx, req.ID)
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
	ctx.JSON(http.StatusOK, sta)
}

// list stadiums request
type listStadiumsRequest struct {
	PageIndex int32 `form:"page_index" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *HttpServer) listStadiums(ctx *gin.Context) {
	var req listStadiumsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	arg := db.ListStadiumsParams{
		Limit:  req.PageSize,
		Offset: (req.PageIndex - 1) * req.PageSize,
	}

	stas, err := server.Pgdbc.ListStadiums(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, stas)
}

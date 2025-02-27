package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type listTicketsOfMatchRequest struct {
	MatchID   int32 `form:"match_id" binding:"required"`
	PageIndex int32 `form:"page_index" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *HttpServer) listTicketsOfMatch(ctx *gin.Context) {
	var req listTicketsOfMatchRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
}

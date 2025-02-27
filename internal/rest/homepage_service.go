package rest

import (
	"net/http"
	db "plbooking_go_structure1/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

type HomepageRequest struct {
	Range     string `form:"range"`
	SortBy    string `form:"sort_by"`
	PageIndex int32  `form:"page_index" binding:"required,min=1"`
	PageSize  int32  `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *HttpServer) getHomepage(ctx *gin.Context) {
	var req listMatchesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 1. get keys from request used for query
	// w=week, m=month, t=today, a=all
	if req.Range != "w" && req.Range != "m" && req.Range != "t" {
		req.Range = "a"
	}
	// d=date, p=price
	if req.SortBy != "d" && req.SortBy != "p" {
		req.SortBy = "d"
	}
	// 2. query data for homepage
	arg := db.FetchHomePageDetailsParams{
		Range:  req.Range,
		Sortby: req.SortBy,
		Limit:  req.PageSize,
		Offset: (req.PageIndex - 1) * req.PageSize,
	}
	data, err := server.Pgdbc.FetchHomePageDetails(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 3. homepage data (match, minimum ticket price, nums of ticker available)
	ctx.JSON(http.StatusOK, data)
}

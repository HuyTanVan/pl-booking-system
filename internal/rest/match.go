package rest

import (
	"database/sql"

	"net/http"
	"time"

	db "plbooking_go_structure1/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

// get a match request
type getMatchRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *HttpServer) getMatchByID(ctx *gin.Context) {
	var req getMatchRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	match, err := server.Pgdbc.GetMatchByID(ctx, req.ID)
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
	ctx.JSON(http.StatusOK, match)
}

// list matchs request
type listMatchesRequest struct {
	Range     string `form:"range"`
	SortBy    string `form:"sort_by"`
	PageIndex int32  `form:"page_index" binding:"required,min=1"`
	PageSize  int32  `form:"page_size" binding:"required,min=5,max=10"`
}
type listMatchesResponse struct {
	HomeTeam  string    `json:"home_team"`
	AwayTeam  string    `json:"away_team"`
	Stadium   string    `json:"stadium"`
	MatchDate time.Time `json:"match_date"`
	Session   string    `json:"session"`
	Status    string    `json:"status"`
}

func (server *HttpServer) listMatchesWithDetails(ctx *gin.Context) {
	var req listMatchesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if req.Range != "w" && req.Range != "m" && req.Range != "t" {
		req.Range = "a"
	}
	if req.SortBy != "d" && req.SortBy != "p" {
		req.SortBy = "d"
	}
	arg := db.ListMatchesWithDetailsParams{
		Range:  req.Range,
		Sortby: req.SortBy,
		Limit:  req.PageSize,
		Offset: (req.PageIndex - 1) * req.PageSize,
	}
	matches, err := server.Pgdbc.ListMatchesWithDetails(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, matches)
}

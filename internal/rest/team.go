package rest

import (
	"database/sql"
	"net/http"

	db "plbooking_go_structure1/internal/db/sqlc"

	"github.com/gin-gonic/gin"
)

// get a team request
type getTeamRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *HttpServer) getTeamByID(ctx *gin.Context) {
	var req getTeamRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Syntax Error": err.Error(),
		})
		return
	}

	team, err := server.Pgdbc.GetTeamByID(ctx, req.ID)
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
	ctx.JSON(http.StatusOK, team)
}

// list users request
type listTeamsRequest struct {
	PageIndex int32 `form:"page_index" binding:"required,min=1"`
	PageSize  int32 `form:"page_size" binding:"required,min=5,max=20"`
}

func (server *HttpServer) listTeams(ctx *gin.Context) {
	var req listTeamsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	arg := db.ListTeamsParams{
		Limit:  req.PageSize,
		Offset: (req.PageIndex - 1) * req.PageSize,
	}

	teams, err := server.Pgdbc.ListTeams(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, teams)
}

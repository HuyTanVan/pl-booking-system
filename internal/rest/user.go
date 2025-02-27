package rest

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	db "plbooking_go_structure1/internal/db/sqlc"
	"plbooking_go_structure1/internal/token"

	// "github.com/HuyTanVan/soccer_booking_ticket/util"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

/*
	This includes all functions serving apis.
	User can UPDATE their FirstName, LastName, PhoneNumber, and Password
*/

// create a get user request

// user response
type userResponse struct {
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type getUserRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *HttpServer) getUser(ctx *gin.Context) {
	var req getUserRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Syntax Error": err.Error(),
		})
		return
	}

	user, err := server.Pgdbc.GetUser(ctx, req.ID)
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
	authPayload := ctx.MustGet("authorization_payload").(*token.Payload)

	if user.Email != authPayload.Email {

		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "not authorized to access this user",
		})
		return
	}
	var userRes = userResponse{
		FirstName:   user.FirstName.String,
		LastName:    user.LastName.String,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber.String,
	}
	ctx.JSON(http.StatusOK, userRes)
}

// create a update user request
type updateUserRequest struct {
	// ID			int32 		   `json:id`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

// http:localhost:8000/api/users/:id [UPDATE]
func (server *HttpServer) updateUser(ctx *gin.Context) {
	var req *updateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error ShouldBindJSON": err.Error(),
		})
		return
	}

	arg := db.UpdateUserParams{
		FirstName:   sql.NullString{String: req.FirstName, Valid: req.FirstName != ""},
		LastName:    sql.NullString{String: req.LastName, Valid: req.LastName != ""},
		PhoneNumber: sql.NullString{String: req.PhoneNumber, Valid: req.PhoneNumber != ""},
		ID:          int32(6),
	}
	fmt.Println("MY ARG", arg)
	err := server.Pgdbc.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error UpdateUser": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})
}

// type loginUserRequest struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required,min=6"`
// }
// type loginUserResponse struct {
// 	SessionID             uuid.UUID    `json"session_id"`
// 	AccessToken           string       `json:"access_token"`
// 	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
// 	RefreshToken          string       `json:"refresh_token"`
// 	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
// 	User                  userResponse `json:"user"`
// }

// func (server *HttpServer) loginUser(ctx *gin.Context) {
// 	var req loginUserRequest
// 	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	user, err := server.Pgdbc.GetUserByEmail(ctx, req.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			ctx.JSON(http.StatusNotFound, err.Error())
// 			return
// 		}
// 		ctx.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	err = util.CheckPassword(req.Password, user.Password)
// 	if err != nil {
// 		ctx.JSON(http.StatusUnauthorized, err.Error())
// 		return
// 	}

// 	accessToken, accessPayload, err := server.tokenmaker.CreateToken(req.Email, server.config.AccessTokenDuration)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	refreshToken, refreshPayload, err := server.tokenmaker.CreateToken(req.Email, server.config.RefreshTokenDuration)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}
// 	session, err := server.Pgdbc.CreateSession(ctx, db.CreateSessionParams{
// 		ID:           refreshPayload.ID,
// 		Email:        user.Email,
// 		RefreshToken: refreshToken,
// 		UserAgent:    "",
// 		ClientIp:     "",
// 		IsBlocked:    false,
// 		ExpiresAt:    refreshPayload.ExpiredAt,
// 	})
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	res := loginUserResponse{
// 		SessionID:             session.ID,
// 		AccessToken:           accessToken,
// 		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
// 		RefreshToken:          refreshToken,
// 		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
// 		User: userResponse{FirstName: user.FirstName.String,
// 			LastName:    user.LastName.String,
// 			Email:       user.Email,
// 			PhoneNumber: user.PhoneNumber.String,
// 		},
// 	}
// 	ctx.JSON(http.StatusOK, res)
// }

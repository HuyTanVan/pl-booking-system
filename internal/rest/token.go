package rest

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *HttpServer) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	refreshPayload, err := server.Token.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	session, err := server.Pgdbc.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, err.Error())
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if session.Email != refreshPayload.Email {
		err := fmt.Errorf("incorrect session user")
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		ctx.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	dur, err := time.ParseDuration(server.Config.JWTToken.AccessTokenDuration)
	if err != nil {
		err := fmt.Errorf("can not parse time")
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	accessToken, accessPayload, err := server.Token.CreateToken(
		refreshPayload.Email,
		dur,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK, rsp)

}

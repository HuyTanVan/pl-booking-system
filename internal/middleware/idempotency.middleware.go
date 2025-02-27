package middleware

import (
	"database/sql"
	"errors"
	"net/http"
	db "plbooking_go_structure1/internal/db/sqlc"
	"plbooking_go_structure1/internal/utils"

	// db "github.com/HuyTanVan/soccer_booking_ticket/db/sqlc"
	// "github.com/HuyTanVan/soccer_booking_ticket/util"

	"github.com/gin-gonic/gin"
)

const (
	idempotencyHeaderKey = "idempotency-key"
	idempotencyKey       = "idempotency_key"
)

func IdempotencyMiddleware(store *db.Store) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idemHeader := ctx.GetHeader(idempotencyHeaderKey)
		if idemHeader == "" {

			err := errors.New("idempotency-key header is not provided")
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		idem, err := store.GetIdempotencyByIdempotencyKey(ctx, idemHeader)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.Set(idempotencyKey, idemHeader)
				ctx.Next()
				return
			}
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}
		// not done
		status_code := int(idem.ResponseStatusCode)
		rspEncode := utils.EncodeByteToBase64(idem.ResponseBody)
		rsp, _ := utils.DecodeFromBase64(rspEncode)
		ctx.AbortWithStatusJSON(status_code, string(rsp))
	}
}

package rest

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	db "plbooking_go_structure1/internal/db/sqlc"
	worker "plbooking_go_structure1/internal/redis_workers"
	"plbooking_go_structure1/internal/token"
	"plbooking_go_structure1/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/paymentintent"
)

// stripe payment_intents create --amount="9999" --currency="usd"
type paymentIntentRequest struct {
	Amount        int64  `form:"amount" binding:"required,min=100"`
	PaymentMethod string `form:"payment_method" binding:"required"` // only support card now
}
type paymentIntentResponse struct {
	ClientSecret string `json:"client_secret"`
}

func (server *HttpServer) createPaymentIntent(ctx *gin.Context) {
	var req paymentIntentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// initialize Stripe Payment Intent Params
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(false),
		},
		// PaymentMethod: stripe.String("pmc_1QaqMOP6qrjtDGURckR4mwID"),
		// PaymentMethod: stripe.String("pm_card_visa"),
		PaymentMethodTypes: []*string{
			stripe.String(req.PaymentMethod),
		},
	}

	// create new payment intent
	result, err := paymentintent.New(params)
	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"stripe error": stripeErr.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// convert amount integer type to string
	// Convert integer to string
	str := strconv.FormatInt(req.Amount, 10)
	// fmt.Println(str)
	// Insert a decimal point at the correct position
	formatted_amount := str[:2] + "." + str[2:]
	// create a payment record with payment_intent ID and status=intent
	pay, err := server.Pgdbc.CreatePayment(ctx, db.CreatePaymentParams{
		PaymentKey:      result.ID,
		Amount:          formatted_amount,
		PaymentMethodID: 1,
		Status:          "intent",
		Date:            time.Now(),
		UpdatedAt:       time.Now(),
	})
	fmt.Println(pay)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error payment": err.Error(),
		})
		return
	}

	rps := paymentIntentResponse{ClientSecret: result.ClientSecret}
	ctx.JSON(http.StatusOK, rps)
}

type confirmPaymentRequest struct {
	// IdempotencyKey string `form:"idempotency_key"`
	TicketID     int32  `form:"ticket_id"`
	Quantity     int32  `form:"quantity_selected"`
	ClientSecret string `form:"client_secret"`
}

// subtract seat_avalable in seats table by amount of request and create order record
func (server *HttpServer) confirmPaymentIntent(ctx *gin.Context) {
	var req confirmPaymentRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	authPayload := ctx.MustGet("authorization_payload").(*token.Payload)
	user_id, err := server.Pgdbc.GetUserByEmail(ctx, authPayload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "authorization_payload",
		})
		return
	}
	// get idem_key
	idem_key, exists := ctx.Get("idempotency_key")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "can not get idempotency_key",
		})
		return
	}
	// get ticket obj
	tick, err := server.Pgdbc.GetTicketByID(ctx, req.TicketID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// get seat obj for seat avalability checking
	seat, err := server.Pgdbc.GetSeatByID(ctx, tick.SeatID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx_arg := db.BuyTicketTxParams{
		Seat: seat,
		UpdateSeatAvailableParams: db.UpdateSeatAvailableParams{
			ID:            seat.ID,
			SeatAvailable: seat.SeatAvailable - req.Quantity,
		},
		CreateOrderParams: db.CreateOrderParams{
			UserID:           user_id.ID,
			TicketID:         req.TicketID,
			Quantity:         req.Quantity,
			TotalPrice:       "9999",
			AdditionalFeesID: 1,
			PaymentKey:       req.ClientSecret,
		},
		CreateIdempotencyParams: db.CreateIdempotencyParams{
			UserID:         user_id.ID,
			IdempotencyKey: idem_key.(string),
		},
	}
	rsp, err := server.Pgdbc.PurchaseTicketTx(ctx, tx_arg)
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error order": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, rsp)
}

type checkoutRequest struct {
	TicketID int32 `form:"ticket_id" binding:"required"`
	Quantity int8  `form:"quantity_selected" binding:"required"`
}

type ticketDetailResponse struct {
	ID              int32     `json:"id"`
	MatchID         int32     `json:"match_id"`
	SeatID          int32     `json:"seat_id"`
	Price           string    `json:"price"`
	TicketAvailable bool      `json:"ticket_available"`
	HomeTeam        string    `json:"home_team"`
	AwayTeam        string    `json:"away_team"`
	MatchDate       time.Time `json:"match_date"`
	StadiumLocation string    `json:"stadium_location"`
	Block           string    `json:"block"`
	Row             string    `json:"row"`
	SeatColumn      string    `json:"seat_column"`
}
type checkoutResponse struct {
	TicketDetail   ticketDetailResponse `json:"ticket_detail"`
	Quantity       int32                `json:"quantity"`
	AdditionalFees string               `json:"additional_fees"`
	TotalPrice     string               `json:"total_price"`
}

// total amount is sum of ticket_price*quantity + (ticket_price*quantity*additional_fees)
func (server *HttpServer) reviewCheckout(ctx *gin.Context) {
	var req checkoutRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	fees, err := server.Pgdbc.GetTotalFeesByID(ctx, 1)
	totalFeesStr := string(fees.TotalFees)
	fmt.Println(totalFeesStr)
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
	ticket, err := server.Pgdbc.GetTicketWithDetails(ctx, req.TicketID)
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

	totalPriceStr, err := utils.CalculateTotalPrice(ticket.Price, int32(req.Quantity), totalFeesStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	rsp := &checkoutResponse{TicketDetail: ticketDetailResponse{
		ID:              ticket.ID,
		MatchID:         ticket.MatchID,
		SeatID:          ticket.SeatID,
		Price:           ticket.Price,
		TicketAvailable: ticket.TicketAvalable,
		HomeTeam:        ticket.HomeTeam.String,
		AwayTeam:        ticket.AwayTeam.String,
		MatchDate:       ticket.MatchDate.Time,
		StadiumLocation: ticket.StadiumLocation.String,
		Block:           ticket.Block.String,
		Row:             ticket.Row.String,
		SeatColumn:      ticket.SeatColumn.String,
	},
		Quantity:       int32(req.Quantity),
		AdditionalFees: totalFeesStr,
		TotalPrice:     totalPriceStr,
	}
	ctx.JSON(http.StatusOK, rsp)

}

type webhookSuccessResponse struct {
	Event       string `json:"event"`
	Amount      string `json:"amount"`
	Status      string `json:"status,omitempty"`
	ReceiptLink string `json:"receipt_link,omitempty"`
}
type webhookFailureResponse struct {
	Event          string `json:"event"`
	Amount         string `json:"amount"`
	FailureMessage string `json:"failure_message"`
	Status         string `json:"status"`
}

// nested structs to collect client_secrect from webhook request
// iinitial
type PaymentIntent struct {
	ID string `json:"id"`
}
type EventData struct {
	Object PaymentIntent `json:"object"`
}
type webHookRequest struct {
	Type string    `json:"type"`
	Data EventData `json:"data"`
}

func (server *HttpServer) handleWebhook(ctx *gin.Context) {
	var req webHookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println("WEBHOOK TYPE:", req.Type, req.Data.Object.ID)

	switch req.Type {
	case "payment_intent.created":
		ctx.JSON(http.StatusOK, nil)
		return
	case "payment_intent.succeeded":
		payment, err := server.Pgdbc.GetPaymentByPaymentKey(ctx, req.Data.Object.ID)
		if err != nil {
			// add no row err here
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// can get order record here to compare payment.paymenty_key and order.payment_key
		// or payment.total and order.total. Extra security to ensure existence of records
		err = server.Pgdbc.UpdatePaymentStatusByPaymentKey(ctx,
			db.UpdatePaymentStatusByPaymentKeyParams{
				PaymentKey: payment.PaymentKey,
				Status:     "succeeded",
			})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		err = server.TaskDistributor.DistributeTaskSendOrderSuccess(ctx, &worker.PayloadSendOrderSuccess{
			OrderID: 1,
			Email:   "huyhandsome189@gmail.com"})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		return
	case "charge.succeeded":
		return
	case "charge.updated":
		return
	case "charge.failed":
		return
	case "payment_intent.payment_failed":
		payment, err := server.Pgdbc.GetPaymentByPaymentKey(ctx, req.Data.Object.ID)
		if err != nil {
			// add no row err here
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// can get order record here to compare payment.paymenty_key and order.payment_key
		// or payment.total and order.total. Extra security to ensure existence of records
		err = server.Pgdbc.UpdatePaymentStatusByPaymentKey(ctx,
			db.UpdatePaymentStatusByPaymentKeyParams{
				PaymentKey: payment.PaymentKey,
				Status:     "failed",
			})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		err = server.Pgdbc.UpdateSeatAvailable(ctx, db.UpdateSeatAvailableParams{SeatAvailable: 100})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		return
	default:
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": "Event Type not found",
		})
		return
	}

	// if req.Type == "payment_intent.created" {
	// 	ctx.JSON(http.StatusOK, nil)
	// 	return
	// }
	// // if payment_intent.succeeded, update payment_status=succeeded
	// if req.Type == "payment_intent.succeeded" {
	// 	// query payment and order if they both have record then update payment_status=succeeded

	// 	ctx.JSON(http.StatusOK, nil)
	// 	return
	// }
	// if req.Type == "charge.succeeded" {
	// 	// payment is ok, but no content being returned
	// 	ctx.JSON(http.StatusNoContent, nil)
	// 	return
	// }

}

// // delete later
// type confirmPaymentRequest struct {
// 	ConfirmPayment string `form:"confirm_payment"`
// }

// func (server *HttpServer) confirmPaymentIntent(ctx *gin.Context) {
// 	var req confirmPaymentRequest
// 	if err := ctx.ShouldBindQuery(&req); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	params := &stripe.PaymentMethodParams{
// 		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
// 		Card: &stripe.PaymentMethodCardParams{
// 			Token: stripe.String("tok_visa"), // Use the test token directly
// 			// ExpMonth: stripe.Int64(12),
// 			// ExpYear:  stripe.Int64(2023),
// 			// CVC:      stripe.String("123"),
// 		},
// 	}
// 	p, err := paymentmethod.New(params)

// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		},
// 		)
// 	}
// 	// ctx.JSON(http.StatusOK, paymentMethod)
// 	s := &stripe.PaymentIntentConfirmParams{
// 		PaymentMethod: stripe.String(p.ID), // Use the actual payment method ID you want to confirm

// 	}
// 	paymentIntent, err := paymentintent.Confirm(req.ConfirmPayment, s)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		},
// 		)
// 	}
// 	ctx.JSON(http.StatusOK, paymentIntent)
// }

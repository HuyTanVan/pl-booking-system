package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
)

type BuyTicketTxParams struct {
	Seat
	UpdateSeatAvailableParams
	CreateOrderParams
	CreateIdempotencyParams
	// AfterCreate func(user User, token EmailVerificationToken) error
}
type BuyTicketTxResult struct {
	Order       Order
	Idempotency string `json:"idempotency_key"`
}

const ()

func (store *Store) PurchaseTicketTx(ctx context.Context, arg BuyTicketTxParams) (BuyTicketTxResult, error) {
	var result BuyTicketTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		// check if seat_available enough for request
		if !(arg.Seat.SeatAvailable-arg.CreateOrderParams.Quantity >= 0) {
			err = fmt.Errorf("not enough seats available: requested %d, available %d",
				arg.CreateOrderParams.Quantity, arg.Seat.SeatAvailable)
			return err
		}
		// update seat_availabitily
		err = q.UpdateSeatAvailable(ctx, arg.UpdateSeatAvailableParams)
		if err != nil {
			return err
		}
		// create an order record
		result.Order, err = q.CreateOrder(ctx, arg.CreateOrderParams)
		if err != nil {
			return err
		}
		// assign idemkey before creating record
		result.Idempotency = arg.CreateIdempotencyParams.IdempotencyKey
		// marshal result.Order into binary
		orderRes, err := json.Marshal(result)
		if err != nil {
			fmt.Println("error marshalling Order:", err)
			return err
		}

		idemParams := CreateIdempotencyParams{
			UserID:             arg.CreateIdempotencyParams.UserID,
			IdempotencyKey:     arg.CreateIdempotencyParams.IdempotencyKey,
			ResponseStatusCode: 201,
			ResponseHeaders:    []byte("{\"authorization\": \"authorization_payload\", \"idempotency-key\": \"idempotency_key\"}"),
			ResponseBody:       orderRes,
			CreatedAt:          time.Now(),
		}
		// create idempotency
		_, err = q.CreateIdempotency(ctx, idemParams)
		if err != nil {
			err = fmt.Errorf("can not create idempotency_key %s",
				arg.CreateIdempotencyParams.IdempotencyKey)
			return err
		}
		return err
	})
	return result, err
}

// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type AdditionalFee struct {
	ID                   int32     `json:"id"`
	ServiceFee           string    `json:"service_fee"`
	VatFee               string    `json:"vat_fee"`
	PaymentProcessingFee string    `json:"payment_processing_fee"`
	DeliveryFee          string    `json:"delivery_fee"`
	VipFee               string    `json:"vip_fee"`
	SpecificEventFee     string    `json:"specific_event_fee"`
	Description          string    `json:"description"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type EmailVerificationToken struct {
	ID     int32  `json:"id"`
	UserID int32  `json:"user_id"`
	Token  string `json:"token"`
}

type Idempotency struct {
	UserID             int32     `json:"user_id"`
	IdempotencyKey     string    `json:"idempotency_key"`
	ResponseStatusCode int16     `json:"response_status_code"`
	ResponseHeaders    []byte    `json:"response_headers"`
	ResponseBody       []byte    `json:"response_body"`
	CreatedAt          time.Time `json:"created_at"`
}

type Match struct {
	ID         int32     `json:"id"`
	HomeTeamID int32     `json:"home_team_id"`
	AwayTeamID int32     `json:"away_team_id"`
	StadiumID  int32     `json:"stadium_id"`
	MatchDate  time.Time `json:"match_date"`
	Session    string    `json:"session"`
	Status     string    `json:"status"`
}

type Order struct {
	ID       int32 `json:"id"`
	UserID   int32 `json:"user_id"`
	TicketID int32 `json:"ticket_id"`
	// must be positive
	Quantity         int32     `json:"quantity"`
	TotalPrice       string    `json:"total_price"`
	AdditionalFeesID int32     `json:"additional_fees_id"`
	PaymentKey       string    `json:"payment_key"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type Payment struct {
	ID int32 `json:"id"`
	// stripe payment key
	PaymentKey string `json:"payment_key"`
	// can not be negative
	Amount          string    `json:"amount"`
	PaymentMethodID int32     `json:"payment_method_id"`
	Status          string    `json:"status"`
	Date            time.Time `json:"date"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type PaymentMethod struct {
	ID         int32  `json:"id"`
	UserID     int32  `json:"user_id"`
	Type       string `json:"type"`
	CardNumber string `json:"card_number"`
}

type Seat struct {
	ID            int32  `json:"id"`
	StadiumID     int32  `json:"stadium_id"`
	Block         string `json:"block"`
	Row           string `json:"row"`
	SeatColumn    string `json:"seat_column"`
	SeatAvailable int32  `json:"seat_available"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Stadium struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	// can not be negative
	Capacity    int32 `json:"capacity"`
	IsAvailable bool  `json:"is_available"`
}

type Team struct {
	ID        int32  `json:"id"`
	Name      string `json:"name"`
	StadiumID int32  `json:"stadium_id"`
	// url for an image
	Logo sql.NullString `json:"logo"`
}

type Ticket struct {
	ID      int32 `json:"id"`
	MatchID int32 `json:"match_id"`
	SeatID  int32 `json:"seat_id"`
	// can not be negative
	Price       string    `json:"price"`
	IsAvailable bool      `json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type User struct {
	ID          int32          `json:"id"`
	FirstName   sql.NullString `json:"first_name"`
	LastName    sql.NullString `json:"last_name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	PhoneNumber sql.NullString `json:"phone_number"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

package rest

import (
	"fmt"
	"log"
	"plbooking_go_structure1/global"
	db "plbooking_go_structure1/internal/db/sqlc"
	"plbooking_go_structure1/internal/middleware"
	"plbooking_go_structure1/internal/token"
	"plbooking_go_structure1/pkg/setting"

	worker "plbooking_go_structure1/internal/redis_workers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stripe/stripe-go/v81"
)

type HttpServer struct {
	Config          setting.Config
	Pgdbc           db.Store      // postgresql with sqlc tool
	Rdb             *redis.Client // a single redis server
	Router          *gin.Engine
	Token           token.IMaker // jwt token maker
	TaskDistributor worker.TaskDistributor
}

// HTTP
func (server *HttpServer) StartServer() error {
	server.setupRouter()
	server.setupStripeSecretKey()
	log.Println("http server running on port", server.Config.HTTPServer.Port)
	return server.Router.Run(fmt.Sprintf(":%d", server.Config.HTTPServer.Port))
}

// Assigns API routes
func (server *HttpServer) setupRouter() {
	// Login Route
	// router.POST("api/v1/users/login", server.loginUser)

	// server.Router.POST("api/v1/tokens/renew_access", server.renewAccessToken)
	// api/v1 group
	apiV1Group := server.Router.Group("/api/v1")
	apiV1Group.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Allow the frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	//
	apiV1Group.GET("/getHomePage", server.getHomepage)
	apiV1Group.GET("/matches/:id/tickets", server.listTicketsOfMatch)
	// apiV1Group.GET("/getHomePage", server.getHomepage)

	apiV1Group.GET("/teams", server.listTeams)
	apiV1Group.GET("/teams/:id", server.getTeamByID)
	// stadiums
	apiV1Group.GET("/stadiums", server.listStadiums)
	apiV1Group.GET("/stadiums/:id", server.getStadiumByID)
	// matches
	apiV1Group.GET("/matches", server.listMatchesWithDetails)
	apiV1Group.GET("/matches/:id", server.getMatchByID)
	// seats
	apiV1Group.GET("/seats", server.listSeats)
	apiV1Group.GET("/seats/:id", server.getSeatByID)
	// tickets
	apiV1Group.GET("/tickets", server.listTickets)

	apiV1Group.GET("/tickets/:id", server.getTicketbyID)
	apiV1Group.GET("/tickets/min_prices", server.listMinPriceOfTickets)

	// auth group
	authGroup := server.Router.Group("/auth/api/v1")
	authGroup.Use(middleware.CORSMiddleware())
	authGroup.OPTIONS("/*cors", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})
	authGroup.Use(middleware.AuthMiddleware(server.Token))
	authGroup.GET("/users/:id", server.getUser)    // A user gets their information
	authGroup.PUT("/users/:id", server.updateUser) // A user update themself

	authGroup.GET("/checkout/review", server.reviewCheckout)
	authGroup.POST("/checkout/create_payment_intent", server.createPaymentIntent)
	authGroup.POST("/checkout/confirm_payment_intent", middleware.IdempotencyMiddleware(&server.Pgdbc), server.confirmPaymentIntent)
	authGroup.POST("/checkout/webhook", server.handleWebhook)
}

func (server *HttpServer) setupStripeSecretKey() {
	stripe.Key = global.Config.Stripe.StripeSecretKey
	log.Println("stripe service is running", stripe.Key)
}

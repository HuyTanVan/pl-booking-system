package externalrouters

// import (
// 	"github.com/gin-gonic/gin"
// )

// type IUserRouter interface {
// 	InitUserRouter(server *HttpServer, Router *gin.RouterGroup)
// }
// type UserRouter struct {
// }

// func (ur *UserRouter) InitUserRouter(server *HttpServer, Router *gin.RouterGroup) IUserRouter {
// 	// public router
// 	userRouterPublic := Router.Group("/users")
// 	{
// 		userRouterPublic.POST("/register")
// 		userRouterPublic.POST("/login")
// 		// userRouterPublic.POST("/login")
// 	}
// 	// teams
// 	teamRouterPublic := Router.Group("/teams", server.listTeams)
// 	{
// 		teamRouterPublic.GET("/teams")
// 		teamRouterPublic.GET("/teams/:id")
// 	}
// 	// match
// 	matchRouterPublic := Router.Group("/matches")
// 	{
// 		matchRouterPublic.GET("/matches")
// 		matchRouterPublic.GET("/matches/:id")
// 	}
// 	// stadium
// 	stadiumRouterPublic := Router.Group("/stadiums")
// 	{
// 		stadiumRouterPublic.GET("/stadiums")
// 		stadiumRouterPublic.GET("/stadiums/:id")
// 	}
// 	// ticket
// 	seatRouterPublic := Router.Group("/seats")
// 	{
// 		seatRouterPublic.GET("/seats")
// 		seatRouterPublic.GET("/seats/:id")
// 	}
// 	// seat
// 	ticketRouterPublic := Router.Group("/tickets")
// 	{
// 		ticketRouterPublic.GET("/tickets")
// 		ticketRouterPublic.GET("/tickets/:id")
// 	}

// 	// private router(authenticated user group)
// 	userRouterPrivate := Router.Group("/users")
// 	userRouterPrivate.Use()
// 	{
// 		userRouterPrivate.GET("/get_info/:id")
// 	}
// }

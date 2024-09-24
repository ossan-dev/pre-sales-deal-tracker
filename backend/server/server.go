package server

import (
	"fmt"
	"time"

	db "github.com/ekefan/deal-tracker/internal/db/sqlc"
	"github.com/ekefan/deal-tracker/internal/token"
	"github.com/ekefan/deal-tracker/internal/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Server contains fields required by a server instance
type Server struct {
	Router     *gin.Engine      // Router an instance of gin.Engine
	Store      db.Store         // Store the database interface for interacting with the db
	EnvVar     Config           // EnvVar holds the environment variables loaded into the server instance
	TokenMaker token.TokenMaker // interface for creating and managing tokens
}

// NewServer create a server instance, having a router that connect api endpoints
func NewServer(store db.Store, config Config) (*Server, error) {
	tokenMaker, err := token.NewPaseto(config.SymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("couln't create token: %w", err)
	}
	return &Server{
		Store:      store,
		EnvVar:     config,
		TokenMaker: tokenMaker,
	}, nil
}

// SetupRouter sets up a router, register cors, and routes for api endpoints
func (s *Server) SetupRouter() {
	router := gin.Default()
	// register validation to gin context
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("valid-role", utils.RoleValidator)
	}
	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Location"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(corsConfig))
	authRoute := router.Group("/").Use(authMiddleware(s.TokenMaker))

	// FIXME: we've to discern about the "users" resource. We have three options:
	// 1. "users" are only used to keep track of logins. Stick to addresses like "/register", "login", "logout", "password-reset", and so on.
	// 2. "users" are part of the domain of our application. In this case use "GET /users", "POST /users", "PUT /users/:id", and so on.
	// 3. Mix & Match: clearly separate those endpoints: the one used for signing-up, login, logout, ecc. from the ones used as domain in our application. Potentially, we could also have separated table.

	router.POST("/users", s.adminCreateUserHandler)
	router.POST("/users/login", s.userLogin)

	authRoute.POST("/sales/pitch-requests", s.salesCreatePitchReqHandler)
	authRoute.POST("/admin/deals", s.adminCreateDealHandler)
	authRoute.GET("/deals", s.getOngoingDeals)
	authRoute.GET("/admin/deals", s.getFilteredDeals)
	authRoute.GET("/admin/pitch-requests", s.adminGetPitchRequests)
	authRoute.GET("/sales/pitch-requests", s.salesViewPitchRequests)
	authRoute.GET("/admin/deals/:deal_id", s.getDealsById)
	authRoute.GET("/sales/deals", s.getSalesDeals)
	authRoute.GET("/admin/users", s.listUsersHandler)

	// I may not have gotten the desing right from the beginning, but I think from what I've done,
	// password is a sub resource from the user... Can I explain in our next f
	authRoute.PATCH("/admin/users/passwordresets", s.resetPassword)
	authRoute.PATCH("/admin/deals", s.adminUpdateDealHandler)
	authRoute.PATCH("/admin/users", s.adminUpdateUserHandler)
	authRoute.PATCH("/users/passwords", s.updatePassWordLoggedIn)
	authRoute.PATCH("/sales/pitch-requests", s.updatePitchReqHandler)
	authRoute.DELETE("admin/users/:id", s.adminDeleteUserHandler)
	authRoute.DELETE("/admin/deals/:deal_id", s.adminDeleteDealHandler)
	authRoute.DELETE("/sales/pitch-requests/:sales_username/:sales_rep_id/:pitch_id", s.salesDeletePitchReqHandler)

	s.Router = router
}

// StartServer starts the app server, takes the hostAddress
// listens and serves on that address
func (s *Server) StartServer(hostAddress string) error {
	err := s.Router.Run(hostAddress)
	if err != nil {
		return err
	}
	return nil
}

// errorResponse a custom error reponse handler for reusability
// takes the  error (err) returns a gin.H{} struct with an error
// field equal to err.Error()
func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

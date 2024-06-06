package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rachitkawar/boilerplate-go/src/common"
	"github.com/rachitkawar/boilerplate-go/src/internal/domain/jwt"
	v1 "github.com/rachitkawar/boilerplate-go/src/internal/server/v1"
	"net/http"
)

type Server struct {
	router     *gin.Engine
	httpServer *http.Server
	v1Jwt      *jwt.TokenMaster
}

func InitializeServer(v1Jwt *jwt.TokenMaster) *Server {
	gin.SetMode(common.GetEnv("GIN_MODE"))
	router := gin.New()

	server := &Server{
		router: router,
		v1Jwt:  v1Jwt,
	}
	server.setupMiddleware()
	server.setupErrorHandling()
	server.setupRoutes()

	return server

}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		common.Log.Error(err)
		return err
	}
	return nil
}

func (s *Server) Run(Port string) {
	s.httpServer = &http.Server{
		Addr:    Port,
		Handler: s.router,
	}
	err := s.router.Run(Port)
	if err != nil {
		common.Log.Fatal(err)
	}

}

// setupMiddleware adds middleware to the Gin router
func (s *Server) setupMiddleware() {
	s.router.Use(Logger())
	s.router.Use(gin.Recovery())
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	s.router.Use(cors.New(config))
	// Add any custom middleware here
}

// setupErrorHandling manages error handling for the Gin router
func (s *Server) setupErrorHandling() {
	s.router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	s.router.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"message": "Method Not Allowed"})
	})

}

// setupRoutes defines the routes for the application
func (s *Server) setupRoutes() {
	api := s.router.Group("/api")

	//pass the pointer address for all the domains required
	v1.InitializeV1Routes(api, s.v1Jwt)
}

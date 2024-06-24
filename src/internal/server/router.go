package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rachitkawar/boilerplate-go/src/internal/domain"
	"github.com/rachitkawar/boilerplate-go/src/internal/server/docs"
	v1 "github.com/rachitkawar/boilerplate-go/src/internal/server/v1"
	"github.com/rachitkawar/boilerplate-go/src/utils"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Server struct {
	router     *gin.Engine
	httpServer *http.Server
	srv        *domain.Service
}

func InitializeServer(
	srv *domain.Service,

) *Server {
	gin.SetMode(utils.GetEnv("GIN_MODE"))
	router := gin.New()

	server := &Server{
		router: router,
		srv:    srv,
	}
	server.setupMiddleware()
	server.setupErrorHandling()
	server.setupRoutes()
	server.setupDocs()

	return server

}

func (s *Server) setupDocs() {
	docs.SwaggerInfo.Title = "API documentation"
	docs.SwaggerInfo.Description = "API Documentation for Basic Auth server"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "authSystem.swagger.io"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	s.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//	http://localhost:8080/swagger/index.html
}

func (s *Server) Shutdown(ctx context.Context) error {
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		utils.Log.Error(err)
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
		utils.Log.Fatal(err)
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
	v1.InitializeV1Routes(api, s.srv)
}

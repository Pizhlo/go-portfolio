package api

import (
	"fmt"
	"net/http"

	db "github.com/Pizhlo/go-portfolio/db/sqlc"
	"github.com/Pizhlo/go-portfolio/token"
	"github.com/Pizhlo/go-portfolio/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for banking service
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

const (
	ContentTypeBinary = "application/octet-stream"
	ContentTypeForm   = "application/x-www-form-urlencoded"
	ContentTypeJSON   = "application/json"
	ContentTypeHTML   = "text/html; charset=utf-8"
	ContentTypeText   = "text/plain; charset=utf-8"
)

func (server *Server) setupRouter() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("static/", "./static/")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "animated_text.html", gin.H{"title": "Главная страница"})
	})

	router.GET("/admin/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth_admin.html", gin.H{"title": "Войти как админ"})
	})

	router.POST("/admin", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", gin.H{"title": "Админ панель"})
	})

	router.Run(":9090")

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

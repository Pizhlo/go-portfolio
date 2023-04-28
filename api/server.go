package api

import (
	"net/http"

	db "github.com/Pizhlo/go-portfolio/db/sqlc"
	"github.com/Pizhlo/go-portfolio/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for banking service
type Server struct {
	config util.Config
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("static/", "./static/")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Главная страница"})
	})

	router.Run()

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

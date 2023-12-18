package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Host string
	Port string
}

type Server struct {
	Config Config
}

// NewServer Creates a new server instance.
func NewServer(c Config) Server {
	return Server{Config: c}
}

// Run Starting HTTP server.
func (s *Server) Run() {
	gin.SetMode(gin.ReleaseMode)
	endpoints := gin.New()

	log.Println("init HTTP middleware")
	endpoints.Use(s.Middleware)

	log.Println("init HTTP endpoints and handlers")
	endpoints.POST("/short", s.GenerateShortUrl)
	endpoints.GET("/:code", s.Redirect)

	// Start server
	go func() {
		log.Printf("server start listen and serve at '%s:%s'", s.Config.Host, s.Config.Port)
		var err = endpoints.Run(fmt.Sprintf("%s:%s", s.Config.Host, s.Config.Port))
		if err != nil {
			log.Println(err.Error())
		}
	}()

	// Pretty exit
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
	<-sigC

	log.Println("server stop listening")
}

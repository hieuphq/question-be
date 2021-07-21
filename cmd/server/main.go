package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hieuphq/question-be/pkg/config"
	"github.com/hieuphq/question-be/pkg/handler"
	"github.com/hieuphq/question-be/pkg/service/repo"
	"github.com/hieuphq/question-be/pkg/service/repo/pg"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// App api app instance
type App struct {
	l     *log.Logger
	cfg   config.Config
	store repo.Store
}

func main() {
	cfg := config.LoadConfig()
	s, close := pg.NewPostgresStore(&cfg)
	defer close()

	a := App{
		l:     log.Default(),
		cfg:   cfg,
		store: s,
	}

	router := a.setupRouter()

	quit := make(chan os.Signal)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", a.cfg.Port),
		Handler: router,
	}

	go func() {
		// service connections
		a.l.Println("listening on ", a.cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
		quit <- os.Interrupt
	}()

	select {
	case <-quit:

		a.l.Println("Shutdown Server ...")
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			a.l.Println("Server Shutdown:", err)
		}
		a.l.Println("Server exiting")
	}
}

func (a *App) setupRouter() *gin.Engine {
	r := gin.New()

	r.Use(cors.New(
		cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
			AllowHeaders: []string{"Origin", "Host",
				"Content-Type", "Content-Length",
				"Accept-Encoding", "Accept-Language", "Accept",
				"X-CSRF-Token", "Authorization", "X-Requested-With", "X-Access-Token"},
			ExposeHeaders:    []string{"MeAllowMethodsntent-Length"},
			AllowCredentials: true,
		},
	))

	h := handler.New(a.cfg, a.store)

	// handlers
	r.GET("/healthz", h.Healthz)
	r.POST("/topics", h.CreateTopic)
	r.GET("/topics/:code", h.GetTopic)
	return r
}

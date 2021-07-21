package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hieuphq/question-be/pkg/config"
	"github.com/hieuphq/question-be/pkg/service/repo"
	"github.com/hieuphq/question-be/pkg/service/repo/pg"
)

// Handler for app
type Handler struct {
	log   *log.Logger
	cfg   config.Config
	repo  repo.Repo
	store repo.Store
}

// New will return an instance of Auth struct
func New(cfg config.Config, s repo.Store) *Handler {

	r := pg.NewRepo()

	return &Handler{
		log:   log.Default(),
		cfg:   cfg,
		store: s,
		repo:  r,
	}
}

// Healthz handler
// Return "OK"
func (h *Handler) Healthz(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("OK"))

}

// GetHoster handler
// Return "Hoster"
func (h *Handler) GetHoster(c *gin.Context) {
	c.Header("Content-Type", "text/plain")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte("Hoster is anonymous"))
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hieuphq/question-be/pkg/model"
	"github.com/hieuphq/question-be/pkg/util"
)

type newTopicRequest struct {
	Name string `json:"name" binding:"required"`
}

type newTopicResponse = model.Topic

type topicsResponse struct {
	Data []model.Topic `json:"data"`
}

// CreateTopic ...
func (h *Handler) CreateTopic(c *gin.Context) {
	r := newTopicRequest{}
	err := c.ShouldBindJSON(&r)
	if err != nil {
		h.log.Println("Cannot parse upsert mobile device request body", err)
		util.HandleError(c, err)
		return
	}

	topic, err := h.repo.CreateTopic(h.store, model.Topic{
		Name:   r.Name,
		Status: "active",
		Code:   util.RandomString(4),
	})
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, topic)
}

// GetTopic ...
func (h *Handler) GetTopic(c *gin.Context) {
	topicCode := c.Param("code")

	topic, err := h.repo.GetTopicByCode(h.store, topicCode)
	if err != nil {
		util.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, topic)
}

package repo

import (
	"github.com/hieuphq/question-be/pkg/model"
)

// Repo is persistent interface
type Repo interface {
	CreateTopic(store Store, t model.Topic) (*model.Topic, error)
	GetTopicByCode(store Store, code string) (*model.Topic, error)
}

package pg

import (
	"github.com/hieuphq/question-be/pkg/model"
	"github.com/hieuphq/question-be/pkg/service/repo"
)

type pgRepo struct {
}

// NewRepo new pg repo implimentation
func NewRepo() repo.Repo {
	return &pgRepo{}
}

func (r *pgRepo) CreateTopic(store repo.Store, t model.Topic) (*model.Topic, error) {
	return &t, store.DB().Create(&t).Error
}

func (r *pgRepo) GetTopicByCode(store repo.Store, code string) (*model.Topic, error) {
	t := model.Topic{}
	return &t, store.DB().Where("code=?", code).First(&t).Error
}

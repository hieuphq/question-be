package pg

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/k0kubun/pp"

	"github.com/hieuphq/question-be/pkg/config"
	"github.com/hieuphq/question-be/pkg/model/errors"
	"github.com/hieuphq/question-be/pkg/service/repo"
	"github.com/hieuphq/question-be/pkg/util"
)

// store is implimentation of repository
type store struct {
	database *gorm.DB
}

// DB database connection
func (s *store) DB() *gorm.DB {
	return s.database
}

// NewTransaction for database connection
func (s *store) NewTransaction() (newRepo repo.Store, finallyFn repo.FinallyFunc) {
	newDB := s.database.Begin()

	finallyFn = func(err error) error {
		if err != nil {
			nErr := newDB.Rollback().Error
			if nErr != nil {
				return errors.NewStringError(nErr.Error(), http.StatusInternalServerError)
			}
			return errors.NewStringError(err.Error(), util.ParseErrorCode(err))
		}

		cErr := newDB.Commit().Error
		if cErr != nil {
			return errors.NewStringError(cErr.Error(), http.StatusInternalServerError)
		}
		return nil
	}

	return &store{database: newDB}, finallyFn
}

// NewPostgresStore postgres init by gorm
func NewPostgresStore(cfg *config.Config) (repo.Store, func() error) {
	ds := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPass,
		cfg.DBHost, cfg.DBPort, cfg.DBName,
	)
	pp.Println(ds)
	db, err := gorm.Open("postgres", ds)
	if err != nil {
		panic(err)
	}

	return &store{database: db}, db.Close
}

// NewStore postgres init by gorm
func NewStore(db *gorm.DB) repo.Store {
	return &store{database: db}
}

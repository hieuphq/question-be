package repo

import "github.com/jinzhu/gorm"

// FinallyFunc function to finish a transaction
type FinallyFunc = func(error) error

// Store persistent data interface
type Store interface {
	DB() *gorm.DB
	NewTransaction() (Store, FinallyFunc)
}

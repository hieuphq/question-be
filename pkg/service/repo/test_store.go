package repo

import "github.com/jinzhu/gorm"

// NewTestStore .
func NewTestStore() Store {
	return &testStore{}
}

type testStore struct {
}

func (*testStore) DB() *gorm.DB {
	return nil
}

func (s *testStore) NewTransaction() (Store, FinallyFunc) {
	return s, func(err error) error { return err }
}

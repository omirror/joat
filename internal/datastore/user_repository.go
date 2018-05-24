package datastore

import (
	"github.com/asdine/storm"
)

func NewUserRepository(db Bucket) UserRepository {
	return &userRepository{db: db}
}

type UserRepository interface {
}

type userRepository struct {
	db storm.Node
}

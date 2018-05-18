package datastore

import (
	"github.com/asdine/storm"
)

func newAuthRepository(db Bucket) AuthRepository {
	return &repository{db: db}
}

type AuthRepository interface {
}

type repository struct {
	db storm.Node
}

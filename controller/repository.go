package controller

import (
	"github.com/asdine/storm"
)

func newRepository(db storm.Node) Repository {
	return &repository{db: db}
}

type Repository interface {
}

type repository struct {
	db storm.Node
}

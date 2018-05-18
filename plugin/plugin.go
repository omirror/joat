package plugin

import (
	"github.com/asdine/storm"
	"github.com/rs/zerolog"
)

type Plugin interface {
	BundleID() string
	Initialize(db storm.Node, logger *zerolog.Logger)
}

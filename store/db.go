package store

import (
	"fmt"
	"path/filepath"

	"github.com/asdine/storm"
	"github.com/rs/zerolog/log"
	"github.com/ubiqueworks/joat/util"
	"github.com/ubiqueworks/joat/store/codec"
)

type Config struct {
	DataDir string
}

func NewStore(conf *Config) (Store, error) {
	stormStore := &stormStore{dataDir:conf.DataDir}
	if err := stormStore.initialize(); err != nil {
		return nil, err
	}
	return stormStore, nil
}

type Store interface {
	Bucket(name string) storm.Node
}

type stormStore struct {
	dataDir string
	db      *storm.DB
}

func (s *stormStore) initialize() error {
	dbPath := filepath.Join(s.dataDir, "db", "joat.db")
	if err := util.EnsurePath(dbPath, true); err != nil {
		log.Error().Err(err).Msgf("error creating database directory")
		return fmt.Errorf("error creating database directory: %s", filepath.Dir(dbPath))
	}

	if db, err := storm.Open(dbPath, storm.Codec(codec.JSON)); err != nil {
		return err
	} else {
		s.db = db
	}
	return nil
}

func (s *stormStore) Bucket(name string) storm.Node {
	return s.db.From(name)
}

package datastore

import (
	"fmt"
	"path/filepath"

	"github.com/asdine/storm"
	"github.com/rs/zerolog/log"
	"github.com/ubiqueworks/joat/datastore/codec"
	"github.com/ubiqueworks/joat/util"
)

type Config struct {
	DataDir string
}

type Bucket interface {
	storm.Node
}

func NewStore(conf *Config) (Store, error) {
	stormStore := &stormStore{dataDir: conf.DataDir}
	if err := stormStore.initialize(); err != nil {
		return nil, err
	}
	return stormStore, nil
}

type Store interface {
	Bucket(name string) Bucket
}

type stormStore struct {
	dataDir string
	db      *storm.DB
}

func (s *stormStore) initialize() error {
	dbPath := filepath.Join(s.dataDir, "db", "joat.db")
	log.Info().Msgf("Initializing datastore: %s", dbPath)

	if err := util.EnsurePath(dbPath, false); err != nil {
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

func (s *stormStore) Bucket(name string) Bucket {
	return s.db.From(name)
}

/*
	What time doors opens?
	Is it allowed to bring food/beverages?
*/

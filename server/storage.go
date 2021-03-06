package server

import (
	"math/rand"

	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

const (
	maxRandomPool = 50
)

// NewStorage creates instance of the storage.
func NewStorage(db *leveldb.DB) Storage {
	return Storage{db: db}
}

// Storage manages records.
type Storage struct {
	db *leveldb.DB
}

// Add stores record using specified topic.
func (s Storage) Add(topic string, record enr.Record) (string, error) {
	key := make([]byte, 0, len([]byte(topic))+len(record.NodeAddr()))
	key = append(key, []byte(topic)...)
	key = append(key, record.NodeAddr()...)
	data, err := rlp.EncodeToBytes(record)
	if err != nil {
		return "", err
	}
	return string(key), s.db.Put(key, data, nil)
}

// RemoveBykey removes record from storage.
func (s *Storage) RemoveByKey(key string) error {
	return s.db.Delete([]byte(key), nil)
}

// GetRandom reads random records for specified topic up to specified limit.
func (s *Storage) GetRandom(topic string, limit uint) (rst []enr.Record, err error) {
	iter := s.db.NewIterator(util.BytesPrefix([]byte(topic)), nil)
	defer iter.Release()
	pool := make([]enr.Record, 0, maxRandomPool)
	for iter.Next() || len(pool) == maxRandomPool {
		var record enr.Record
		if err = rlp.DecodeBytes(iter.Value(), &record); err != nil {
			return
		}
		pool = append(pool, record)
	}
	if limit >= uint(len(pool)) {
		return pool, nil
	}
	chosen := make([]byte, len(pool))
	for uint(len(rst)) < limit {
		n := rand.Intn(len(pool))
		if chosen[n] == 1 {
			continue
		}
		chosen[n] = 1
		rst = append(rst, pool[n])
	}
	return
}

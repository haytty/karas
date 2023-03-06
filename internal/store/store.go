package store

import (
	"fmt"
	"sync"

	"github.com/haytty/karas/internal/util"
)

type Store struct {
	result map[string]string
	mu     sync.Mutex
}

var (
	cachedStore *Store
)

func NewStore() *Store {
	if cachedStore != nil {
		return cachedStore
	}

	m := make(map[string]string)
	cachedStore = &Store{result: m}
	return cachedStore
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	s.result[key] = value
	s.mu.Unlock()
}

func (s *Store) Dump(format string) error {
	switch format {
	case "json":
		return s.jsonDump()
	default:
		return s.jsonDump()
	}
	return nil
}

func (s *Store) jsonDump() error {
	b, err := util.PrettyJSON(s.result)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

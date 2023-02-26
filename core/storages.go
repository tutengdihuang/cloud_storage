package core

import (
	"errors"
	"fmt"
)

// ErrExistsDriver can not register same driver_default twice.
var ErrExistsDriver = errors.New("driver_default already exists")

// Storages store all supported drivers.
type Storages struct {
	storages map[string]Storage
}

// Get driver_default by name.
func (s *Storages) Get(name string) (Storage, error) {
	if s.storages == nil {
		return nil, errors.New("no valid driver_default")
	}

	storage, ok := s.storages[name]
	if !ok {
		return nil, fmt.Errorf("unsupported driver_default: %s", name)
	}

	return storage, nil
}

// Register a new driver_default.
func (s *Storages) Register(name string, storage Storage) error {
	if s.storages == nil {
		s.storages = make(map[string]Storage)
	}

	if _, ok := s.storages[name]; ok {
		return ErrExistsDriver
	}

	s.storages[name] = storage

	return nil
}

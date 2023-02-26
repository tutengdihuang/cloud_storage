package core

import (
	"cloud_storage/drivers"
	"strings"
)

// Kernel is the core struct of driver_default, it plays the role of a driver_default manager.
type Kernel struct {
	StorageConfig *drivers.StorageConfig
	storages      Storages
	Storage       Storage
}

// New create a new instance of Kernel.
func New(scfg *drivers.StorageConfig) Kernel {
	app := Kernel{
		StorageConfig: scfg,
		storages:      Storages{},
	}

	return app
}

// UseDriver is used to switch to the specified driver_default.
func (a *Kernel) UseDriver(driver Driver) error {
	storage, err := a.storages.Get(strings.ToLower(driver.Name()))
	if err != nil {
		return err
	}

	a.Storage = storage

	return nil
}

// RegisterDriver is used to register new driver_default.
func (a *Kernel) RegisterDriver(driver Driver) error {
	storage, err := driver.Storage()
	if err != nil {
		return err
	}

	return a.storages.Register(driver.Name(), storage)
}

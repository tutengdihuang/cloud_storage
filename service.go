package cloud_storage

import (
	"cloud_storage/core"
	"cloud_storage/driver_default"
	"cloud_storage/drivers"
	"errors"
)

// Goss is the wrapper for core.Kernel
type Goss struct {
	core.Kernel
}

// New creates a new instance based on the configuration file pointed to by configPath.
func New(scfg *drivers.StorageConfig) (core.Storage, error) {
	if string(scfg.Driver) == "" {
		return nil, errors.New("Driver is empty. ")
	}
	goss := Goss{
		core.New(scfg),
	}

	//真正实现new存储对象 driver
	driver, err := driver_default.DefaultDriver(scfg)
	if err != nil {
		return nil, err
	}

	//放入map
	err = goss.RegisterDriver(driver)
	if err != nil {
		return nil, err
	}

	//从map中取出
	err = goss.UseDriver(driver)
	if err != nil {
		return nil, err
	}

	return goss.Storage, nil
}

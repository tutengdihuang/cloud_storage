package driver_default

import (
	"errors"
	"cloud_storage/drivers/minio"

	"cloud_storage/drivers/s3"

	"cloud_storage/core"
	"cloud_storage/drivers"
	"cloud_storage/drivers/aliyun"
	"cloud_storage/drivers/huawei"
	"cloud_storage/drivers/qiniu"
	"cloud_storage/drivers/tencent"
)

var (
	// ErrNoDefaultDriver no default driver_default configured error.
	ErrNoDefaultDriver = errors.New("no default driver_default set")

	// ErrDriverNotExists driver_default not registered error.
	ErrDriverNotExists = errors.New("driver_default not exists")
)

// defaultDriver get the driver_default specified by "driver_default" in the configuration file.
func DefaultDriver(scfg *drivers.StorageConfig) (core.Driver, error) {

	switch scfg.Driver {
	case drivers.Aliyun:
		return aliyun.NewDriver(scfg), nil
	case drivers.Tencent:
		return tencent.NewDriver(scfg), nil
	case drivers.Qiniu:
		return qiniu.NewDriver(scfg), nil
	case drivers.Huawei:
		return huawei.NewDriver(scfg), nil
	case drivers.S3:
		return s3.NewDriver(scfg), nil
	case drivers.Minio:
		return minio.NewDriver(scfg), nil
	default:
		return nil, ErrDriverNotExists
	}
}

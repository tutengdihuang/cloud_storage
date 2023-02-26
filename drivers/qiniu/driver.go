package qiniu

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"cloud_storage/core"
	"cloud_storage/drivers"
)

type Driver struct {
	cfg *drivers.StorageConfig
}

func NewDriver(scfg *drivers.StorageConfig) core.Driver {
	return &Driver{cfg: scfg}
}

func (d *Driver) Storage() (core.Storage, error) {

	if d.cfg.Bucket == "" || d.cfg.AccessKeyID == "" || d.cfg.AccessKeySecret == "" {
		return nil, core.ErrorConfigEmpty
	}

	mac := qbox.NewMac(d.cfg.AccessKeyID, d.cfg.AccessKeySecret)
	cfg := storage.Config{
		UseHTTPS: true,
	}

	store := Store{
		config:        d.cfg,
		bucketManager: storage.NewBucketManager(mac, &cfg),
	}

	return core.NewStorage(&store), nil
}

func (d Driver) Name() string {
	return "qiniu"
}

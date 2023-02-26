package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
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

	if d.cfg.Bucket == "" || d.cfg.AccessKeyID == "" || d.cfg.AccessKeySecret == "" || d.cfg.Endpoint == "" {
		return nil, core.ErrorConfigEmpty
	}

	minioClient, err := minio.New(d.cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(d.cfg.AccessKeyID, d.cfg.AccessKeySecret, ""),
		Secure: d.cfg.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	store := Store{
		client: minioClient,
		config: d.cfg,
	}

	return core.NewStorage(&store), nil
}

func (d Driver) Name() string {
	return "minio"
}

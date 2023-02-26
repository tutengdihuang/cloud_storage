package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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
	bucket, err := ossBucket(d.cfg)
	if err != nil {
		return nil, err
	}

	store := Store{
		Bucket: bucket,
	}

	return core.NewStorage(&store), nil
}

func ossBucket(conf *drivers.StorageConfig) (*oss.Bucket, error) {

	if conf.Endpoint == "" || conf.AccessKeyID == "" || conf.AccessKeySecret == "" {
		return nil, core.ErrorConfigEmpty
	}

	client, err := oss.New(conf.Endpoint, conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	return client.Bucket(conf.Bucket)
}

func (d Driver) Name() string {
	return "aliyun"
}

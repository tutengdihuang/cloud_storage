package huawei

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
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

	client, err := getClient(d.cfg)
	if err != nil {
		return nil, err
	}

	store := Store{
		client: client,
		config: d.cfg,
	}

	return core.NewStorage(&store), nil
}

func getClient(conf *drivers.StorageConfig) (*obs.ObsClient, error) {

	if conf.Endpoint == "" || conf.Region == "" || conf.Bucket == "" || conf.AccessKeyID == "" || conf.AccessKeySecret == "" {
		return nil, core.ErrorConfigEmpty
	}

	return obs.New(conf.AccessKeyID, conf.AccessKeySecret, conf.Endpoint)
}

func (d Driver) Name() string {
	return "huawei"
}

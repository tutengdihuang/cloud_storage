package tencent

import (
	"cloud_storage/drivers"
	"net/http"
	"net/url"

	"cloud_storage/core"
	"github.com/tencentyun/cos-go-sdk-v5"
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
	u, err := url.Parse(d.cfg.Url)
	if err != nil {
		return nil, err
	}

	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  d.cfg.AccessKeyID,
			SecretKey: d.cfg.AccessKeySecret,
		},
	})

	store := Store{
		client: client,
	}

	return core.NewStorage(&store), nil
}

func (d *Driver) Name() string {
	return "tencent"
}

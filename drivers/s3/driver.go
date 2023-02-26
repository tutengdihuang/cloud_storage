package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

	creds := credentials.NewStaticCredentials(d.cfg.AccessKeyID, d.cfg.AccessKeySecret, "")
	_, err := creds.Get()
	if err != nil {
		return nil, err
	}

	awsConfig := &aws.Config{
		Region:      aws.String(d.cfg.Region),
		Endpoint:    aws.String(d.cfg.Endpoint),
		DisableSSL:  aws.Bool(true),
		Credentials: creds,
	}
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)

	store := Store{
		s3:     svc,
		config: d.cfg,
	}

	return core.NewStorage(&store), nil
}

func (d Driver) Name() string {
	return "s3"
}

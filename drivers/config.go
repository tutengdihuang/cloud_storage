package drivers

type DriverType string

const (
	Aliyun  DriverType= "aliyun"
	Tencent DriverType= "tencent"
	Qiniu   DriverType= "qiniu"
	Huawei  DriverType= "huawei"
	S3      DriverType= "s3"
	Minio   DriverType= "minio"
)

type StorageConfig struct {
	Driver          DriverType
	Region          string
	Endpoint        string
	Url             string //# 腾讯云 bucket 对应的的 url
	Bucket          string
	AccessKeyID     string
	AccessKeySecret string
	UseSSL          bool
	Private         bool
	Domain          string //qiniu
}

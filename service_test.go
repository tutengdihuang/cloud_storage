package cloud_storage

import (
	"cloud_storage/core"
	"cloud_storage/drivers"
	"cloud_storage/utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"path/filepath"
	"testing"
)

var accessKeyID = ""
var accessKeySecret = ""

var (
	storage core.Storage

	bucket *oss.Bucket

	key          = "test/foo.txt"
	testdata     string
	fooPath      string
	localFooPath string
)

func Test_Aliyun(t *testing.T) {
	var cfg = &drivers.StorageConfig{}
	cfg.AccessKeyID = accessKeyID
	cfg.AccessKeySecret = accessKeySecret
	cfg.Endpoint = "https://oss-cn-shanghai.aliyuncs.com"
	cfg.Bucket = "top-demo-sh"
	cfg.Region = "shanghai"
	cfg.Driver = drivers.Aliyun

	storeInstance, err := New(cfg)
	if err != nil {
		panic(err)
	}
	// storage 是云存储对象
	storage = storeInstance

	testdata = filepath.Join(utils.RootDir(), "testdata")
	fooPath = filepath.Join(testdata, "foo.txt")
	localFooPath = filepath.Join(testdata, "foo1.txt")

	f, err := os.Open(fooPath)
	if err != nil {
		t.Fatal(err)
	}

	err = storage.Put(key, f)
	if err != nil {
		panic(err)
	}

}

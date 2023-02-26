//go:build integration

package aliyun

import (
	"cloud_storage/drivers"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"cloud_storage/core"
	"cloud_storage/utils"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	fs "github.com/eleven26/go-filesystem"
	"github.com/stretchr/testify/assert"
)

var (
	storage core.Storage

	bucket *oss.Bucket

	key          = "test/foo.txt"
	testdata     string
	fooPath      string
	localFooPath string
)

func Init() {
	/*
	   oss_key_id: "LTAI5tCSkwcbL531j8wPpqv4"
	   oss_key_secret: "I2nJM73QonXJoSx5IK3LOjqSR1Bd4T"
	   oss_end_point: "oss-cn-shanghai"
	   oss_bucket: "top-demo-sh"
	   oss_host: "https://oss-cn-shanghai.aliyuncs.com"
	   prefix: "image/test-"
	   time_to_live: 1
	*/
	var err error
	var cfg = &drivers.StorageConfig{}
	cfg.AccessKeyID = accessKeyID
	cfg.AccessKeySecret = accessKeySecret
	cfg.Endpoint = "https://oss-cn-shanghai.aliyuncs.com"
	cfg.Bucket = "top-demo-sh"
	cfg.Region = "shanghai"
	cfg.Driver = drivers.Aliyun

	d := NewDriver(cfg)
	storage, err = d.Storage()
	if err != nil {
		log.Fatal(err)
	}
	bucket, err = ossBucket(cfg)
	if err != nil {
		log.Fatal(err)
	}

	testdata = filepath.Join(utils.RootDir(), "testdata")
	fooPath = filepath.Join(testdata, "foo.txt")
	localFooPath = filepath.Join(testdata, "foo1.txt")
}

func setUp(t *testing.T) {
	err := bucket.PutObjectFromFile(key, fooPath)
	if err != nil {
		t.Fatal(err)
	}
}

func tearDown(t *testing.T) {
	deleteLocal(t)
	deleteRemote(t)
}

func deleteRemote(t *testing.T) {
	err := bucket.DeleteObject(key)
	if err != nil {
		t.Fatal(err)
	}
}

func deleteLocal(t *testing.T) {
	exists, _ := fs.Exists(localFooPath)
	if exists {
		err := fs.Delete(localFooPath)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestPut(t *testing.T) {
	Init()
	defer tearDown(t)
	f, err := os.Open(fooPath)
	if err != nil {
		t.Fatal(err)
	}

	err = storage.Put(key, f)
	assert.Nil(t, err)

	exists, err := bucket.IsObjectExist(key)
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestPutFromFile(t *testing.T) {
	Init()
	defer tearDown(t)

	err := storage.PutFromFile(key, fooPath)
	assert.Nil(t, err)

	exists, err := bucket.IsObjectExist(key)
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestGet(t *testing.T) {
	Init()

	setUp(t)
	defer tearDown(t)

	rc, err := storage.Get(key)
	defer func(rc io.ReadCloser) {
		err = rc.Close()
		if err != nil {
			t.Fatal(err)
		}
	}(rc)
	assert.Nil(t, err)

	bs, err := io.ReadAll(rc)
	assert.Nil(t, err)
	assert.Equal(t, string(bs), "foo")

	rc, err = storage.Get(key + "not_exists")
	assert.Nil(t, rc)
	assert.Equal(t, http.StatusNotFound, err.(oss.ServiceError).StatusCode)
}

func TestGetString(t *testing.T) {
	Init()
	setUp(t)
	defer tearDown(t)

	content, err := storage.GetString(key)
	assert.Nil(t, err)
	assert.Equal(t, content, "foo")

	content, err = storage.GetString(key + "not_exists")
	assert.Empty(t, content)
	assert.Equal(t, http.StatusNotFound, err.(oss.ServiceError).StatusCode)
}

func TestGetBytes(t *testing.T) {
	Init()
	setUp(t)
	defer tearDown(t)

	bs, err := storage.GetBytes(key)
	assert.Nil(t, err)
	assert.Equal(t, string(bs), "foo")

	bs, err = storage.GetBytes(key + "not_exists")
	assert.Nil(t, bs)
	assert.Equal(t, http.StatusNotFound, err.(oss.ServiceError).StatusCode)
}

func TestSave(t *testing.T) {
	Init()
	setUp(t)
	defer tearDown(t)
	err := storage.GetToFile(key, localFooPath)
	assert.Nil(t, err)
	assert.Equal(t, "foo", fs.MustGetString(localFooPath))
}

func TestSize(t *testing.T) {
	Init()
	setUp(t)
	defer tearDown(t)

	size, err := storage.Size(key)

	var expectedSize int64 = 3
	assert.Nil(t, err)
	assert.Equal(t, expectedSize, size)

	var s int64 = 0
	size, err = storage.Size(key + "not_exists")
	assert.Equal(t, s, size)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.(oss.ServiceError).StatusCode)
}

func TestDelete(t *testing.T) {
	Init()
	setUp(t)

	err := storage.Delete(key)
	assert.Nil(t, err)

	exists, err := bucket.IsObjectExist(key)
	assert.Nil(t, err)
	assert.False(t, exists)
}

func TestExists(t *testing.T) {
	Init()
	setUp(t)
	defer tearDown(t)

	exists, err := storage.Exists(key)
	assert.Nil(t, err)
	assert.True(t, exists)
}

func TestFiles(t *testing.T) {
	Init()
	setUp(t)
	defer tearDown(t)

	files, err := storage.Files("test/")
	assert.Nil(t, err)
	assert.Len(t, files, 1)

	var expectedSize int64 = 3
	assert.Equal(t, key, files[0].Key())
	assert.Equal(t, expectedSize, files[0].Size())
}

func TestFilesWithMultiPage(t *testing.T) {
	Init()
	// Testdata was prepared before.
	dir := "test_all/"

	files, err := storage.Files(dir)
	assert.Nil(t, err)
	assert.Len(t, files, 200)
}

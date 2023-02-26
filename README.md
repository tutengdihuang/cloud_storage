# cloud_storage

✨ `cloud_storage` 是一个简洁的云存储 golang 库，支持**阿里云**、**腾讯云**、**aws s3**。

[![Go Reference](https://pkg.go.dev/badge/github.com/eleven26/go-filesystem.svg)](https://pkg.go.dev/goss)
[![Go Report Card](https://goreportcard.com/badge/github.com/eleven26/go-filesystem)](https://goreportcard.com/report/goss)
[![Go](https://goss/actions/workflows/go.yml/badge.svg)](https://goss/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/eleven26/goss/branch/main/graph/badge.svg?token=UU4lLD2n4k)](https://codecov.io/gh/eleven26/goss)
[![GitHub license](https://img.shields.io/github/license/eleven26/goss)](https://goss/blob/main/LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/eleven26/goss)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/eleven26/goss)


## ⚙️ 配置

### 阿里云
```golang
	var cfg = &drivers.StorageConfig{}
	cfg.AccessKeyID = accessKeyID
	cfg.AccessKeySecret = accessKeySecret
	cfg.Endpoint = "https://oss-cn-shanghai.aliyuncs.com"
	cfg.Bucket = "top-demo-sh"
	cfg.Region = "top-demo-sh"
    cfg.Driver = drivers.Aliyun
```
### 腾讯云
```golang
	var cfg = &drivers.StorageConfig{}
	cfg.AccessKeyID = accessKeyID
	cfg.AccessKeySecret = accessKeySecret
	cfg.Region = "ap-beijing"
	cfg.Url = "https://dev-allen-test-1301664974.cos.ap-beijing.myqcloud.com"
	cfg.Endpoint = "https://dev-allen-test-1301664974.cos.ap-beijing.myqcloud.com"
	cfg.Bucket = "dev-allen-test-1301664974"
	cfg.Driver = drivers.Tencent
```

### AWS

```golang
	var cfg = &drivers.StorageConfig{}
	cfg.AccessKeyID = accessKeyID
	cfg.AccessKeySecret = accessKeySecret
	cfg.Region = endpoints.ApEast1RegionID
	cfg.Endpoint = "http://s3.ap-east-1.amazonaws.com"
	cfg.Bucket = "ecasebucket"
    cfg.Driver = drivers.S3
```
## 💡 基本用法-案例

```golang
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

	testdata = filepath.Join(utils.RootDir(), "testdata")
	fooPath = filepath.Join(testdata, "foo.txt")
	localFooPath = filepath.Join(testdata, "foo1.txt")

	f, err := os.Open(fooPath)
	if err != nil {
		t.Fatal(err)
	}

    err:=storeInstance.Put(key, f)
	if err!=nil{
		panic(err)
	}
```

## 📚 接口

`goss` 支持以下操作：

- [Put](#Put)
- [PutFromFile](#PutFromFile)
- [Get](#Get)
- [GetString](#GetString)
- [GetBytes](#GetBytes)
- [GetToFile](#GetToFile)
- [Delete](#Delete)
- [Exists](#Exists)
- [Files](#Files)
- [Size](#Size)

### Put

上传文件到云存储。第一个参数是 key，第二个参数是 `io.Reader`。

```go
data := []byte("this is some data stored as a byte slice in Go Lang!")
r := bytes.NewReader(data)
err := storage.Put("test/test.txt", r)
```

### PutFromFile

上传文件到云存储。第一个参数是 key，第二个参数是本地文件路径。

```go
err := storage.PutFromFile("test/test.txt", "/path/to/test.txt")
```

### Get

从云存储获取文件。参数是 key。返回值是 `io.ReadCloser` 和 `error`。

```go
// rc 是 `io.ReadCloser`
rc, err := storage.Get("test/test.txt")
defer rc.Close()

bs, err := io.ReadAll(rc)
fmt.Println(string(bs))
```

### GetString

从云存储获取文件。参数是 key。返回值是 `string` 和 `error`

```go
content, err := storage.GetString("test/test.txt")
fmt.Println(content)
```

### GetBytes

从云存储获取文件。参数是 key。返回值是 `[]byte` 和 `error`

```go
bs, err := storage.GetBytes("test/test.txt")
fmt.Println(string(bs))
```

### GetToFile

下载云存储文件到本地。第一个参数是 key，第二个参数是本地路径。

```go
// 第一个参数是云端路径，第二个参数是本地路径
err := storage.GetToFile("test/test.txt", "/path/to/local")
```

### Delete

删除云存储文件。

```go
err := storage.Delete("test/test.txt")
```

### Exists

判断云存储文件是否存在。

```go
exists, err := storage.Exists("test/test.txt")
```

### Files

根据前缀获取文件列表。

> minio 最多返回 1000 个，其他的有多少返回多少。

```go
exists, err := storage.Files("test/")
```

### Size

获取云存储文件大小。

```go
size, err := storage.Size("test/test.txt")
```

## 参考文档

1. [阿里云对象存储](https://help.aliyun.com/product/31815.html)
2. [腾讯云对象存储](https://cloud.tencent.com/document/product/436)
5. [aws s3](https://docs.aws.amazon.com/sdk-for-go/api/service/s3/)


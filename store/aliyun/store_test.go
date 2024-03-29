package aliyun_test

import (
	"os"
	"testing"

	"github.com/cloud-station/store"
	"github.com/cloud-station/store/aliyun"
	"github.com/stretchr/testify/assert"
)

var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

//在单元测试中需要有具体的实现
var (
	uploader store.Uploader
)

//Aliyun oss store upload 测试用例
func TestUpload(t *testing.T) {
	//使用assert 编写测试用例的断言
	//通过new获取一个断言实例
	should := assert.New(t)

	err := uploader.Upload(BucketName, "test.txt", "store_test.go")
	//如果不用断言库 就需要自己写if err == nil{}
	//封装：assert库 两个对象的比较封装为(断言)
	if should.NoError(err) {
		//没有Erro开启下一个步骤
		t.Log("upload ok")
	}
}

//添加测试err的用例
func TestUploadError(t *testing.T) {
	//使用assert 编写测试用例的断言
	//通过new获取一个断言实例
	should := assert.New(t)

	err := uploader.Upload(BucketName, "test.txt", "store_testxxx.go")
	should.Error(err, "An error is expected but got nil.")

}

//通过init 编写uploader 的实例化逻辑
func init() {

	ali, err := aliyun.NewDefaultAliOssStore()

	if err != nil {
		panic(err)
	}
	uploader = ali

}

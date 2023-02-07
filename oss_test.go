package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

//定一个全局变量用于客户端连接,在包初始化init 加载
var (
	client *oss.Client
)

//定义环境变量
var (
	AccessKey    = os.Getenv("ALI_AK")
	AccessSecret = os.Getenv("ALI_SK")
	OssEndpoint  = os.Getenv("ALI_OSS_ENDPOINT")
	BucketName   = os.Getenv("ALI_BUCKET_NAME")
)

//测试用例
//测试阿里云oss的sDK的基础能力
func TestBucketList(t *testing.T) {

	lsRes, err := client.ListBuckets()
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

//测试TestUploadFile的功能是否正常
func TestUploadFile(t *testing.T) {
	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}
}

//初始化oss-client，等下给其他的所有测试用例使用
func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	if err != nil {
		//如果有报错就让程序奔溃
		panic(err)
		// HandleError(err)
	}
	client = c
}

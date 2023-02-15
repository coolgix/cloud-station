package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
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
	bucket, err := client.Bucket(BucketName) //属于客户端的方法
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}

	//bucket的操作，增删改查
	//PutObjectFromFile 上传文件的操作
	//常见我们文件需要创建一个文件夹
	//云商的ossserver 会根据你的key 路径结构,自动帮你创建目录
	//objectKey 上传到bucket里面对象的名称有路径的
	//mydir/test.go ,oss server 会自动创建一个mydir目录，执行一个mkdir -pv 的操作
	//测试 给予一个具体的路径名称,然后把测试用例文件上传到了mydir下
	err = bucket.PutObjectFromFile("mydir/test.go", "oss_test.go")
	if err != nil {
		// HandleError(err)
		t.Log(err)
	}
}

//初始化oss-client，等下给其他的所有测试用例使用
func init() {
	c, err := oss.New(OssEndpoint, AccessKey, AccessSecret)
	fmt.Println(OssEndpoint)
	if err != nil {
		//如果有报错就让程序奔溃
		panic(err)
		// HandleError(err)
	}
	client = c
}

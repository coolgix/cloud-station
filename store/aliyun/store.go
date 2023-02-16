package aliyun

//面向对象

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloud-station/store"
)

var (
	//接口对象是否实现接口的强制约束
	_ store.Uploader = &AliOssStore{}
)

//AliOssStore 对象的构造函数
//构造一个客户端链接 oss
//endpoint, accesskey, accessSecret string 构造出链接oss的对象
func NewAliOssStore(endpoint, accesskey, accessSecret string) (*AliOssStore, error) {
	c, err := oss.New(endpoint, accesskey, accessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client: c,
	}, nil
}

//需要阿里的客户端使用
type AliOssStore struct {
	client *oss.Client
}

//基于NewAliOssStore 构造函数实现文件上传功能
func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	//2、获取bucket对象
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return nil
	}

	//3、上传文件到这个bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName); err != nil {
		return nil
	}

	//4、输出上传文件的url
	downloadURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60*60*24)
	if err != nil {
		return err
	}
	fmt.Printf("文件下载URL: %s \n", downloadURL)
	fmt.Println("请在1天之内下载.")

	return nil
}

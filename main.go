package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//var (
//	Endpoint = ""
//	AccessKeyId  = ""
//	AccessKeySecret  = ""
//	bucketName = ""
//	uploadFile = ""
//)

//func upload(filepath string) error {
//	client, err := oss.New(endpoint, accessKey, secretKey)
//	if err != nil {
//		return err
//	}
//
//	bucket, err := client.Bucket(bucketName)
//	if err != nil {
//		return err
//	}
//
//	err = bucket.PutObjectFromFile(filepath, filepath)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}

//串联逻辑，程序入口
func main() {
	//if err := upload(uploadFile); err != nil {
	//	fmt.Printf("upload file error, %s", err)
	//	os.Exit(1)
	//}

	//测试云商的功能是否符合预期
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "LTAI5tH96bPQPWaSE3QVki1b", "HXyH6QvVbTdZwfGiw0gXbpl040UsGj")
	if err != nil {
		// HandleError(err)
	}

	lsRes, err := client.ListBuckets()
	if err != nil {
		// HandleError(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
	
}

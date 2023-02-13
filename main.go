package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

//修改这些变量，控制程序的运行逻辑
var (
	//程序内置
	acessKey    = "xx"
	acessSecret = "xx"
	endpoint    = "xx"

	//默认配置
	bucketName = "xx"

	//用户需要传递参数
	//期望用户自己输入（cli/GUI）
	uploadFile = ""

	//help
	help = false
)

//实现文件上传的函数
func upload(file_Path string) error {
	//1、实例化客户端
	client, err := oss.New(endpoint, acessKey, acessSecret)
	if err != nil {
		return err
	}

	//2、获取bucket对象
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil
	}

	//3、上传文件到这个bucket
	return bucket.PutObjectFromFile(file_Path, file_Path)
}

//参数合法性检查
func validata() error {
	//校验三个参数全部不为空
	if endpoint == "" || acessKey == "" || acessSecret == "" {
		return fmt.Errorf("endpoint,access_key,access_secret has one empty")
	}
	//校验上传文件的路径不为空
	if uploadFile == " " {
		return fmt.Errorf("upload file path requile")
	}
	//没文件就返回nil
	return nil
}

//cli参数加载
func loadParams() {
	//设置参数
	flag.BoolVar(&help, "h", false, "打印帮助信息")
	flag.StringVar(&uploadFile, "f", "", "上传文件的名称")
	//通过动作传递参数，parse解析“-f ”
	flag.Parse()
}

//打印使用说明
func usage() {
	//1、打印一些描述信息
	fmt.Fprintf(os.Stderr, `cloud-station version:1.0 
	Usage: cloud-station [-h] -f <upload_file_path>
	Options:
	`)
	//2、参数选项信息,打印当前默认的参数选项
	flag.PrintDefaults()
}

func main() {
	//参数加载
	loadParams()

	//参数验证
	if err := validata(); err != nil {
		fmt.Printf("参数校验异常：%s\n", err)
		os.Exit(1)
	}

	if err := upload(uploadFile); err != nil {
		fmt.Printf("上传文件异常：%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("文件：%s 上传完成\n", uploadFile)

}

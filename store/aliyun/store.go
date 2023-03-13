package aliyun

//面向对象

import (
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/cloud-station/store"
)

var (
	//接口对象是否实现接口的强制约束
	//_ store.Uploader 我不需要这个变量的值 只做变量类型的判断
	//&AliOssStore 这个对象必须满足store.Uploader 这个类型

	// _ store.Uploader = &AliOssStore{} 申明了一个空对象，至少需要一个指针
	//nil 空指针，nil有类型
	//申明一个指针类型
	// a *AliOssStore = nil a是这样一个AliOssStore类型的指针
	//如何把一个nil 转换为 有指定类型的变量
	// a int =16
	// b int64 = int64(a)
	// int64 是类型，a 是指针
	// (int64类型)(值)
	// (*AliOssStore)(nil) //获得类型
	//类型断言
	_ store.Uploader = &AliOssStore{}
)

//抽象出一个结构体，把参数选项传递到结构体里面
type Options struct {
	Endpoint     string
	Accesskey    string
	AccessSecret string
}

func (o *Options) Validate() error {
	//校验参数逻辑
	//校验三个参数全部不为空
	if o.Endpoint == "" || o.Accesskey == "" || o.AccessSecret == "" {
		return fmt.Errorf("endpoint,access_key,access_secret has one empty")
	}
	return nil
}

//抽象出链接客户端的环境变量用于单元测试的init()链接客户端函数
func NewDefaultAliOssStore() (*AliOssStore, error) {
	return NewAliOssStore(&Options{
		Endpoint:     os.Getenv("ALI_OSS_ENDPOINT"),
		Accesskey:    os.Getenv("ALI_AK"),
		AccessSecret: os.Getenv("ALI_SK"),
	})
}

//AliOssStore 对象的构造函数
//构造一个客户端链接 oss
//endpoint, accesskey, accessSecret string 构造出链接oss的对象
func NewAliOssStore(opts *Options) (*AliOssStore, error) {
	//链接客户端前调用下参数校验逻辑
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	//客户端链接逻辑
	c, err := oss.New(opts.Endpoint, opts.Accesskey, opts.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &AliOssStore{
		client:   c,
		listener: NewDefaultProgressListener(),
	}, nil
}

//需要阿里的客户端使用
type AliOssStore struct {
	//阿里云oss client，私有变脸，不运行外部
	client *oss.Client
	//进度条实现,依赖Listener的实现
	listener oss.ProgressListener
}

//基于NewAliOssStore 构造函数实现文件上传功能
func (s *AliOssStore) Upload(bucketName string, objectKey string, fileName string) error {
	//2、获取bucket对象
	bucket, err := s.client.Bucket(bucketName)
	if err != nil {
		return nil
	}

	//3、上传文件到这个bucket
	if err := bucket.PutObjectFromFile(objectKey, fileName, oss.Progress(s.listener)); err != nil {
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

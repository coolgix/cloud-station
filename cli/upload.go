package cli

import (
	"fmt"

	"github.com/cloud-station/store"
	"github.com/cloud-station/store/aliyun"
	"github.com/cloud-station/store/aws"
	"github.com/cloud-station/store/tx"
	"github.com/spf13/cobra"
)

var (
	ossProvider  string
	ossEndpoint  string
	accessKey    string
	accessSecret string
	bucketName   string
	uploadFile   string
)

//定义命令名字
var UploadCmd = &cobra.Command{
	Use:     "upload",
	Long:    "upload 文件上传",
	Short:   "upload 文件上传",
	Example: "upload -f filename",
	//upload的文件上传逻辑.
	RunE: func(cmd *cobra.Command, args []string) error {
		//参数的校验
		var (
			uploader store.Uploader
			err      error
		)
		switch ossProvider {
		case "aliyun":
			uploader, err = aliyun.NewAliOssStore(&aliyun.Options{
				Endpoint:     ossEndpoint,
				Accesskey:    accessKey,
				AccessSecret: accessSecret,
			})
		case "tx":
			uploader = tx.NewTxOssStore()

		case "aws":
			uploader = aws.NewAwsOssStore()

		default:
			return fmt.Errorf("not support oss storage provider")
		}
		if err != nil {
			return nil
		}

		//使用Uploader 来上传文件
		return uploader.Upload(bucketName, uploadFile, uploadFile)
	},
}

//参数的校验

func init() {
	//为upload 命令添加持久化的flag
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider[aliyun/tx/aws]")
	f.StringVarP(&ossEndpoint, "ossEndpoint", "e", "oss-cn-beijing.aliyuncs.com", "oss storage provider endpoint")
	f.StringVarP(&bucketName, "bucket_name", "b", "devcloud-station-a", "oss storage bucket name")
	f.StringVarP(&accessKey, "access_key", "k", "", "oss storage provider ak")
	f.StringVarP(&accessSecret, "access_secret", "s", "", "oss storage provider sk")
	f.StringVarP(&uploadFile, "upload_file", "f", "", "upload filename")
	//作为root命令的子命令存在
	//注册给root命令
	RootCmd.AddCommand(UploadCmd)
}

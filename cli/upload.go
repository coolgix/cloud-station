package cli

import (
	"fmt"

	"github.com/cloud-station/store"
	"github.com/cloud-station/store/aliyun"
	"github.com/cloud-station/store/aws"
	"github.com/spf13/cobra"
)

var (
	ossProvider string
	OssEndpoint string
	access_key string
	access_secret string
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
			aliyun.NewAliOssStore(
				Endpoint: OssEndpoint,


			)
		case "tx":
			uploader = tecenter.NewTxOssStore()

		case "aws":
			uploader = aws.NewAwsOssStore()

		default:
			return fmt.Errorf("not support oss storage provider")
		}
		return nil

	},
}

//参数的校验

func init() {
	//为upload 命令添加持久化的flag
	f := UploadCmd.PersistentFlags()
	f.StringVarP(&ossProvider, "provider", "p", "aliyun", "oss storage provider[aliyun/tx/aws]")
	f.StringVarP(&OssEndpoint, "OssEndpoint", "e", "oss-cn-beijing.aliyuncs.com", "oss storage provider endpoint")

	//作为root命令的子命令存在
	//注册给root命令
	RootCmd.AddCommand(UploadCmd)
}

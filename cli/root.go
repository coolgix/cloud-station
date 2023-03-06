package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

//定义参数
var (
	version bool
)

var RootCmd = &cobra.Command{
	Use:     "cloud-station-cli",
	Long:    "cloud-station-cli 云中转站",
	Short:   "cloud-station-cli 云中转站",
	Example: "cloud-station-cli cmds",
	//rootcmd的逻辑
	RunE: func(cmd *cobra.Command, args []string) error {
		if version {
			fmt.Println("cloud-station-cli v0.0.1")
		}

		return nil
	},
}

func init() {
	//绑定一些参数
	f := RootCmd.PersistentFlags()
	//f.StringVarP(&ossProvider,"provider","p","aliyun", "oss storage provider[aliyun/tx/aws]")
	f.BoolVarP(&version, "version", "v", false, "cloud station 版本信息")

}

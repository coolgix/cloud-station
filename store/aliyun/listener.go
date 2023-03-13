package aliyun

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/schollz/progressbar/v3"
)

//进度条方法

//构造函数
func NewDefaultProgressListener() *ProgressListener {
	return &ProgressListener{
		//初始化bar对象

	}
}

type ProgressListener struct {
	//进度条展示
	bar *progressbar.ProgressBar
}

func (p *ProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		p.bar = progressbar.DefaultBytes(
			event.TotalBytes,
			"文件上传中",
		)
	case oss.TransferDataEvent:
		p.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\n文件上传完成\n")
	case oss.TransferFailedEvent:
		fmt.Printf("\n文件上传失败\n")
	default:
	}

}

package models

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
)

var (
	UploadZlMediaKitMediaServerIp   string
	DownloadZlMediaKitMediaServerIp string
	ZlMediaKitMediaServer           string
	VideoType                       string
	ZlMediaKitDataPath              string
	RoutineNum                      int
)

func init() {
	logs.Info("Init global variable")
	UploadZlMediaKitMediaServerIp, _ = config.String("upload_zl_media_kit_media_server_ip")
	ZlMediaKitDataPath, _ = config.String("zl_media_kit_data_path")
	DownloadZlMediaKitMediaServerIp, _ = config.String("download_zl_media_kit_media_server_ip")

	ZlMediaKitMediaServer, _ = config.String("zl_media_kit_media_server")
	VideoType, _ = config.String("video_type")
	RoutineNum, _ = config.Int("routine_num")
}

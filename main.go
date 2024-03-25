package main

import (
	_ "VideoShow/routers"
	"VideoShow/utils"
	"os"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	if len(os.Args) > 1 {
		utils.LoadVideos()
	} else {
		web.Run()
	}
}

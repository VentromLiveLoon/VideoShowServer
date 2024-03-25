package utils

import (
	"VideoShow/models"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/beego/beego/v2/core/logs"
)

func LoadVideos() {
	if len(os.Args) < 3 {
		fmt.Println("Usage VideoShow scan /path/to/load/directory/")
	} else if os.Args[1] == "scan" {
		vp := new(models.VideoPath)
		Routines := make(chan struct{}, models.RoutineNum)
		var wg sync.WaitGroup

		filepath.Walk(os.Args[2], func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				panic(err.Error())
			}
			if filepath.Ext(path) == ".mp4" || filepath.Ext(path) == ".flv" {

				filename := info.Name()
				// Save name to mysql
				vp.Name = filename
				// Check is saved!
				s := vp.SelectByName()
				if s != "" {
					vp.Id = 0
					s2 := vp.Insert()
					if s2 != "" {
						panic(s2)
					}
					logs.Info("Loaded::", path)
					logs.Info("Cmmand::", "ffmpeg ", " -re", " -i ", path, " -c", " copy", " -f", " flv ", models.UploadZlMediaKitMediaServerIp+models.ZlMediaKitDataPath+filename)
					c := exec.Command("ffmpeg", "-re", "-i", path, "-c", "copy", "-f", "flv", models.UploadZlMediaKitMediaServerIp+models.ZlMediaKitDataPath+filename)

					Routines <- struct{}{}
					wg.Add(1)
					go func() {
						defer wg.Done()
						err2 := c.Run()
						if err2 != nil {
							logs.Error(err2.Error())
						}
						<-Routines
					}()

				} else {
					logs.Info("Video::", filename, "Already exists in zlmedia server and mysql")
				}

			}

			return nil
		})
		logs.Info("Waiting all routines exit...")
		wg.Wait()
	}
}

package tasks

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/TsukiCore/tsuki/INTERX/common"
	"github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/TsukiCore/tsuki/INTERX/types"
)

// CacheDataCheck is a function to check cache data if it's expired.
func CacheDataCheck(rpcAddr string, isLog bool) {
	for {
		err := filepath.Walk(config.GetResponseCacheDir(),
			func(path string, info os.FileInfo, err error) error {
				if _, err := os.Stat(path); os.IsNotExist(err) {
					return nil
				}

				if err != nil {
					return err
				}

				delete := false

				if !info.IsDir() && info.Size() != 0 {
					// check cache json data

					common.Mutex.Lock()
					// check if file or path exists
					if _, err := os.Stat(path); os.IsNotExist(err) {
						return nil
					}
					data, _ := ioutil.ReadFile(path)
					common.Mutex.Unlock()

					result := types.InterxResponse{}
					err := json.Unmarshal([]byte(data), &result)

					if err == nil && result.ExpireAt.Before(time.Now()) && result.Response.Block != common.NodeStatus.Block {
						delete = true
					}
				}

				if path != config.GetResponseCacheDir() && delete {
					if isLog {
						common.GetLogger().Info("[cache] Deleting file: ", path)
					}

					common.Mutex.Lock()
					// check if file or path exists
					if _, err := os.Stat(path); os.IsNotExist(err) {
						return nil
					}
					err := os.Remove(path)
					common.Mutex.Unlock()

					if err != nil {
						if isLog {
							common.GetLogger().Error("[cache] Error deleting file: ", err)
						}
						return err
					}

					return nil
				}

				return nil
			})

		if err != nil {
			log.Println(err)
		}

		time.Sleep(2 * time.Second)
	}
}

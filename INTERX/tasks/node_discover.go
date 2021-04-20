package tasks

import (
	"time"

	"github.com/TsukiCore/tsuki/INTERX/common"
	"github.com/TsukiCore/tsuki/INTERX/config"
	"github.com/TsukiCore/tsuki/INTERX/types"
)

var (
	NodeList []types.KnownAddress
)

func NodeDiscover(isLog bool) {
	for {
		uniqueAddresses := config.LoadUniqueAddresses()
		NodeList = uniqueAddresses
		common.GetLogger().Info(config.Config.AddrBooks)

		common.GetLogger().Info("[node-discover] addresses = ", uniqueAddresses)

		time.Sleep(2 * time.Second)
	}
}

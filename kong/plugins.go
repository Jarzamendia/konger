package kong

import (
	"fmt"
	"log"

	models "github.com/jarzamendia/konger/models"
	"github.com/kevholditch/gokong"
)

//GetPlugins asd
func GetPlugins() []models.PluginInfo {

	var list []models.PluginInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")
	}

	plugins, err := gokong.NewClient(gokong.NewDefaultConfig()).Plugins().List()

	if err != nil {

		log.Panic(err)

	}

	for _, plugin := range plugins.Results {

		p := models.PluginInfo{
			ID:         plugin.Id,
			Name:       plugin.Name,
			Enabled:    plugin.Enabled,
			ServiceID:  gokong.IdToString(plugin.ServiceId),
			RouteID:    gokong.IdToString(plugin.RouteId),
			ConsumerID: gokong.IdToString(plugin.ConsumerId),
			Config:     plugin.Config,
		}

		list = append(list, p)

	}

	return list

}

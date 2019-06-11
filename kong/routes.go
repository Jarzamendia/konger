package kong

import (
	"fmt"
	"log"

	models "github.com/jarzamendia/konger/models"
	"github.com/kevholditch/gokong"
)

//GetRouteByServiceID Retorna uma rota pelo ID.
func GetRouteByServiceID(service models.ServiceInfo) []models.RouteInfo {

	var list []models.RouteInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")
	}

	if err == nil {

		routes, err := gokong.NewClient(gokong.NewDefaultConfig()).Routes().GetRoutesFromServiceId(service.ID)

		if err == nil {

			for _, route := range routes {

				r := models.RouteInfo{
					ID:        *route.Id,
					Hosts:     gokong.StringValueSlice(route.Hosts),
					Paths:     gokong.StringValueSlice(route.Paths),
					ServiceID: string(*route.Service),
				}

				list = append(list, r)

			}

		} else {

			log.Panic(err)

		}

	} else {

		log.Panic(err)

	}

	return list

}

//GetRouteByServiceName Retorna uma rota pelo ID.
func GetRouteByServiceName(service models.ServiceInfo) []models.RouteInfo {

	var list []models.RouteInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")
	}

	if err == nil {

		routes, err := gokong.NewClient(gokong.NewDefaultConfig()).Routes().GetRoutesFromServiceName(service.Name)

		if err == nil {

			for _, route := range routes {

				r := models.RouteInfo{
					ID:        *route.Id,
					Hosts:     gokong.StringValueSlice(route.Hosts),
					Paths:     gokong.StringValueSlice(route.Paths),
					ServiceID: string(*route.Service),
				}

				list = append(list, r)

			}

		} else {

			log.Panic(err)

		}

	} else {

		log.Panic(err)

	}

	return list

}

//CreateRouteByServiceID Cria um Route a partir de um ServiceID.
func CreateRouteByServiceID(route models.RouteInfo) models.RouteInfo {

	var r models.RouteInfo

	routeRequest := &gokong.RouteRequest{
		Protocols:    gokong.StringSlice([]string{route.Protocols}),
		Methods:      gokong.StringSlice([]string{route.Methods}),
		Hosts:        gokong.StringSlice(route.Hosts),
		StripPath:    gokong.Bool(route.StripPath),
		PreserveHost: gokong.Bool(route.PreserveHost),
		Service:      gokong.ToId(route.ServiceID),
		Paths:        gokong.StringSlice(route.Paths),
	}

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")
	}

	if err == nil {

		createdRoute, err := kongClient.Routes().Create(routeRequest)

		if err == nil {

			r = models.RouteInfo{
				ID:        *createdRoute.Id,
				Hosts:     gokong.StringValueSlice(createdRoute.Hosts),
				Paths:     gokong.StringValueSlice(createdRoute.Paths),
				ServiceID: string(*createdRoute.Service),
			}

		} else {

			log.Panic(err)
		}

	} else {

		log.Panic(err)

	}

	return r

}

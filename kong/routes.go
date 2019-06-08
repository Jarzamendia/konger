package kong

import (
	"fmt"
	"log"

	models "github.com/jarzamendia/konger/models"
	"github.com/kevholditch/gokong"
)

//GetRoutes Get all routes.
func GetRoutes() {

	//var list []RouteInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")
	}

	if err == nil {

		query := &gokong.RouteQueryString{
			Offset: 0,
			Size:   100,
		}

		fmt.Println(query)

		routes, err := gokong.NewClient(gokong.NewDefaultConfig()).Routes().List(query)

		fmt.Println("Routes verificadas.")

		if err == nil {

			for _, route := range routes {

				fmt.Println(gokong.StringValueSlice(route.Hosts))

			}

		} else {

			log.Panic(err)

		}

	} else {

		log.Panic(err)

	}

}

//GetRouteByServiceID Retorna uma rota pelo ID.
func GetRouteByServiceID(ID string) []models.RouteInfo {

	var list []models.RouteInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")
	}

	if err == nil {

		routes, err := gokong.NewClient(gokong.NewDefaultConfig()).Routes().GetRoutesFromServiceId(ID)

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
func GetRouteByServiceName(Name string) []models.RouteInfo {

	var list []models.RouteInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")
	}

	if err == nil {

		routes, err := gokong.NewClient(gokong.NewDefaultConfig()).Routes().GetRoutesFromServiceName(Name)

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

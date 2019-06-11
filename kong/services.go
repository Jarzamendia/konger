package kong

import (
	"fmt"
	"log"

	models "github.com/jarzamendia/konger/models"
	"github.com/kevholditch/gokong"
)

//GetServices Retorna todos os services.
func GetServices() []models.ServiceInfo {

	var list []models.ServiceInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if err != nil {

		log.Panic(err)

	}

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	services, err := gokong.NewClient(gokong.NewDefaultConfig()).Services().GetServices(&gokong.ServiceQueryString{
		Size:   1000,
		Offset: 0,
	})

	if err != nil {

		log.Panic(err)

	}

	for _, service := range services {

		routes, err := gokong.NewClient(gokong.NewDefaultConfig()).Routes().GetRoutesFromServiceId(*service.Id)

		if err != nil {

		}

		var routeList []models.RouteInfo

		for _, route := range routes {

			r := models.RouteInfo{
				ID:        *route.Id,
				Hosts:     gokong.StringValueSlice(route.Hosts),
				Paths:     gokong.StringValueSlice(route.Paths),
				ServiceID: string(*route.Service),
			}

			routeList = append(routeList, r)

		}

		s := models.ServiceInfo{
			ID:       *service.Id,
			Name:     *service.Name,
			Hosts:    *service.Host,
			Protocol: *service.Protocol,
			Path:     *service.Path,
			Port:     *service.Port,
			Routes:   routeList,
		}

		list = append(list, s)

	}

	return list

}

//GetServiceByID Retorna um service pelo ID.
func GetServiceByID(service models.ServiceInfo) models.ServiceInfo {

	var s models.ServiceInfo
	var routeList []models.RouteInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	if err == nil {

		result, err := gokong.NewClient(gokong.NewDefaultConfig()).Services().GetServiceById(service.ID)

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

					routeList = append(routeList, r)

				}

			}

			s = models.ServiceInfo{
				ID:       *result.Id,
				Name:     *result.Name,
				Hosts:    *result.Host,
				Protocol: *result.Protocol,
				Path:     *result.Path,
				Port:     *result.Port,
				Routes:   routeList,
			}

		} else {

			log.Panic(err)

		}

	} else {

		log.Panic(err)

	}

	return s

}

//GetServiceByName Retorna um service pelo Nome.
func GetServiceByName(service models.ServiceInfo) models.ServiceInfo {

	var s models.ServiceInfo
	var routeList []models.RouteInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	if err == nil {

		result, err := gokong.NewClient(gokong.NewDefaultConfig()).Services().GetServiceByName(service.Name)

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

					routeList = append(routeList, r)

				}

			}

			s = models.ServiceInfo{
				ID:       *result.Id,
				Name:     *result.Name,
				Hosts:    *result.Host,
				Protocol: *result.Protocol,
				Path:     *result.Path,
				Port:     *result.Port,
				Routes:   routeList,
			}

		} else {

			log.Panic(err)

		}

	} else {

		log.Panic(err)

	}

	return s

}

//CreateService Cria um serviço.
func CreateService(service models.ServiceInfo) models.ServiceInfo {

	var s models.ServiceInfo

	serviceRequest := &gokong.ServiceRequest{
		Name:     gokong.String(service.Name),
		Protocol: gokong.String(service.Protocol),
		Host:     gokong.String(service.Hosts),
	}

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	if err == nil {

		result, err := kongClient.Services().Create(serviceRequest)

		if err == nil {

			s = models.ServiceInfo{
				ID:       *result.Id,
				Name:     *result.Name,
				Hosts:    *result.Host,
				Protocol: *result.Protocol,
				Path:     *result.Path,
				Port:     *result.Port,
			}

		} else {

			log.Panic(err)

		}

	} else {

		log.Panic(err)

	}

	return s

}

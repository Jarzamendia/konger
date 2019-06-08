package kong

import (
	"fmt"
	"log"

	models "github.com/jarzamendia/konger/models"
	"github.com/kevholditch/gokong"
)

//GetConsumers Retorna uma lista de consumers.
func GetConsumers() []models.ConsumersInfo {

	var list []models.ConsumersInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")
	}

	if err == nil {

		consumers, err := kongClient.Consumers().List()

		if err == nil {

			for _, consumer := range consumers.Results {

				p := models.ConsumersInfo{
					ID:       consumer.Id,
					Username: consumer.Username,
				}

				list = append(list, p)

			}

		} else {

			log.Panic(err)

		}

	} else {

		log.Panic(err)

	}

	return list

}

//GetConsumerByID Retorna um Consumer pelo seu ID.
func GetConsumerByID(ID string) models.ConsumersInfo {

	var p models.ConsumersInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	if err == nil {

		consumer, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().GetById(ID)
		if err == nil {

			if consumer != nil {

				p = models.ConsumersInfo{
					ID:       consumer.Id,
					Username: consumer.Username,
				}

			}
		}

	} else {

		log.Panic(err)

	}

	return p

}

//GetConsumerByName Retorna um Consumer pelo seu Name.
func GetConsumerByName(Name string) models.ConsumersInfo {

	var p models.ConsumersInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	if err == nil {

		consumer, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().GetByUsername(Name)

		if err == nil {

			if consumer != nil {

				p = models.ConsumersInfo{
					ID:       consumer.Id,
					Username: consumer.Username,
				}

			}
		}

	} else {

		log.Panic(err)

	}

	return p

}

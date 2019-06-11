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
func GetConsumerByID(consumer models.ConsumersInfo) models.ConsumersInfo {

	var p models.ConsumersInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	if err == nil {

		result, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().GetById(consumer.ID)

		if err == nil {

			if result != nil {

				p = models.ConsumersInfo{
					ID:       result.Id,
					Username: result.Username,
				}

			}
		}

	} else {

		log.Panic(err)

	}

	return p

}

//GetConsumerByName Retorna um Consumer pelo seu Name.
func GetConsumerByName(consumer models.ConsumersInfo) models.ConsumersInfo {

	var p models.ConsumersInfo

	//Open a new connection.
	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	if err == nil {

		result, err := gokong.NewClient(gokong.NewDefaultConfig()).Consumers().GetByUsername(consumer.Username)

		if err == nil {

			if result != nil {

				p = models.ConsumersInfo{
					ID:       result.Id,
					Username: result.Username,
				}

			}
		}

	} else {

		log.Panic(err)

	}

	return p

}

//CreateConsumer Criar um consumer, CustomID "".
func CreateConsumer(consumer models.ConsumersInfo) models.ConsumersInfo {

	var c models.ConsumersInfo

	consumerRequest := &gokong.ConsumerRequest{
		Username: consumer.Username,
		CustomId: "",
	}

	kongClient := gokong.NewClient(gokong.NewDefaultConfig())

	status, err := kongClient.Status().Get()

	if status == nil {

		fmt.Println("Falha na verificação de status.")

		log.Panic("Status failed")

	}

	consumerCreated, err := kongClient.Consumers().Create(consumerRequest)

	if err == nil {

		c = models.ConsumersInfo{
			ID:       consumerCreated.Id,
			Username: consumerCreated.Username,
		}

	} else {

		fmt.Errorf("Err: %s", err)

		//fmt.Println("Falha ao criar Consumer.")

		//log.Panic("Consumer Create failed")

	}

	return c

}

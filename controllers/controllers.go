package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jarzamendia/konger/kong"
	"github.com/jarzamendia/konger/models"
)

//HomeHandler Handler
func HomeHandler(w http.ResponseWriter, r *http.Request) {

}

//ConsumersHandler Handler
func ConsumersHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	consumers := kong.GetConsumers()

	jsonResult, err := json.Marshal(consumers)

	if err == nil {

		w.Write(jsonResult)

	} else {

		log.Panic(err)

	}

}

//ConsumersNameHandler Handler
func ConsumersNameHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var consumer models.ConsumersInfo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&consumer)

	if err != nil {
		panic(err)
	}

	consumer = kong.GetConsumerByName(consumer)

	json.NewEncoder(w).Encode(&consumer)

}

//ConsumersIDHandler Handler
func ConsumersIDHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var consumer models.ConsumersInfo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&consumer)

	if err != nil {
		panic(err)
	}

	consumer = kong.GetConsumerByID(consumer)

	json.NewEncoder(w).Encode(&consumer)

}

//ConsumersCreateNameHandler Handler
func ConsumersCreateNameHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var consumer models.ConsumersInfo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&consumer)

	if err != nil {
		panic(err)
	}

	consumer = kong.CreateConsumer(consumer)

	json.NewEncoder(w).Encode(&consumer)

}

//ServicesHandler Handler
func ServicesHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	services := kong.GetServices()

	jsonResult, err := json.Marshal(services)

	fmt.Println(err)

	if err == nil {

		w.Write(jsonResult)

	} else {

		fmt.Println("Erro ao fazer o parse do Json.")
		log.Panic(err)

	}

}

//ServicesNameHandler models.ServiceInfo / models.ServiceInfo
func ServicesNameHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var service models.ServiceInfo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&service)

	if err != nil {
		panic(err)
	}

	service = kong.GetServiceByName(service)

	json.NewEncoder(w).Encode(&service)

}

//ServicesIDHandler models.ServiceInfo / models.ServiceInfo
func ServicesIDHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var service models.ServiceInfo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&service)

	if err != nil {
		panic(err)
	}

	service = kong.GetServiceByID(service)

	json.NewEncoder(w).Encode(&service)

}

//ServicesCreateHandler models.ServiceInfo / models.ServiceInfo
func ServicesCreateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var service models.ServiceInfo

	_ = json.NewDecoder(r.Body).Decode(&service)

	service = kong.CreateService(service)

	json.NewEncoder(w).Encode(&service)

}

//RoutesNameHandler Handler
func RoutesNameHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var service models.ServiceInfo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&service)

	if err != nil {
		panic(err)
	}

	route := kong.GetRouteByServiceName(service)

	json.NewEncoder(w).Encode(&route)

}

//RoutesIDHandler Handler
func RoutesIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var service models.ServiceInfo

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&service)

	if err != nil {
		panic(err)
	}

	route := kong.GetRouteByServiceID(service)

	json.NewEncoder(w).Encode(&route)

}

//PluginsHandler Handler
func PluginsHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	plugins := kong.GetPlugins()

	jsonResult, err := json.Marshal(plugins)

	if err == nil {

		w.Write(jsonResult)

	} else {

		log.Panic(err)

	}

}

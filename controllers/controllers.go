package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jarzamendia/konger/kong"
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

	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	consumerName := vars["Name"]

	consumer := kong.GetConsumerByName(consumerName)

	jsonResult, err := json.Marshal(consumer)

	if err == nil {

		w.Write(jsonResult)

	} else {

		log.Panic(err)

	}

}

//ConsumersIDHandler Handler
func ConsumersIDHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	consumerID := vars["ID"]

	consumer := kong.GetConsumerByID(consumerID)

	jsonResult, err := json.Marshal(consumer)

	if err == nil {

		w.Write(jsonResult)

	} else {

		log.Panic(err)

	}

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

//ServicesNameHandler Handler
func ServicesNameHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	serviceName := vars["Name"]

	service := kong.GetConsumerByName(serviceName)

	jsonResult, err := json.Marshal(service)

	if err == nil {

		w.Write(jsonResult)

	} else {

		log.Panic(err)

	}

}

//ServicesIDHandler Handler
func ServicesIDHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	serviceID := vars["ID"]

	service := kong.GetServiceByID(serviceID)

	jsonResult, err := json.Marshal(service)

	if err == nil {

		w.Write(jsonResult)

	} else {

		log.Panic(err)

	}

}

//RoutesNameHandler Handler
func RoutesNameHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	routeName := vars["Name"]

	routes := kong.GetRouteByServiceName(routeName)

	jsonResult, err := json.Marshal(routes)

	if err == nil {

		w.Write(jsonResult)

	} else {

		log.Panic(err)

	}

}

//RoutesIDHandler Handler
func RoutesIDHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)

	routeID := vars["ID"]

	routes := kong.GetRouteByServiceID(routeID)

	jsonResult, err := json.Marshal(routes)

	if err == nil {

		w.Write(jsonResult)

	} else {

		log.Panic(err)

	}

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

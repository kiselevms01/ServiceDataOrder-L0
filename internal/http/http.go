package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/kiselevms01/wbProject_L0/internal/service"
)

type Config struct {
	Port int `yaml:"port"`
}

type HttpRouter struct {
	port   int
	server *http.ServeMux
	client *service.Client
}

func NewHttpRouter(config Config, client *service.Client) *HttpRouter {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir("www")))
	httpRouter := HttpRouter{
		port:   config.Port,
		server: serveMux,
		client: client,
	}
	serveMux.HandleFunc("/order", GetOrder(&httpRouter))
	return &httpRouter
}

func (h *HttpRouter) Start() error {
	err := http.ListenAndServe(fmt.Sprintf(":%d", h.port), h.server)
	if err != nil {
		return err
	}
	return nil
}

func GetOrder(h *HttpRouter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//set the content type
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		orderUid := r.URL.Query().Get("order_uid")
		id, err := h.client.GetOrder(orderUid)
		if err != nil {
			warning := make(map[string]string)
			warning["error"] = "The order with the entered ID not found!"
			err = json.NewEncoder(w).Encode(warning)
			if err != nil {
				http.NotFound(w, r)
				log.Printf("error json Encode: %v", err)
				return
			}
			return
		}

		err = json.NewEncoder(w).Encode(id)
		if err != nil {
			log.Printf("error json Encode: %v", err)
			http.NotFound(w, r)
		}
	}
}

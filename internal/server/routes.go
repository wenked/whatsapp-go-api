package server

import (
	"encoding/json"
	"log"
	"net/http"
	"whatsapp-go-api/pkg/wbots"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.helloWorldHandler)
	r.Post("/connect", s.connect)

	return r
}

func (s *Server) connect(w http.ResponseWriter, r *http.Request)  {

	response := make(map[string]string)
	err := wbots.InitSession(nil,"554299488471")

	if err != nil {
		response["message"] = "Error connecting to WhatsApp"
		jsonResp, err := json.Marshal(response)
		if err != nil {
			log.Fatalf("error handling JSON marshal. Err: %v", err)
		}
		w.Write(jsonResp)
		
	}

	response["message"] = "QR Code generated"
	jsonResp, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)

}

func (s *Server) helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	w.Write(jsonResp)
}

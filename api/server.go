package api

import (
	"cars-go/store"
	"cars-go/types"
	"encoding/json"
	"net/http"
)

type Server struct {
	listenPort string
	store      store.CarStore
}

func CreateNewServer(listenPort string, store store.CarStore) *Server {
	server := &Server{
		listenPort: listenPort,
		store:      store,
	}

	server.registerRoutes()
	return server
}

func (s *Server) registerRoutes() {
	http.HandleFunc("/car/create", CreateCarMiddleware(s.CreateCar))
	http.HandleFunc("/car/one/", GetCarMiddleware(s.GetCar))
	http.HandleFunc("/car/all", GetAllCarsMiddleware(s.GetAllCars))
	http.HandleFunc("/car/update/", UpdateCarMiddleware(s.UpdateCar))
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.listenPort, nil)
}

func WriteHeaders(res http.ResponseWriter, statusCode int) {
	res.Header().Set("Content-Type", "application/json")

	status := http.StatusAccepted
	switch statusCode {
	case 201:
		status = http.StatusCreated
	case 200:
		status = http.StatusCreated
		break
	case 400:
		status = http.StatusBadRequest
		break
	case 500:
		status = http.StatusInternalServerError
		break
	}

	res.WriteHeader(status)
}

func SendErrorResponse(res http.ResponseWriter, err error) {
	WriteHeaders(res, 400)
	errRes := types.FailureResponse{
		Status:  400,
		Message: err.Error(),
	}
	json.NewEncoder(res).Encode(errRes)
}

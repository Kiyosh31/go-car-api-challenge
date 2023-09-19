package api

import (
	"cars-go/types"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) CreateCar(res http.ResponseWriter, req *http.Request) {
	var newCar types.Car
	err := json.NewDecoder(req.Body).Decode(&newCar)
	if err != nil {
		SendErrorResponse(res, err)
		return
	}

	createdCar, err := s.store.Create(newCar)
	if err != nil {
		SendErrorResponse(res, err)
		return
	}

	WriteHeaders(res, 201)
	json.NewEncoder(res).Encode(createdCar)
}

func (s *Server) GetCar(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/car/one/")
	idInt, _ := strconv.Atoi(id)

	car, err := s.store.Get(idInt)
	if err != nil {
		SendErrorResponse(res, err)
		return
	}

	WriteHeaders(res, 200)
	json.NewEncoder(res).Encode(car)
}

func (s *Server) GetAllCars(res http.ResponseWriter, req *http.Request) {
	cars, err := s.store.GetAll()
	if err != nil {
		SendErrorResponse(res, err)
		return
	}

	WriteHeaders(res, 200)
	json.NewEncoder(res).Encode(cars)
}

func (s *Server) UpdateCar(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/car/update/")
	idInt, _ := strconv.Atoi(id)

	var carToUpdate types.Car
	err := json.NewDecoder(req.Body).Decode(&carToUpdate)
	if err != nil {
		SendErrorResponse(res, err)
		return
	}

	updatedCar, err := s.store.Update(idInt, carToUpdate)
	if err != nil {
		SendErrorResponse(res, err)
		return
	}

	WriteHeaders(res, 200)
	json.NewEncoder(res).Encode(updatedCar)
}

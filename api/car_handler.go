package api

import (
	"cars-go/types"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func (s *Server) CreateCar(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
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
	} else {
		err := errors.New("Metodo no permitido")
		SendErrorResponse(res, err)
	}
}

func (s *Server) GetCar(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		id := strings.TrimPrefix(req.URL.Path, "/car/one/")
		idInt, _ := strconv.Atoi(id)

		car, err := s.store.Get(idInt)
		if err != nil {
			SendErrorResponse(res, err)
			return
		}

		WriteHeaders(res, 200)
		json.NewEncoder(res).Encode(car)
	} else {
		err := errors.New("Metodo no permitido")
		SendErrorResponse(res, err)
	}
}

func (s *Server) GetAllCars(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		cars, err := s.store.GetAll()
		if err != nil {
			SendErrorResponse(res, err)
			return
		}

		WriteHeaders(res, 200)
		json.NewEncoder(res).Encode(cars)
	} else {
		err := errors.New("Metodo no permitido")
		SendErrorResponse(res, err)
	}
}

func (s *Server) UpdateCar(res http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
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
	} else {
		err := errors.New("Metodo no permitido")
		SendErrorResponse(res, err)
	}
}

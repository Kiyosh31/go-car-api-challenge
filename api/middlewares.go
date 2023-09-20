package api

import (
	"bytes"
	"cars-go/types"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func isCarBodyValid(car io.ReadCloser) error {
	var newCar types.Car
	err := json.NewDecoder(car).Decode(&newCar)
	if err != nil {
		return err
	}

	if newCar.Make == "" {
		err := errors.New("Make prop is missing")
		return err
	} else if newCar.Model == "" {
		err := errors.New("Model prop is missing")
		return err
	} else if newCar.Package == "" {
		err := errors.New("Package prop is missing")
		return err
	} else if newCar.Color == "" {
		err := errors.New("Color prop is missing")
		return err
	} else if newCar.Year == 0 {
		err := errors.New("Year prop is missing")
		return err
	} else if newCar.Category == "" {
		err := errors.New("Category prop is missing")
		return err
	} else if newCar.Mileage == 0 {
		err := errors.New("Mileage prop is missing")
		return err
	} else if newCar.Price == 0 {
		err := errors.New("Price prop is missing")
		return err
	}

	return nil
}

func CreateCarMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			err := errors.New("Method not allowed")
			SendErrorResponse(res, err)
			return
		}

		buf, _ := io.ReadAll(req.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))

		if err := isCarBodyValid(rdr1); err != nil {
			SendErrorResponse(res, err)
			return
		}

		req.Body = rdr2
		next.ServeHTTP(res, req)
	}
}

func GetCarMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			err := errors.New("Method not allowed")
			SendErrorResponse(res, err)
			return
		}

		id := strings.TrimPrefix(req.URL.Path, "/car/one/")

		if id == "" {
			err := errors.New("You must provide an ID")
			SendErrorResponse(res, err)
			return
		}

		next.ServeHTTP(res, req)
	}
}

func GetAllCarsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			err := errors.New("Method not allowed")
			SendErrorResponse(res, err)
			return
		}

		next.ServeHTTP(res, req)
	}
}

func UpdateCarMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			err := errors.New("Method not allowed")
			SendErrorResponse(res, err)
			return
		}

		id := strings.TrimPrefix(req.URL.Path, "/car/update/")

		if id == "" {
			err := errors.New("You must provide an ID")
			SendErrorResponse(res, err)
			return
		}

		buf, _ := io.ReadAll(req.Body)
		rdr1 := io.NopCloser(bytes.NewBuffer(buf))
		rdr2 := io.NopCloser(bytes.NewBuffer(buf))

		if err := isCarBodyValid(rdr1); err != nil {
			SendErrorResponse(res, err)
			return
		}
		req.Body = rdr2

		next.ServeHTTP(res, req)
	}
}

package store

import (
	"cars-go/types"
	"fmt"
)

type CarStore struct{}

var carDatabase = []types.Car{}

func NewCarStore() *CarStore {
	return &CarStore{}
}

func (s *CarStore) Create(car types.Car) (types.Car, error) {
	if len(carDatabase) == 0 {
		car.Id = 1
	} else {
		lastId := carDatabase[len(carDatabase)-1].Id
		car.Id = lastId + 1
	}

	carDatabase = append(carDatabase, *&car)
	return carDatabase[len(carDatabase)-1], nil
}

func (s *CarStore) Get(id int) (types.Car, error) {
	for _, car := range carDatabase {
		if car.Id == id {
			return car, nil
		}
	}

	return types.Car{}, fmt.Errorf("Car not found")
}

func (s *CarStore) GetAll() ([]types.Car, error) {
	if len(carDatabase) == 0 {
		return nil, fmt.Errorf("No elements in DB")
	} else {
		return carDatabase, nil
	}
}

func (s *CarStore) Update(id int, car types.Car) (types.Car, error) {
	carIndex, err := s.getCarPositionInList(id)
	if err != nil {
		return types.Car{}, err
	}

	car.Id = id
	carDatabase[carIndex] = car
	return carDatabase[carIndex], nil
}

func (s *CarStore) getCarPositionInList(id int) (int, error) {
	for index, car := range carDatabase {
		if car.Id == id {
			return index, nil
		}
	}

	return 0, fmt.Errorf("Car not found")
}

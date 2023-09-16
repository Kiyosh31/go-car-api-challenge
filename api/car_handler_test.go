package api

import (
	"cars-go/types"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func (s *Server) TestCreateCar(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(s.CreateCar))
	resp, err := http.Post(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("expected 201 but got: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	expected := types.Car{
		Id:       1,
		Make:     "Toyota",
		Model:    "Prius",
		Package:  "Premium",
		Color:    "white",
		Year:     2023,
		Category: "Sedan",
		Mileage:  100,
		Price:    2456.77,
	}

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	if res != expected {
		t.Errorf("Expected %s but got %s", expected, res)
	}

}

package types

type Car struct {
	Id       int
	Make     string  `json:"make"`
	Model    string  `json:"model"`
	Package  string  `json:"package"`
	Color    string  `json:"color"`
	Year     int     `json:"year"`
	Category string  `json:"category"`
	Mileage  int     `json:"mileage"`
	Price    float32 `json:"price"`
}

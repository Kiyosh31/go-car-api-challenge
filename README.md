# Introduction

This codebase is a challenge to create a very simple REST api about cars, does not involve databases or any other golang framework, just standard go libraries

# Usage

## Pre-requisites

1. Go 1.18.1
2. ability to run Makefiles

## Run

1. Open a terminal in root folder
2. Run command
   ```console
   make dev
   ```

At this point API is ready to receive HTTP requests

# Requests

## [POST] Create a car

Endpoint: `http://localhost:3000/car/create`
Body:

```json
{
  "make": "toyota",
  "model": "prius",
  "package": "Base",
  "color": "white",
  "year": 2023,
  "category": "Sedan",
  "mileage": 10000,
  "price": 256.89
}
```

## [GET] Get a car

Endpoint: `http://localhost:3000/car/one/{id}`

## [GET] Get all cars

Endpoint: `http://localhost:3000/car/all`

## [PUT] Update a car

Endpoint: `http://localhost:3000/car/update/{id}`
Body:

```json
{
  "make": "Ford",
  "model": "Tacoma",
  "package": "Premium",
  "color": "red",
  "year": 2012,
  "category": "pickup",
  "mileage": 10000,
  "price": 25645.89
}
```

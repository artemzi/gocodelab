package api

import "github.com/artemzi/gocodelab/storage"

type (
	Location struct {
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	}
	Payload struct {
		Timestamp int64    `json:"timestamp"`
		DriverID  int      `json:"driver_id"`
		Location  Location `json:"location"`
	}
)

type (
	// Структура для возврата ответа по умолчанию
	DefaultResponse struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	}
	// Для возврата ответа, когда мы запрашиваем водителя
	DriverResponse struct {
		Success bool            `json:"success"`
		Message string          `json:"message"`
		Driver  *storage.Driver `json:"driver"`
	}
	// Для возврата ближайших водителей
	NearestDriverResponse struct {
		Success bool              `json:"success"`
		Message string            `json:"message"`
		Drivers []*storage.Driver `json:"drivers"`
	}
)

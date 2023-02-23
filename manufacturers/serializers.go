package manufacturers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ManufacturersSerializer struct {
	C             *gin.Context
	Manufacturers []Manufacturer
}

type ManufacturerSerializer struct {
	C *gin.Context
	Manufacturer
}

type ManufacturerResponse struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"string"`
	Origin string    `json:"origin"`
}

func (self *ManufacturerSerializer) Response() ManufacturerResponse {
	response := ManufacturerResponse{
		ID:     self.ID,
		Name:   self.Name,
		Origin: self.Origin,
	}
	return response
}

func (self *ManufacturersSerializer) Response() []ManufacturerResponse {
	response := []ManufacturerResponse{}
	for _, manufacturer := range self.Manufacturers {
		serializer := ManufacturerSerializer{self.C, manufacturer}
		response = append(response, serializer.Response())
	}
	return response
}

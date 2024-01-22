package seller

import (
	"GabrielMSosa/crud-api/internal/service"
	"errors"

	"github.com/gin-gonic/gin"
)

var (
	ErrDataEmpty         = errors.New("error data empty")
	ErrInvalidNumber     = errors.New("Invalid telephone numbers")
	ErrInvalidData       = errors.New("Invalid address or Company name")
	ErrCidInvalid        = errors.New("invalid Cid")
	ErrInvalidName       = errors.New("error invalid name data")
	ErrInvalidType       = errors.New("error invalid type data")
	ErrINvalidLocalityId = errors.New("invalid LocalityId")
	ErrLocalityId        = errors.New("Locality not exist")
)

type Seller struct {
	srv service.Service
}

func NewSeller(s service.Service) *Seller {
	return &Seller{srv: s}
}

func (h *Seller) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		list, err := h.srv.GetAll(ctx)

	}
}

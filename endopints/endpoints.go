package endopints

import (
	"github.com/go-kit/kit/endpoint"
	"context"
	"baction/services"
)
type ConRequest struct {
	Msgo string
	Msgt string
}

type ConResponse struct {
	Msg string
}

type DifRequest struct {
	Msgo string
	Msgt string
}

type DifResponse struct {
	Msg string
}

type HealRequest struct {}

type HealResponse struct {
	Status bool
}

type Bendpoints struct {
	ConEndopint endpoint.Endpoint
	DifEndpoint endpoint.Endpoint
	HealEndpoint endpoint.Endpoint
}

func ConEndopint (service services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		conr := request.(ConRequest)
		rep := service.Concat(conr.Msgo , conr.Msgt)
		return ConResponse{Msg: rep} , nil
	}
}
func DifEndpoint (service services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		difr := request.(DifRequest)
		rep := service.Diff(difr.Msgo , difr.Msgt)
		return DifResponse{Msg: rep} , err
	}
}

func HealEndpoint (service services.Services) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		rep := service.Health()
		return HealResponse{Status: rep} , nil
	}
}
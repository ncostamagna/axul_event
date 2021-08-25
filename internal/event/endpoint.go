package event

import (
	"context"
	"fmt"
	"github.com/digitalhouse-dev/dh-kit/response"
)

type (
	StoreReq struct {
		UserID      string `json:"user_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Date        string `json:"date"`
		Time        string `json:"time"`
	}

	GetAllReq struct {
		ID      []string `json:"id"`
		UserID  []string `json:"user_id"`
		Days    int64    `json:"days"`
		Preload string   `json:"preload"`
		Limit   int      `json:"limit"`
		Page    int      `json:"page"`
	}

	GetReq struct {
		ID      string `json:"id"`
		Preload string `json:"preload"`
	}
)

type Controller func(ctx context.Context, request interface{}) (interface{}, error)

//Endpoints struct
type Endpoints struct {
	Get    Controller
	GetAll Controller
	Store  Controller
	Update Controller
	Delete Controller
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Store:  makeStoreEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeGetEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.Success("", nil, nil, nil), nil

	}
}

func makeGetAllEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllReq)
		filters := Filters{ID: req.ID, Days: req.Days}

		events, err := service.GetAll(ctx, filters, 0, 0, req.Preload)
		if err != nil {
			return nil, response.BadRequest(err.Error())
		}

		return response.Success("", events, nil, nil), nil
	}
}

func makeStoreEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StoreReq)
		fmt.Println(req)
		event, err := service.Create(ctx, req.UserID, req.Title, req.Description, req.Date, req.Time)
		if err != nil {
			return nil, response.BadRequest(err.Error())
		}

		return response.Success("success", event, nil, nil), nil
	}
}

func makeUpdateEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.Success("success", nil, nil, nil), nil
	}
}

func makeDeleteEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		return response.Success("", nil, nil, nil), nil
	}
}

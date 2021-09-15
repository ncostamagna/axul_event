package event

import (
	"context"
	"github.com/digitalhouse-dev/dh-kit/response"
)

type (
	StoreReq struct {
		Auth        Authentication
		UserID      string `json:"user_id"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Date        string `json:"date"`
		Time        string `json:"time"`
	}

	GetAllReq struct {
		Auth    Authentication
		ID      []string `json:"id"`
		UserID  []string `json:"user_id"`
		Days    int64    `json:"days"`
		Expired int16    `json:"expired"`
		Preload string   `json:"preload"`
		Limit   int      `json:"limit"`
		Page    int      `json:"page"`
	}

	GetReq struct {
		Auth    Authentication
		ID      string `json:"id"`
		Preload string `json:"preload"`
	}

	Authentication struct {
		ID    string
		Token string
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

		if err := service.authorization(ctx, req.Auth.ID, req.Auth.Token); err != nil {
			return response.BadRequest(err.Error()), nil
		}

		filters := Filters{ID: req.ID, Days: req.Days}
		if req.Expired > 0 {
			filters.Expired = true
		}
		events, err := service.GetAll(ctx, filters, 0, 0, req.Preload)
		if err != nil {
			return response.BadRequest(err.Error()), nil
		}

		return response.Success("", events, nil, nil), nil
	}
}

func makeStoreEndpoint(service Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StoreReq)

		if err := service.authorization(ctx, req.Auth.ID, req.Auth.Token); err != nil {
			return response.BadRequest(err.Error()), nil
		}

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

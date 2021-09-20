package event

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/digitalhouse-dev/dh-kit/logger"
	"github.com/ncostamagna/axul_event/pkg/client"
)

type Filters struct {
	ID      []string
	UserId  []string
	Expired bool
	Days    int64
}

type Service interface {
	Get(ctx context.Context, id, pload string) (*Event, error)
	GetAll(ctx context.Context, filters Filters, offset, limit int, pload string) (*[]Event, error)
	Create(ctx context.Context, userID, title, description, date, time string) (*Event, error)
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
	authorization(ctx context.Context, id, token string) error
}

type service struct {
	repo     Repository
	userTran client.Transport
	logger   logger.Logger
}

//NewService is a service handler
func NewService(repo Repository, userTran client.Transport, logger logger.Logger) Service {
	return &service{
		repo:     repo,
		userTran: userTran,
		logger:   logger,
	}
}

func (s *service) Get(ctx context.Context, id, pload string) (*Event, error) {
	return nil, nil
}

func (s *service) GetAll(ctx context.Context, filters Filters, offset, limit int, pload string) (*[]Event, error) {
	events, err := s.repo.GetAll(ctx, filters, offset, limit)
	if err != nil {
		return nil, s.logger.CatchError(err)
	}

	s.logger.DebugMessage(fmt.Sprintf("Get %d Enrollment", len(*events)))
	return events, nil
}

func (s *service) Create(ctx context.Context, userID, title, description, date, times string) (*Event, error) {
	fmt.Println(date)
	v := strings.Split(date, "/")
	var year, month, day, hour, minute int
	var err error
	if times != "" {
		tS := strings.Split(times, ":")
		hour, err = strconv.Atoi(tS[0])
		if err != nil {
			return nil, errors.New("invalid time format")
		}
		minute, err = strconv.Atoi(tS[1])
		if err != nil {
			return nil, errors.New("invalid time format")
		}
	}
	if len(v) != 3 {
		fmt.Println(len(v))
		return nil, errors.New("invalid date format")
	}
	day, err = strconv.Atoi(v[0])
	if err != nil {
		return nil, errors.New("invalid date format")
	}
	month, err = strconv.Atoi(v[1])
	if err != nil {
		return nil, errors.New("invalid date format")
	}
	year, err = strconv.Atoi(v[2])
	if err != nil {
		return nil, errors.New("invalid date format")
	}
	event := Event{
		UserID:      userID,
		Title:       title,
		Description: description,
		Date:        time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
	}
	fmt.Println(event.Date)
	if err := s.repo.Create(ctx, &event); err != nil {
		return nil, s.logger.CatchError(err)
	}
	s.logger.DebugMessage(fmt.Sprintf("Create %s ProductUser", event.ID))

	return &event, nil

}

func (s *service) Update(ctx context.Context, id string) error {
	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	return nil
}

func (s *service) authorization(ctx context.Context, id, token string) error {
	a, err := s.userTran.GetAuth(id, token)

	if err != nil {
		fmt.Println(err)
		return errors.New("invalid authentication")
	}

	if a < 1 {
		return errors.New("invalid authorization")
	}

	return nil
}

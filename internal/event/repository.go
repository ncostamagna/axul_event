package event

import (
	"context"
	"github.com/digitalhouse-dev/dh-kit/logger"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	GetAll(ctx context.Context, filters Filters, offset, limit int) (*[]Event, error)
	Get(ctx context.Context, id string) (*Event, error)
	Create(ctx context.Context, event *Event) error
	Update(ctx context.Context, id string) error
	Delete(ctx context.Context, id string) error
}

type repo struct {
	db     *gorm.DB
	logger logger.Logger
}

func NewRepository(db *gorm.DB, log logger.Logger) Repository {
	return &repo{db, log}
}

func (r *repo) GetAll(ctx context.Context, filters Filters, offset, limit int) (*[]Event, error) {
	var tx *gorm.DB
	var events []Event
	tx = r.db.WithContext(ctx).Model(&events)
	currentTime := time.Now().UTC()
	first := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, time.UTC)

	if filters.Days != 0 {
		second := first.AddDate(0, 0, int(filters.Days)).Add(time.Hour * 20)
		tx = tx.Order("date").Where("DATE_FORMAT(date,'%Y%m%d') between DATE_FORMAT(?,'%Y%m%d') and DATE_FORMAT(?,'%Y%m%d')", first, second)
	}

	if !filters.Expired {
		tx = tx.Order("date").Where("DATE_FORMAT(date,'%Y%m%d') >= DATE_FORMAT(CURDATE(),'%Y%m%d')")
	}

	result := tx.Order("created_at desc").Find(&events)

	for i := range events {

		bd := time.Date(events[i].Date.Year(), events[i].Date.Month(), events[i].Date.Day(), 0, 0, 0, 0, time.UTC)
		events[i].Days = int64(bd.Sub(first).Hours() / 24)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &events, nil
}

func (r *repo) Get(ctx context.Context, id string) (*Event, error) {
	return nil, nil
}

func (r *repo) Create(ctx context.Context, event *Event) error {
	event.ID = uuid.New().String()
	return r.db.Create(&event).Error
}

func (r *repo) Update(ctx context.Context, id string) error {
	return nil
}

func (r *repo) Delete(ctx context.Context, id string) error {
	return nil
}

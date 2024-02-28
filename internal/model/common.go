package model

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm/clause"
)

type Interface interface {
	QueryAvailable() error
	WriteAvailable() error
	GetGettingQuery() (clause.Expression, error)
}

type TimeMeta struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type ObjectMeta struct {
	ID   uint   `gorm:"primarykey" json:"-"`
	UUID string `json:"uuid" gorm:"column:uuid;<-:create;index:,;unique;size:64"`
}

func (o *ObjectMeta) QueryAvailable() error {
	if o.UUID == "" && o.ID == 0 {
		return errors.New("uuid or id is required")
	}
	return nil
}

func (o *ObjectMeta) WriteAvailable() error {
	if o.UUID == "" {
		return errors.New("uuid is required")
	}
	return nil
}

func (o *ObjectMeta) GetGettingQuery() (clause.Expression, error) {
	if err := o.QueryAvailable(); err != nil {
		return nil, err
	}
	if o.ID != 0 {
		return clause.Eq{Column: "id", Value: o.ID}, nil
	}
	return clause.Eq{Column: "uuid", Value: o.UUID}, nil
}

func (o *ObjectMeta) SetUUID() string {
	o.UUID = uuid.NewV4().String()
	return o.UUID
}

package entities

import (
	"time"

	"github.com/google/uuid"
)

type Interface interface {
	GenerateId()
	SetCreatedAt()
	SetUpdateAt()
	TableName() string
	GetMap() map[string]interface{}
	GetFilterId() map[string]interface{}
}

type Base struct {
	Id       uuid.UUID
	CreateAt time.Time
	UpdateAt time.Time
}

func (b *Base) GenerateId() {
	b.Id = uuid.New()
}

func (b *Base) SetCreatedAt() {
	b.CreateAt = time.Now()
}

func (b *Base) SetUpdateAt() {
	b.UpdateAt = time.Now()
}

func GetTimeFormat() string {
	return time.RFC3339
}

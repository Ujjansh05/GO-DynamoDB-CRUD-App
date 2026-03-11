package entities

import (
	"github.com/google/uuid"
	"time"
)

type Interface interface{
	GenerateID()
	SetCreatedAt()
	SetUpdateAt()
	TableName() string
	GetMap() map[string]interface{}
	GetFilterId() map[string]interface{}
}


type Base struct{
	ID				uuid.UUID		`json:"_id"`
	CreatedAt		time.Time		`json:"createdAt"`
	UpdateAt 		time.Time		`json:"updatedAt"`
}


func (b *Base) GenerateID(){
	b.ID = uuid.New()
}

func(b *Base) SetCreatedAt(){
	b.CreatedAt=time.Now()
}
func(b *Base) SetUpdateAt(){
	b.UpdateAt = time.Now()
}

func GetTimeFormat() string{
	return "2010-01-02T15:04:05-0700"
}
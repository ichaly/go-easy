package base

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Entity struct {
	ID        uint64         `gorm:"primary_key;comment:ID;"`
	State     int8           `gorm:"comment:状态;"`
	Remark    string         `gorm:"comment:备注;"`
	CreatedBy *uint64        `gorm:"comment:创建人;"`
	UpdatedBy *uint64        `gorm:"comment:更新人;"`
	DeletedBy *uint64        `gorm:"comment:删除人;"`
	CreatedAt time.Time      `gorm:"comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:逻辑删除;"`
}

func (my Entity) GetID() uint64 {
	return my.ID
}

func (my Entity) MarshalBinary() ([]byte, error) {
	return json.Marshal(my)
}

func (my *Entity) BeforeCreate(tx *gorm.DB) error {
	if my.ID == 0 {
		if id, err := GenerateID(); err != nil {
			return err
		} else {
			my.ID = id
		}
	}
	return nil
}

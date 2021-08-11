package utils

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	UUID      string         `gorm:"type:varchar(36);primaryKey;not null" json:"uuid"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy sql.NullString `gorm:"type:varchar(36);null;default:NULL" json:"created_by"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedBy sql.NullString `gorm:"type:varchar(36);null;default:NULL" json:"updated_by"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	DeletedBy sql.NullString `gorm:"type:varchar(36);null;default:NULL" json:"deleted_by"`
}

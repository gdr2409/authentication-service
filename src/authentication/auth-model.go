package authentication

import (
	"time"

	"github.com/jinzhu/gorm"
)

// AuthToken model stores user's data
type AuthToken struct {
	gorm.Model
	Token      string    `gorm:"size:60;column:token;NOT NULL;UNIQUE_INDEX"`
	ExpiryAt   time.Time `gorm:"column:expiry_at;NOT NULL"`
	ObjectId   string    `gorm:"size:30;column:object_id;NOT NULL"`
	ObjectType string    `gorm:"size:30;column:object_type;NOT NULL"`
}

// TableName for AuthToken is `authtoken`
func (AuthToken) TableName() string {
	return "authtoken"
}

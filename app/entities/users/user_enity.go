package users

import "time"

type UserEntity struct {
	ID       int64     `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	Updated  time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

// TableName overrides the table name used by User to `users`
func (UserEntity) TableName() string {
	return "users"
}

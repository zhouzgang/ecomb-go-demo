package user

import "time"

type User struct {
	UserId    int64     `gorm:"column:user_id;primaryKey"`
	UserName  string    `gorm:"colum:user_name"`
	NickName  string    `gorm:"colum:nick_name"`
	Password  string    `gorm:"colum:password"`
	Phone     string    `gorm:"colum:phone"`
	Status    int8      `gorm:"colum:status"`
	Gender    int8      `gorm:"colum:gender"`
	City      string    `gorm:"colum:city"`
	CreatedAt time.Time `gorm:"colum:created_at"`
	UpdatedAt time.Time `gorm:"colum:updated_at"`
}

func (User) TableName() string {
	return "t_user"
}

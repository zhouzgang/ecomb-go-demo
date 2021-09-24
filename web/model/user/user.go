package user

type User struct {
	UserId   int64  `gorm:"column:user_id;primaryKey"`
	NickName string `gorm:"colum:nick_name"`
	RealName string `gorm:"colum:real_name"`
	Phone    string `gorm:"colum:phone"`
	CreateAt string `gorm:"colum:create_at"`
}

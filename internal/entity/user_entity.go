package entity

type User struct {
	ID        string    `gorm:"column:id;primary_key;"`
	Password  string    `gorm:"column:password;"`
	Name      string    `gorm:"column:name;"`
	Token     string    `gorm:"column:token;"`
	CreatedAt int64     `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64     `gorm:"column:update_at;autoCreateTime;autoUpdateTime:milli"`
	Contacts  []Contact `gorm:"foreignKet:user_id;references:id"`
}

func (u *User) TableName() string {
	return "users"
}

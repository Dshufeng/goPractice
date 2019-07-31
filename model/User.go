package model

type User struct {
	UID        int64 `gorm:"primary_key"`
	Username   string
	Email      string
	Mobile     string
	Password   string
	Salt       string
	CreateTime int64
}

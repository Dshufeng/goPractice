package model

type People struct {
	UID        int64 `gorm:"primary_key"`
	Username   string
	Email      string
	Mobile     string
	Password   string
	Salt       string
	RegIp      string
	CreateTime int64
}

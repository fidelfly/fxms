package res

import "time"

type User struct {
	Id         int64     `xorm:"autoincr pk id"`
	Code       string    `xorm:"not null unique"`
	Name       string    `xorm:"not null"`
	Email      string    `xorm:"not null unique"`
	Password   string    `xorm:"not null"`
	Avatar     int64     `xorm:"not null"`
	CreateTime time.Time `xorm:"created"`
	UpdateTime time.Time `xorm:"updated"`
}

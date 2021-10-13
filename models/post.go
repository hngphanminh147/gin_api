package models

import "time"

type Post struct {
	Id         uint      `gorm:"primary_key:auto_increment" json:"id"`
	Title      string    `gorm:"type:nvarchar(255)" json:"title"`
	Body       string    `gorm:"type:text" json:"body"`
	Created_at time.Time `gorm:"type:time" json:"created_at"`
	Updated_at time.Time `gorm:"type:time" json:"updated_at"`
	Tag        []*Tag    `gorm:"many2many:post_tag" json:"tag"`
}

package models

type Tag struct {
	Id   uint    `gorm:"primary_key:auto_increasement" json:"id"`
	Name string  `gorm:"type:nvarchar(255)" json:"name"`
	Post []*Post `gorm:"many2many:post_tag" json:"post"`
}

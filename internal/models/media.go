package models

type Media struct {
	ID   int    `gorm:"primaryKey; autoIncrement:true"`
	Link string `gorm:"unique; many2many:media_links;"`
}

package model

type User struct {
	ID        string `gorm:"primaryKey" redis:"id"`
	Firstname string `gorm:"column:firstname" redis:"firstname"`
	Lastname  string `gorm:"column:lastname" redis:"lastname" `
}

package structs

import "time"

type User struct {
	Id        string `form:"Id" binding:"required"`
	Name      string `form:"Name" binding:"required"`
	Cellphone string `form:"Cellphone" binding:"required"`
	Email     string `form:"Email" binding:"required"`
	Password  string `form:"Password" binding:"required"`
}

type Chat struct {
	Id       string
	User     string
	Receiver string
}

type Message[Content Photo | Video | Audio] struct {
	Id        string    `form:"id"`
	Content   Content   `form:"content"`
	Date      string    `form:"date"`
	ChatId    string    `form:"chatid"`
	Timestamp time.Time `form: "timestamp"`
}

type Photo struct {
	URL  string
	Info string
}

type Video struct {
	URL  string
	Info string
}

type Audio struct {
	URL  string
	Info string
}

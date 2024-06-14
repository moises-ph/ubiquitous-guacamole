package structs

type User struct {
	Id        string `form:"id" binding:"required"`
	Name      string `form:"name" binding:"required"`
	Cellphone string `form:"cellphone" binding:"required"`
	Email     string `form:"email" binding:"required"`
	Password  string `form:"password" binding:"required"`
}

type Chat struct {
	Id       string
	User     string
	Receiver string
}

type Message[Content Photo | Video | Audio] struct {
	Id      string  `form:"id"`
	Content Content `form:"content"`
	Date    string  `form:"date"`
	ChatId  string  `form:"chatid"`
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

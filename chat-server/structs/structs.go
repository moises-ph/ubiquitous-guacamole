package structs

type User struct {
	Name      string `json: "name"`
	Cellphone string `json: "cellphone"`
	Email     string `json: "email"`
	Password  string `json: "password"`
}

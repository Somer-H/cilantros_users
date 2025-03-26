package entities

type User struct {
	IdUser int `json:"idUser"`
    Username string `json:"username"`
	Password string `json:"password"`
}
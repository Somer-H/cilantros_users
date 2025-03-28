package entities

type User struct {
	IdUser int `json:"idUser"`
    Username string `json:"username"`
	Password string `json:"password"`
	Role string `json:"role"`
	Gmail string `json:"gmail"`
}
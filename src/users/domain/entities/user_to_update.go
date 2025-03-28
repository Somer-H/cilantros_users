package entities

type UserToUpdate struct {
	Password string `json:"password"`
	Gmail string `json:"gmail"`
}
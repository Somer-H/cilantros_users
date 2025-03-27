package entities

type UserLog struct {
    TokenLog string `json:"tokenLog"`
    Username string `json:"username"`
    ID       int    `json:"id"`
}
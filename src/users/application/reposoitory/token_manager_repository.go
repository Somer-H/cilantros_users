package repository

type TokenManager interface {
    GenerateToken(username string, role string) (string, error)
    ValidateToken(token string) (bool, map[string]interface{}, error)
}
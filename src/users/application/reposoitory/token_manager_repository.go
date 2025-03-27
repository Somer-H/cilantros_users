package repository

type TokenManager interface {
    GenerateToken(userID string) (string, error)
    ValidateToken(token string) (bool, map[string]interface{}, error)
}
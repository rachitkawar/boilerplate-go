package jwt

type TokenMaster struct {
	SecretKey string
}

func (t *TokenMaster) CreateToken() string {
	return "token"
}

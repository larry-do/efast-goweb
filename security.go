package goweb

func GetJwtSecretKey() []byte {
	return config.securityConfig.JwtSecretKey
}
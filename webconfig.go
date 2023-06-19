package goweb

type webConfig struct {
	Port uint `yaml:"port"`
	securityConfig securityConfig `yaml:"security"`
}

type securityConfig struct {
	JwtSecretKey []byte `yaml:"jwt_secret_key"`
}

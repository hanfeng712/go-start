package config

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var C Config = Config{}

func ParseConfig(filePath string) {
	conf.MustLoad(filePath, &C)
}

type Config struct {
	rest.RestConf
	Jwt JwtConfig
}

type JwtConfig struct {
	Secret        string `json:"Secret" yaml:"Secret,omitempty"`
	RSAPrivateKey string `json:"RSAPrivateKey" yaml:"RSAPrivateKey,omitempty"`
	RSAPublicKey  string `json:"RSAPublicKey" yaml:"RSAPublicKey,omitempty"`
	ExpireIn      int    `json:"ExpireIn" yaml:"ExpireIn,omitempty" default:"86400"`
}

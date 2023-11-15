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
	Jwt JwtConfig `toml:"Jwt" yaml:"jwt,omitempty"`
}

type JwtConfig struct {
	Secret        string `toml:"Secret" yaml:"Secret,omitempty"`
	RSAPrivateKey string `toml:"RSAPrivateKey" yaml:"RSAPrivateKey,omitempty"`
	RSAPublicKey  string `toml:"RSAPublicKey" yaml:"RSAPublicKey,omitempty"`
	ExpireIn      int    `toml:"ExpireIn" yaml:"ExpireIn,omitempty" default:"86400"`
}

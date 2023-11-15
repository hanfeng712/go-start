package apigateway

import (
	"fmt"
	"gtstart/config"

	"github.com/zeromicro/go-zero/rest"
)

func RunServer(c *config.Config) {
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	RegisterHandlers(server)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

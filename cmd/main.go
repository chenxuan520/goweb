package main

import (
	"fmt"

	"github.com/chenxuan520/goweb/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	g := gin.Default()
	for _, staticPath := range config.GlobalConfig.StaticPaths {
		g.Static(staticPath.RelativePath, staticPath.RootPath)
	}
	g.Run(config.GlobalConfig.Server.Host + ":" + fmt.Sprintf("%d", config.GlobalConfig.Server.Port))
}

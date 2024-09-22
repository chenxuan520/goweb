package main

import (
	"fmt"
	"log"

	"github.com/chenxuan520/goweb/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	g := gin.Default()
	for _, staticPath := range config.GlobalConfig.StaticPaths {
		g.Static(staticPath.WebPath, staticPath.FilePath)
	}
	// print path of server
	log.Println("Server is running at http://" + config.GlobalConfig.Server.Host + ":" + fmt.Sprintf("%d", config.GlobalConfig.Server.Port))
	g.Run(config.GlobalConfig.Server.Host + ":" + fmt.Sprintf("%d", config.GlobalConfig.Server.Port))
}

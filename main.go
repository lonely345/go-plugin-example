package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lonely345/go-plugin-demo/core"
	"github.com/spf13/viper"
)

type plugin struct {
	core.Plugin
}

const (
	Version     string = "0.1.0"
	Description string = "example description"
	Author      string = "caijunduo"
	Status      bool   = true
)

var (
	Plugin = plugin{}
)

func (p plugin) ConfigHooks(config *viper.Viper) {

	config.Set("plugin.example.version", Version)
	config.Set("plugin.example.description", Description)
	config.Set("plugin.example.author", Author)
	config.Set("plugin.example.status", Status)

	err := config.WriteConfig()
	if err != nil {
		core.Engine.Log.Errorf("[ example插件 ] 配置写入失败，%s", err)
	}
}

func (p plugin) RouterHooks(router *gin.Engine) {
	example := router.Group("/example")
	{
		example.GET("/", func(c *gin.Context) {
			c.String(200, "example plugin get")
		})
	}
}

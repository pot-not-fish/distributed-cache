package main

import (
	"kv-cache/head/pkg/kitex_client"
	"kv-cache/head/pkg/parse"
	"kv-cache/head/pkg/process"
	"kv-cache/head/router"

	"github.com/gin-gonic/gin"
)

func main() {
	process.Init(2)
	kitex_client.Init()
	parse.Init("./config.yaml")

	r := gin.Default()

	router.Register(r)

	r.Run("127.0.0.1:5000")
}

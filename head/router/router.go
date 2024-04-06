package router

import (
	"kv-cache/head/handler/auth"
	"kv-cache/head/handler/group"
	"kv-cache/head/handler/kv"
	"kv-cache/head/handler/node"
	"kv-cache/head/pkg/mw"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// 查看节点个数
	node_route := r.Group("/node")
	node_route.Use(mw.JWTMiddlewareFunc)
	{
		// 展示最多100条node
		node_route.POST("/getall", node.GetAllNode)
	}

	// 键值对操作
	kv_route := r.Group("/key")
	kv_route.Use(mw.JWTMiddlewareFunc)
	{
		kv_route.POST("/get", kv.GetValue)
		// 展示最多100条k-v
		kv_route.POST("/getall", kv.GetAllValue)
		// 通过key设置value
		kv_route.POST("/set", kv.SetValue)
		// 通过key删除value
		kv_route.POST("/del", kv.DelValue)
	}

	// 各个组用于隔离不同的kv
	group_route := r.Group("/group")
	group_route.Use(mw.JWTMiddlewareFunc)
	{
		// 设置组
		group_route.POST("/set", group.SetGroup)
		// 展示最多100个组
		group_route.POST("/getall", group.GetAllGroups)
	}

	// 鉴权模块
	auth_route := r.Group("/auth")
	{
		// 登录
		auth_route.POST("/login", auth.Login)
	}
}

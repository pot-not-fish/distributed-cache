package node

import (
	"kv-cache/head/model/node"
	"kv-cache/head/pkg/process"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllNode(c *gin.Context) {
	var getAllNodeReq node.GetAllNodeRequest
	if err := c.ShouldBind(&getAllNodeReq); err != nil {
		c.JSON(http.StatusOK, node.GetAllNodeResponse{Code: 1, Msg: err.Error()})
	}

	nodes := process.M.GetAll()

	c.JSON(http.StatusOK, &node.GetAllNodeResponse{Code: 0, Msg: "OK", Nodes: nodes})
}

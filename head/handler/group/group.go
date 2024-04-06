package group

import (
	"context"
	"kv-cache/head/model/group"
	"kv-cache/head/pkg/kitex_client"
	"kv-cache/head/pkg/process"
	kitex_node "kv-cache/pkg/kitex_gen/node"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetGroup(c *gin.Context) {
	var (
		setGroupReq group.SetGroupRequest
		err         error
	)
	if err := c.ShouldBind(&setGroupReq); err != nil {
		c.JSON(http.StatusOK, group.SetGroupResponse{Code: 1, Msg: err.Error()})
		return
	}

	var resp *kitex_node.SetGroupResponse
	resp, err = kitex_client.Node1Client.SetGroup(context.Background(), &kitex_node.SetGroupRequest{
		Group:    setGroupReq.Group,
		MaxBytes: 1024,
	})
	if err != nil {
		c.JSON(http.StatusOK, group.SetGroupResponse{Code: 1, Msg: err.Error()})
		return
	}
	if resp.Code == 1 {
		c.JSON(http.StatusOK, group.SetGroupResponse{Code: 1, Msg: "invalid handle"})
	}

	resp, err = kitex_client.Node2Client.SetGroup(context.Background(), &kitex_node.SetGroupRequest{
		Group:    setGroupReq.Group,
		MaxBytes: 1024,
	})
	if err != nil {
		c.JSON(http.StatusOK, group.SetGroupResponse{Code: 1, Msg: err.Error()})
		return
	}
	if resp.Code == 1 {
		c.JSON(http.StatusOK, group.SetGroupResponse{Code: 1, Msg: "invalid handle"})
	}

	process.Group = append(process.Group, setGroupReq.Group)

	c.JSON(http.StatusOK, &group.SetGroupResponse{Code: 0, Msg: "OK"})
}

func GetAllGroups(c *gin.Context) {
	var getAllGroupsReq group.GetAllGroupRequest
	if err := c.ShouldBind(&getAllGroupsReq); err != nil {
		c.JSON(http.StatusOK, group.GetAllGroupResponse{Code: 1, Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &group.GetAllGroupResponse{Code: 0, Msg: "OK", Group: process.Group})
}

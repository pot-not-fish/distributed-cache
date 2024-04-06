package kv

import (
	"context"
	"fmt"
	"kv-cache/head/model/kv"
	"kv-cache/head/pkg/kitex_client"
	"kv-cache/head/pkg/process"
	kitex_node "kv-cache/pkg/kitex_gen/node"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetValue(c *gin.Context) {
	var (
		getValueReq kv.GetValueRequest
		err         error
	)
	if err := c.ShouldBind(&getValueReq); err != nil {
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: err.Error()})
		return
	}

	node := process.M.Get(getValueReq.Key)
	var resp *kitex_node.GetKeyResponse
	fmt.Println(node)
	switch node {
	case "node1":
		resp, err = kitex_client.Node1Client.GetKey(context.Background(), &kitex_node.GetKeyRequest{
			Key:   getValueReq.Key,
			Group: getValueReq.Group,
		})
	case "node2":
		resp, err = kitex_client.Node2Client.GetKey(context.Background(), &kitex_node.GetKeyRequest{
			Key:   getValueReq.Key,
			Group: getValueReq.Group,
		})
	default:
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: "null value"})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: err.Error()})
		return
	}
	if resp.Code == 1 {
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: "invalid handle"})
		return
	}

	c.JSON(http.StatusOK, &kv.GetValueResponse{Code: 0, Msg: "OK", Value: resp.Value})
}

func GetAllValue(c *gin.Context) {
	var (
		getAllValueReq kv.GetAllValueRequest
		err            error
	)
	if err := c.ShouldBind(&getAllValueReq); err != nil {
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: err.Error()})
		return
	}

	var resp *kitex_node.GetAllKeysResponse
	resp, err = kitex_client.Node1Client.GetAll(context.Background(), &kitex_node.GetAllKeysRequest{Group: getAllValueReq.Group})
	if err != nil {
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: err.Error()})
		return
	}
	if resp.Code == 1 {
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: "null value"})
		return
	}
	caches := make([]kv.Cache, 0, len(resp.Caches))
	for i := range resp.Caches {
		caches = append(caches, kv.Cache{Key: resp.Caches[i].Key, Value: resp.Caches[i].Value})
	}
	resp, err = kitex_client.Node2Client.GetAll(context.Background(), &kitex_node.GetAllKeysRequest{Group: getAllValueReq.Group})
	if err != nil {
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: err.Error()})
		return
	}
	if resp.Code == 1 {
		c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 1, Msg: "null value"})
		return
	}
	for i := range resp.Caches {
		caches = append(caches, kv.Cache{Key: resp.Caches[i].Key, Value: resp.Caches[i].Value})
	}

	c.JSON(http.StatusOK, kv.GetAllValueResponse{Code: 0, Msg: "OK", Cache: caches})
}

func DelValue(c *gin.Context) {
	var (
		delValueReq kv.DelValueRequest
		err         error
	)
	if err := c.ShouldBind(&delValueReq); err != nil {
		c.JSON(http.StatusOK, kv.DelValueResponse{Code: 1, Msg: err.Error()})
		return
	}

	node := process.M.Get(delValueReq.Key)
	var resp *kitex_node.DelKeyResponse
	switch node {
	case "node1":
		resp, err = kitex_client.Node1Client.DelKey(context.Background(), &kitex_node.DelKeyRequest{
			Key:   delValueReq.Key,
			Group: delValueReq.Group,
		})
	case "node2":
		resp, err = kitex_client.Node2Client.DelKey(context.Background(), &kitex_node.DelKeyRequest{
			Key:   delValueReq.Key,
			Group: delValueReq.Group,
		})
	default:
		c.JSON(http.StatusOK, kv.DelValueResponse{Code: 1, Msg: "null value"})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, kv.DelValueResponse{Code: 1, Msg: err.Error()})
		return
	}
	if resp.Code == 1 {
		c.JSON(http.StatusOK, kv.DelValueResponse{Code: 1, Msg: "null key"})
		return
	}

	c.JSON(http.StatusOK, kv.DelValueResponse{Code: 0, Msg: "OK"})
}

func SetValue(c *gin.Context) {
	var (
		setValueReq kv.SetValueRequest
		err         error
	)
	if err := c.ShouldBind(&setValueReq); err != nil {
		c.JSON(http.StatusOK, kv.SetValueResponse{Code: 1, Msg: err.Error()})
	}

	node := process.M.Get(setValueReq.Key)
	var resp *kitex_node.SetKeyResponse
	switch node {
	case "node1":
		resp, err = kitex_client.Node1Client.SetKey(context.Background(), &kitex_node.SetKeyRequest{
			Key:   setValueReq.Key,
			Value: setValueReq.Value,
			Group: setValueReq.Group,
		})
	case "node2":
		resp, err = kitex_client.Node2Client.SetKey(context.Background(), &kitex_node.SetKeyRequest{
			Key:   setValueReq.Key,
			Value: setValueReq.Value,
			Group: setValueReq.Group,
		})
	default:
		c.JSON(http.StatusOK, kv.SetValueResponse{Code: 1, Msg: "invalid node"})
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, kv.SetValueResponse{Code: 1, Msg: err.Error()})
		return
	}
	if resp.Code == 1 {
		c.JSON(http.StatusOK, kv.SetValueResponse{Code: 1, Msg: "invalid handle"})
		return
	}

	c.JSON(http.StatusOK, kv.SetValueResponse{Code: 0, Msg: "OK"})
}

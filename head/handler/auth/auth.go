package auth

import (
	"kv-cache/head/model/auth"
	"kv-cache/head/pkg/mw"
	"kv-cache/head/pkg/parse"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginReq auth.LoginRequest
	if err := c.ShouldBind(&loginReq); err != nil {
		c.JSON(http.StatusOK, auth.LoginResponse{Code: 1, Msg: err.Error()})
		return
	}

	if loginReq.Username != parse.ConfigStructure.Admin.Username || loginReq.Password != parse.ConfigStructure.Admin.Password {
		c.JSON(http.StatusOK, auth.LoginResponse{Code: 1, Msg: "invalid username or password"})
		return
	}

	token, err := mw.GenerateJWT()
	if err != nil {
		c.JSON(http.StatusOK, auth.LoginResponse{Code: 1, Msg: err.Error()})
		return
	}

	c.JSON(http.StatusOK, auth.LoginResponse{
		Code:  0,
		Msg:   "OK",
		Token: token,
	})
}

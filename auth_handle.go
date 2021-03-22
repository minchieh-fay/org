package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenInfo struct {
	Token string `json:"token"`
}

type ErrRet struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

// var ret401 = `{"errors":[{"code":"UNAUTHORIZED","message":"authentication required","detail":null}]}`

// @Description 通过basicAuth获取token
// @Tags auth
// @Accept  json
// @Param Authorization header string true "Basic {base64(loginid:password)}"
// @Produce json
// @Success 200 {object} TokenInfo
// @Failure 401 {object} ErrRet
// @Failure 403 {object} ErrRet
// @Router /v1/token [get]
func basicAuth(c *gin.Context) {
	// step.1 提取账号密码
	loginid, password, status := c.Request.BasicAuth()
	if status != true {
		c.JSON(http.StatusUnauthorized, &ErrRet{"authentication required", "Unauthorized"})
		return
	}

	// step.2 获取数据库的用户信息
	user := dbe.getUserByLoginID(loginid)
	if user == nil {
		c.JSON(http.StatusForbidden, &ErrRet{"user not exist", "Unauthorized"})
		return
	}
	var err error

	// step.3 验密
	pwd, err := parsePassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrRet{"parse password failed", err.Error()})
		return
	}
	if pwd != password {
		c.JSON(http.StatusForbidden, &ErrRet{"password error", "Unauthorized"})
		return
	}

	// step.5 生成token
	expireSpan := 3600 * 8
	jwtToken, err := signature(user, expireSpan)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &ErrRet{"signature error", "Unauthorized"})
		return
	}

	c.SetCookie("token", jwtToken, expireSpan, "/", "", false, true)
	c.JSON(http.StatusOK, &TokenInfo{jwtToken})
}

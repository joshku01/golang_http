package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// API [後台]測試Api
// @Summary [後台]測試Api
// @description 測試Golang Swagger Open Api
// @Tags Api
// @Produce  json
// @Success 200 {object} global.APIResult "成功"
// @Router /api/ping [GET]
func ApiRoute(c *gin.Context) {
	var data string
	data = "pong test data json"
	c.JSON(http.StatusOK, data)
}

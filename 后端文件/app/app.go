package app

import (
	"backend/routers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Receive_data(c *gin.Context) map[string]interface{} {
	var jsonData map[string]interface{}
	// 使用 ShouldBindJSON 方法将 JSON 数据绑定到 map 中
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	return jsonData
}
func App_go(c *gin.Context) {
	jsonData := Receive_data(c)
	fmt.Println(jsonData)
	operate_type := jsonData["operate_type"].(string)
	operate := jsonData["operate"].(string)
	switch operate_type {
	case "Operate_documents":
		routers.Operate_documents(operate, c, jsonData)
	case "Operate_users":
		routers.Operate_users(operate, c, jsonData)
	case "Operate_results":
		routers.Operate_results(operate, c, jsonData)
	case "Operate_visualoperations":
		routers.Operate_visualoperations(operate, c, jsonData)
	}
}

package routers

import (
	"backend/gorm"
	"os"

	"github.com/gin-gonic/gin"
)

func Visualoperation_create(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State  string `json:"state"`
		Result string `jsono:"result"`
	}
	c.Request.ParseMultipartForm(10 << 20) // 10 * 1024 * 1024

	filename := jsonData["operationname"].(string) + ".py"
	filePath := "./py/" + filename
	// 使用 os.OpenFile 打开文件，如果文件不存在则创建
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close() // 确保在函数结束时关闭文件

	// 定义要写入的字符串
	content := jsonData["file"].(string)

	// 使用 file.WriteString 方法将字符串写入文件
	if _, err := file.WriteString(content); err != nil {
		panic(err)
	}

	//修改数据库
	result := gorm.Visualoperation_create_sql(jsonData, filePath)

	response := Response{
		State:  "OK",
		Result: result,
	}
	c.JSON(200, response)
}
func Visualoperation_query_all(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string                   `json:"state"`
		Data  []map[string]interface{} `json:"data"`
	}

	jsonMaps := gorm.Visualoperation_query_all_sql(jsonData)
	response := Response{
		State: "OK",
		Data:  jsonMaps,
	}

	c.JSON(200, response)

}
func Visualoperation_delete(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State  string `json:"state"`
		Result string `jsono:"result"`
	}

	//修改数据库
	result := gorm.Visualoperation_delete_sql(jsonData)

	response := Response{
		State:  "OK",
		Result: result,
	}
	c.JSON(200, response)

}
func Visualoperation_update(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State  string `json:"state"`
		Result string `jsono:"result"`
	}

	//修改数据库
	result := gorm.Visualoperation_update_sql(jsonData)

	response := Response{
		State:  "OK",
		Result: result,
	}
	c.JSON(200, response)

}

//选择操作
func Operate_visualoperations(operate string, c *gin.Context, jsonData map[string]interface{}) {

	switch operate {
	case "Visualoperation_create":
		Visualoperation_create(c, jsonData)
	case "Visualoperation_query_all":
		Visualoperation_query_all(c, jsonData)
	case "Visualoperation_delete":
		Visualoperation_delete(c, jsonData)
	case "Visualoperation_update":
		Visualoperation_update(c, jsonData)
	}
}

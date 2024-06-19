package routers

import (
	"backend/gorm"

	"github.com/gin-gonic/gin"
)

// document_create  传type是userid,返回json里有一个字段是state(string默认是OK)返回documentid
func Document_create(c *gin.Context, jsonData map[string]interface{}) {

	type Response struct {
		State      string `json:"state"`
		DocumentID int    `json:"documentID"`
	}

	//修改数据库
	newDocumentIDStr := gorm.Document_create_sql(jsonData)

	response := Response{
		State:      "OK",
		DocumentID: newDocumentIDStr, // 包含 documentID 字段
	}
	c.JSON(200, response)
}

//document_delete 传documentID , userID 返回state
func Document_delete(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string `json:"state"`
	}

	gorm.Document_delete_sql(jsonData)

	response := Response{
		State: "OK",
	}
	c.JSON(200, response)
}

//document_query 传documentID ,userID 返回state和updateTime,path,paragraphs
func Document_query(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State      string `json:"state"`
		UpdateTime string `json:"updateTime"`
	}

	request_jsonData := gorm.Document_query_sql(jsonData)

	response := Response{
		State:      "OK",
		UpdateTime: request_jsonData["updateTime"].(string),
	}
	c.JSON(200, response)
}

func Document_update(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string `json:"state"`
	}
	gorm.Document_update_sql(jsonData)
	response := Response{
		State: "OK",
	}
	c.JSON(200, response)
}

func Document_update_time(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string `json:"state"`
	}

	gorm.Document_update_time_sql(jsonData)
	response := Response{
		State: "OK",
	}
	c.JSON(200, response)
}
func Document_update_is_collected(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string `json:"state"`
	}

	gorm.Document_update_is_collected_sql(jsonData)
	response := Response{
		State: "OK",
	}
	c.JSON(200, response)
}

//Document_query_all传operat,返回所有
func Document_query_all(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string                   `json:"state"`
		Data  []map[string]interface{} `json:"data"`
	}

	jsonMaps := gorm.Document_query_all_sql(jsonData)
	response := Response{
		State: "OK",
		Data:  jsonMaps,
	}

	c.JSON(200, response)
}

func Document_txt(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string `json:"state"`
	}
	gorm.Document_txt_sql(jsonData)
	response := Response{
		State: "OK",
	}
	c.JSON(200, response)
}

//选择操作
func Operate_documents(operate string, c *gin.Context, jsonData map[string]interface{}) {

	switch operate {
	case "Document_query_all":
		Document_query_all(c, jsonData)
	case "Document_create":
		Document_create(c, jsonData)
	case "Document_delete":
		Document_delete(c, jsonData)
	case "Document_txt":
		Document_txt(c, jsonData)
	case "Document_query":
		Document_query(c, jsonData)
	case "Document_update":
		Document_update(c, jsonData)
	case "Document_update_is_collected":
		Document_update_is_collected(c, jsonData)
	}
}

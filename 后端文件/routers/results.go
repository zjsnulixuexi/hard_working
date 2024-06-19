package routers

import (
	"backend/gorm"

	"github.com/gin-gonic/gin"
)

func Result_create(c *gin.Context, jsonData map[string]interface{}) {

	type Response struct {
		State    string `json:"state"`
		Resultid int    `json:"resultid"`
		Result   string `jsono:"result"`
	}

	//修改数据库
	newUseridStr, result := gorm.Result_create_sql(jsonData)

	response := Response{
		State:    "OK",
		Resultid: newUseridStr,
		Result:   result,
	}
	c.JSON(200, response)
}

func Result_delete(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State  string `json:"state"`
		Result string `json:"result"`
	}

	//修改数据库
	result := gorm.Result_delete_sql(jsonData)

	response := Response{
		State:  "OK",
		Result: result,
	}
	c.JSON(200, response)
}

func Result_query_all(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string                   `json:"state"`
		Data  []map[string]interface{} `json:"data"`
	}

	jsonMaps := gorm.Result_query_all_sql(jsonData)
	response := Response{
		State: "OK",
		Data:  jsonMaps,
	}

	c.JSON(200, response)
}

func Operate_results(operate string, c *gin.Context, jsonData map[string]interface{}) {

	switch operate {
	case "result_create":
		Result_create(c, jsonData)
	case "result_delete":
		Result_delete(c, jsonData)
	case "result_query_all":
		Result_query_all(c, jsonData)
	}
}

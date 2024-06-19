package routers

import (
	"backend/gorm"

	"github.com/gin-gonic/gin"
)

//注册用户
func User_registration(c *gin.Context, jsonData map[string]interface{}) {

	type Response struct {
		State  string `json:"state"`
		Userid int    `json:"userid"`
		Result string `jsono:"result"`
	}

	//修改数据库
	newUseridStr, result := gorm.User_registration_sql(jsonData)

	response := Response{
		State:  "OK",
		Userid: newUseridStr,
		Result: result,
	}
	c.JSON(200, response)

}

//注册用户
func User_login(c *gin.Context, jsonData map[string]interface{}) {

	type Response struct {
		State  string `json:"state"`
		Result string `json:"result"`
		Power  int    `json:"power"`
		Userid int    `json:"userid"`
	}

	//修改数据库
	results, power, userid := gorm.User_login_sql(jsonData)

	response := Response{
		State:  "OK",
		Result: results,
		Power:  power,
		Userid: userid,
	}
	c.JSON(200, response)
}

//忘记密码
func Password_recover(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State    string `json:"state"`
		Password string `json:"password"`
	}

	//修改数据库
	password := gorm.Password_recover_sql(jsonData)

	response := Response{
		State:    "OK",
		Password: password,
	}
	c.JSON(200, response)
}

func Manage_user_delete(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State  string `json:"state"`
		Result string `json:"result"`
	}

	//修改数据库
	result := gorm.Manage_user_delete_sql(jsonData)

	response := Response{
		State:  "OK",
		Result: result,
	}
	c.JSON(200, response)
}

func User_query_all(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State string                   `json:"state"`
		Data  []map[string]interface{} `json:"data"`
	}

	jsonMaps := gorm.User_query_all_sql(jsonData)
	response := Response{
		State: "OK",
		Data:  jsonMaps,
	}

	c.JSON(200, response)
}
func User_power_update(c *gin.Context, jsonData map[string]interface{}) {
	type Response struct {
		State     string `json:"state"`
		Result    string `json:"result"`
		New_power int    `json:"new_power"`
	}

	//修改数据库
	result, power := gorm.User_power_update(jsonData)

	response := Response{
		State:     "OK",
		Result:    result,
		New_power: power,
	}
	c.JSON(200, response)
}

func Operate_users(operate string, c *gin.Context, jsonData map[string]interface{}) {

	switch operate {
	case "User_registration":
		User_registration(c, jsonData)
	case "User_login":
		User_login(c, jsonData)
	case "Password_recovery":
		Password_recover(c, jsonData)
	case "Manage_user_delete":
		Manage_user_delete(c, jsonData)
	case "User_query_all":
		User_query_all(c, jsonData)
	case "User_power_update":
		User_power_update(c, jsonData)

	}
}

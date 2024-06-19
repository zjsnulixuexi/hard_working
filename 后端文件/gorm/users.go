package gorm

import (
	"backend/models"
	"encoding/json"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//连接数据库
func Connect_To_Database_b() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("/root/server.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

//userID自增
func Userid_self_increasing(db *gorm.DB) int {
	db.AutoMigrate(&models.Users{})

	var user models.Users
	if err := db.Order("userid DESC").First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//首次创建
			return 1
		} else {
			panic(err)
		}
	}
	return user.Userid + 1
}

func User_registration_sql(jsonData map[string]interface{}) (int, string) {
	db, nil := Connect_To_Database_b()
	db.AutoMigrate(&models.Users{})
	var user models.Users
	db.Where("username = ?", jsonData["username"].(string)).Find(&user)

	if user.Userid == 0 {
		userid := Userid_self_increasing(db)
		doc := models.Users{
			Userid:   userid,
			Password: jsonData["password"].(string),
			Username: jsonData["username"].(string),
			School:   jsonData["school"].(string),
			Sex:      jsonData["sex"].(string),
			Power:    2,
		}
		db.AutoMigrate(&models.Users{})
		result := db.Create(&doc)
		if result.Error != nil {
			panic(result.Error)
		}
		return userid, "注册成功"
	} else {
		return -1, "用户已存在"
	}
}

func User_login_sql(jsonData map[string]interface{}) (string, int, int) {
	if jsonData["username"].(string) == "" {
		return "用户不存在", -1, -1
	}
	db, err := Connect_To_Database_b()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Users{})

	var result models.Users
	if err := db.Where("username = ?", jsonData["username"].(string)).Find(&result).Error; err != nil {
		return "系统错误，请等待修复", -1, -1
	}
	if result.Password == "" {
		return "用户不存在", -1, -1
	}
	if jsonData["password"].(string) != result.Password {
		return "密码错误", -1, -1
	}

	return "登录成功", result.Power, result.Userid
}

func Password_recover_sql(jsonData map[string]interface{}) string {
	db, err := Connect_To_Database_b()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Users{})

	var result models.Users
	if err := db.Where("username = ?", jsonData["username"].(string)).Find(&result).Error; err != nil {
		return "系统错误，请等待修复"
	}
	if result.Password == "" {
		return "用户不存在"
	}

	return result.Password
}

func Manage_user_delete_sql(jsonData map[string]interface{}) string {

	db, err := Connect_To_Database_b()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Users{})

	deleteCondition := models.Users{Userid: int(jsonData["userid"].(float64))}

	result := db.Where(&deleteCondition).Delete(&models.Users{})
	if result.Error != nil {
		return "删除失败"
	}
	return "删除成功"
}

func User_query_all_sql(jsonData map[string]interface{}) []map[string]interface{} {
	db, err := Connect_To_Database_b()

	if err != nil {
		panic(err)
	}
	var users []models.Users
	db.Table("users_view").Find(&users)

	var jsonArray []string
	for _, doc := range users {
		jsonDateConversion, err := json.Marshal(doc)
		if err != nil {
			panic(err)
		}
		jsonArray = append(jsonArray, string(jsonDateConversion))
	}
	// 创建一个空的 []map[string]interface{} 类型的切片，用于存储转换后的数据
	var jsonMaps []map[string]interface{}

	// 遍历每个 JSON 字符串
	for _, jsonString := range jsonArray {
		// 创建一个 map 用于存储解析后的 JSON 数据
		var jsonMap map[string]interface{}

		// 解析 JSON 字符串到 map[string]interface{} 类型
		errs := json.Unmarshal([]byte(jsonString), &jsonMap)
		if errs != nil {
			panic(errs)
		}

		// 将解析后的 map 添加到 jsonMaps 切片中
		jsonMaps = append(jsonMaps, jsonMap)
	}
	return jsonMaps
}

func User_power_update(jsonData map[string]interface{}) (string, int) {
	db, err := Connect_To_Database_b()
	if err != nil {
		return "系统错误", -1
	}
	if int(jsonData["managepower"].(float64)) == 1 {
		return "无权限", -1
	}
	// 更新记录
	if int(jsonData["power"].(float64)) == 2 {
		result := db.Model(&models.Users{}).Where("userid = ?", int(jsonData["userid"].(float64))).Update("power", 1)
		if result.Error != nil {
			return "系统错误", -1
		}
		return "已修改为一级管理员权限", 1
	} else {
		result := db.Model(&models.Users{}).Where("userid = ?", int(jsonData["userid"].(float64))).Update("power", 2)
		if result.Error != nil {
			return "系统错误", -1
		}
		return "已修改为用户权限", 2
	}
}

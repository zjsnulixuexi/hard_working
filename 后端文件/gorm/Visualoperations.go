package gorm

import (
	"backend/models"
	"encoding/json"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect_To_Database_c() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("/root/server.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

//userID自增
func Visualoperationid_self_increasing(db *gorm.DB) int {
	db, err := Connect_To_Database_c()
	if err != nil {
		return -1
	}
	db.AutoMigrate(&models.Visualoperations{})
	var visualoperation models.Visualoperations
	if err := db.Order("operationid DESC").First(&visualoperation).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//首次创建
			return 1
		} else {
			panic(err)
		}
	}
	return visualoperation.Operationid + 1
}
func Visualoperation_create_sql(jsonData map[string]interface{}, filePath string) string {
	db, err := Connect_To_Database_c()
	if err != nil {
		return "系统错误"
	}
	db.AutoMigrate(&models.Visualoperations{})
	operateid := Visualoperationid_self_increasing(db)
	doc := models.Visualoperations{
		Operationid:        operateid,
		Operationname:      jsonData["operationname"].(string),
		Operationintroduce: jsonData["operationintroduce"].(string),
		Operatinopath:      filePath,
	}
	db.AutoMigrate(&models.Visualoperations{})
	result := db.Create(&doc)
	if result.Error != nil {
		return "系统错误"
	}
	return "导入成功"
}

func Visualoperation_query_all_sql(jsonData map[string]interface{}) []map[string]interface{} {
	db, err := Connect_To_Database_c()

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Visualoperations{})
	var visualoperations []models.Visualoperations
	db.Table("visualoperations_view").Find(&visualoperations)
	var jsonArray []string
	for _, doc := range visualoperations {
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

func Visualoperation_delete_sql(jsonData map[string]interface{}) string {
	db, err := Connect_To_Database_b()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Visualoperations{})

	result := db.Where("operationid = ?", int(jsonData["operationid"].(float64))).Delete(&models.Visualoperations{})
	if result.Error != nil {
		return "删除失败"
	}
	return "删除成功"
}

func Visualoperation_update_sql(jsonData map[string]interface{}) string {
	db, err := Connect_To_Database_c()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Visualoperations{})

	// 更新记录
	db.Model(&models.Visualoperations{}).Where("operationid = ?", int(jsonData["operationid"].(float64))).Update("operationname", jsonData["operationname"].(string))
	db.Model(&models.Visualoperations{}).Where("operationid = ?", int(jsonData["operationid"].(float64))).Update("operationintroduce", jsonData["operationintroduce"].(string))
	return "更新成功"
}

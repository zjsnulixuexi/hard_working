package gorm

import (
	"backend/models"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect_To_Database_d() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("/root/server.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Resultid_self_increasing(db *gorm.DB) int {
	db.AutoMigrate(&models.Results{})

	var result models.Results
	if err := db.Order("resultid DESC").First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			//首次创建
			return 1
		} else {
			panic(err)
		}
	}
	return result.Resultid + 1
}

func Result_create_sql(jsonData map[string]interface{}) (int, string) {
	db, nil := Connect_To_Database_d()
	db.AutoMigrate(&models.Results{})
	var text models.Results
	db.Model(&models.Results{}).Where("userid = ? AND documentid = ? AND kind = ?", int(jsonData["userid"].(float64)), int(jsonData["documentid"].(float64)), int(jsonData["kind"].(float64))).Find(&text)
	fmt.Println(text.Resultid)
	if text.Resultid != 0 {
		return -1, "已存在"
	}
	resultid := Resultid_self_increasing(db)

	var visualoperation models.Visualoperations
	db.Model(&models.Visualoperations{}).Where("operationid = ?", int(jsonData["kind"].(float64))).Find(&visualoperation)
	var document models.Documents
	db.Model(&models.Documents{}).Where("userid = ? AND documentid = ?", int(jsonData["userid"].(float64)), int(jsonData["documentid"].(float64))).Find(&document)
	documentid := int(jsonData["documentid"].(float64))

	cmd := exec.Command("/usr/bin/python3", visualoperation.Operatinopath, document.Path)

	output, _ := cmd.CombinedOutput()

	decodedOutput := strings.TrimSpace(string(output))

	doc := models.Results{
		Userid:     int(jsonData["userid"].(float64)),
		Documentid: documentid,
		Resultid:   resultid,
		Resultname: jsonData["resultname"].(string),
		Kind:       int(jsonData["kind"].(float64)),
		Content:    decodedOutput,
	}

	db.AutoMigrate(&models.Results{})
	result := db.Create(&doc)
	if result.Error != nil {
		panic(result.Error)
	}
	return resultid, "结果为"

}

func Result_delete_sql(jsonData map[string]interface{}) string {
	db, err := Connect_To_Database_d()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Results{})

	result := db.Where("userid = ? AND documentid = ? AND resultid = ?", int(jsonData["userid"].(float64)), int(jsonData["documentid"].(float64)), int(jsonData["resultid"].(float64))).Delete(&models.Results{})
	if result.Error != nil {
		return "删除失败"
	}
	return "删除成功"
}

func Result_query_all_sql(jsonData map[string]interface{}) []map[string]interface{} {
	db, err := Connect_To_Database_d()

	if err != nil {
		panic(err)
	}
	var results []models.Results

	db.Table("results_view").Where("userid = ? AND documentid = ? ", int(jsonData["userid"].(float64)), int(jsonData["documentid"].(float64))).Find(&results)
	var jsonArray []string
	for _, doc := range results {
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

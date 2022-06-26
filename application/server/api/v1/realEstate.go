package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type RealEstateRequestBody struct {
// 	AccountId   string  `json:"accountId"`   //操作人ID
// 	Proprietor  string  `json:"proprietor"`  //所有者(客户)(客户AccountId)
// 	TotalArea   float64 `json:"totalArea"`   //总面积
// 	LivingSpace float64 `json:"livingSpace"` //生活空间
// }

type RealEstateRequestBody struct {
	AccountId   string `json:"accountId"`  //操作人ID
	Proprietor  string `json:"proprietor"` //所有者(客户)(客户AccountId)
	Carmodel    string `json:"carmodel"`
	Car_scz     string `json:"car_scz"`
	Car_scd     string `json:"car_scd"`
	Car_scsj    string `json:"car_scsj"`
	Car_lbjph   string `json:"car_lbjph"`
	Car_lbjscz  string `json:"car_lbjscz"`
	Car_lbjscd  string `json:"car_lbjscd"`
	Car_lbjscsj string `json:"car_lbjscsj"`
}

type RealEstateQueryRequestBody struct {
	Proprietor string `json:"proprietor"` //所有者(客户)(客户AccountId)
}

func CreateRealEstate(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.AccountId))
	bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	bodyBytes = append(bodyBytes, []byte(body.Carmodel))
	bodyBytes = append(bodyBytes, []byte(body.Car_scz))
	bodyBytes = append(bodyBytes, []byte(body.Car_scd))
	bodyBytes = append(bodyBytes, []byte(body.Car_scsj))
	bodyBytes = append(bodyBytes, []byte(body.Car_lbjph))
	bodyBytes = append(bodyBytes, []byte(body.Car_lbjscz))
	bodyBytes = append(bodyBytes, []byte(body.Car_lbjscd))
	bodyBytes = append(bodyBytes, []byte(body.Car_lbjscsj))

	//调用智能合约
	resp, err := bc.ChannelExecute("createRealEstate", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

func QueryRealEstateList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RealEstateQueryRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.Proprietor != "" {
		bodyBytes = append(bodyBytes, []byte(body.Proprietor))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryRealEstateList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

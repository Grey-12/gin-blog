package api

import (
	"github.com/Grey-12/gin-blog/models"
	"github.com/Grey-12/gin-blog/pkg/errorCode"
	"github.com/Grey-12/gin-blog/pkg/logging"
	"github.com/Grey-12/gin-blog/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := errorCode.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = errorCode.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = errorCode.SUCCESS
			}
		} else {
			code = errorCode.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.SugarLogger.Errorf("Err key: %s, Err Msg: %s", err.Key, err.Message)
			// log.Println(err.Key, err.Message)
		}
	}
	logging.SugarLogger.Debug("认证成功")
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": errorCode.GetMsg(code),
		"data": data,
	})
}
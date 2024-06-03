package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)
type(
	UserControllerType interface{
		Validate(ctx *gin.Context) string
		SaveData(ctx *gin.Context) string
		handler(ctx *gin.Context)
	}
	UserControllerStruct struct{
	}
	UserInfo struct{
		Uid string `json:"uid" binding:"required"`
		Phone_number string `json:"phone_number" binding:"required"`
		Name string `json:"name" binding:"required"`
		Address string `json:"address" binding:"required"`


	}
)
func UserController(ctx *gin.Context)  {
  userInfo :=&UserInfo{}
if err := ctx.ShouldBindJSON(&userInfo)
err!=nil{
	ctx.Error(err)
	ctx.AbortWithStatus(http.StatusBadRequest)
        return
}
fmt.Println("uid : "+userInfo.Uid)
ctx.JSON(http.StatusOK,gin.H{
	"message":"success",
})
}


func (userController *UserControllerStruct) Validate(ctx *gin.Context) string{
return "Hi"

}
func (userController *UserControllerStruct)SaveDate(ctx *gin.Context) string{
return "Saving  Data"

}
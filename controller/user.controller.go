package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type(
	UserControllerType interface{
		Validate(ctx *gin.Context) string
		SaveData(ctx *gin.Context) string
		handler(ctx *gin.Context)
	}
	UserControllerStruct struct{
	}
)
func UserController(ctx *gin.Context)  {
	ctx.String(http.StatusOK,"Test") 
}


func (userController *UserControllerStruct) Validate(ctx *gin.Context) string{
return "Hi"

}
func (userController *UserControllerStruct)SaveDate(ctx *gin.Context) string{
return "Saving  Data"

}
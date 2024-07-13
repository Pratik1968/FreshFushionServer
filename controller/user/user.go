package user

import (
	"net/http"

	"github.com/PratikShekhar/FreshFusionServer/db"
	"github.com/PratikShekhar/FreshFusionServer/model"
	"github.com/gin-gonic/gin"
)
func GET(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
func POST(ctx *gin.Context){
	var userInfo model.UserInfo
	if err:=ctx.ShouldBindJSON(&userInfo); err!=nil{
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err:= db.InsertUser(&userInfo); err!=nil{
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if err:= db.InsertUID(userInfo.Uid); err!=nil{
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(200,gin.H{"status":"success"})
}

package routes

import (
"github.com/gin-gonic/gin"
"github.com/PratikShekhar/FreshFusionServer/controller/user"
)

func userRouteHandler(router *gin.RouterGroup){
	router.POST("/",user.POST)
router.GET("/",user.GET)
}
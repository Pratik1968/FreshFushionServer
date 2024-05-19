package routes

import (

	"github.com/PratikShekhar/FreshFusionServer/controller"
	"github.com/gin-gonic/gin"
)

func userRouteHandler(router *gin.RouterGroup){
	router.GET("/",controller.UserController)

}
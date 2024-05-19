package routes

import (

	"github.com/gin-gonic/gin"
)

func Route(router *gin.Engine){
root := router.Group("/api")
 {
 userRoute := root.Group("/user")
{
userRouteHandler(userRoute)
}
 }
}
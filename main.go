package main

import (

	"github.com/PratikShekhar/FreshFusionServer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    // initiates a gin Engine with the default logger and recovery middleware
    router := gin.Default(

    )
routes.Route(router)
    router.Run(":3000")
}
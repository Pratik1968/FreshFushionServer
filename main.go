package main

import (
	"github.com/PratikShekhar/FreshFusionServer/db"
	"github.com/PratikShekhar/FreshFusionServer/routes"
	"github.com/gin-gonic/gin"
)

func main() {
 db.Init()
    router := gin.Default()

routes.Route(router)
router.Static("/product_images","./product_images")
    router.Run(":3000")
}

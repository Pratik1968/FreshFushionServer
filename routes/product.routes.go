package routes

import (
	"github.com/PratikShekhar/FreshFusionServer/controller/product"
	"github.com/gin-gonic/gin"
)

func productRouteHandler(router *gin.RouterGroup) {
	router.GET("/get_product_by_id", product.GetProductByID)
	router.GET("/get_product_by_trend", product.GetTrendingProducts)
	router.GET("/get_product_by_tag", product.GetProductByText)
router.GET("/get_product_by_category",product.GetByCategory)
}

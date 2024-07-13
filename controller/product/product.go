package product

import (
	"net/http"

	"github.com/PratikShekhar/FreshFusionServer/db"
	"github.com/PratikShekhar/FreshFusionServer/model"
	"github.com/gin-gonic/gin"
)

func GetProductByID(ctx *gin.Context) {
	var id string = ctx.Query("id")
	var product model.Product = db.GetProduct(id)
	ctx.JSON(http.StatusOK, product)

}
func GetProductByText(ctx *gin.Context) {
	var text string = ctx.Query("search")
	var products []model.Product
	products = db.GetSearchProduct(text)

	ctx.JSON(http.StatusOK, products)

}
func GetTrendingProducts(ctx *gin.Context) {
	results := db.GetTrendingProducts()
	ctx.JSON(http.StatusOK, results)
}
func GetByCategory(ctx *gin.Context){
	category := ctx.Query("category")
var results [] model.Product = db.GetByCategory(category)
ctx.JSON(http.StatusOK,results)
}
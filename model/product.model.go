package model

type Product struct{
	Name string `json:"name" binding:"required"`
Price uint16 `json:"price" binding:"required"`
Description string `json:"description" binding:"required"`
Organic bool `json:"organic" binding:"required"`
Expiration uint16 `json:"expiration" binding:"required"`
Rating float64 `json:"rating" binding:"required"`
Calories uint16 `json:"calories" binding:"required"`
Category string `json:"category" binding:"required"`
Image_id uint16 `json:"image_id" binding:"required"`
}
package model

	type UserInfo  struct{
		Uid string `json:"uid" binding:"required"`
		Phone_number string `json:"phoneNumber" binding:"required"`
		Name string `json:"name" binding:"required"`
		Address string `json:"address" binding:"required"`


	}	

package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"project/docs"
	"project/service"
)

func Router() *gin.Engine {
	r := gin.Default()

	//swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//用户
	r.GET("/user/getUserList", service.GetUserList)
	r.GET("/user/createUser", service.CreateUser)
	r.GET("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.GET("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)
	r.GET("/user/getUserByName", service.GetUserByName)
	r.GET("/user/getUserByPhone", service.GetUserByPhone)

	return r
}

package routers

import (
	"github.com/gin-gonic/gin"
	"zg5/Homework01/client/logic"
)

func GinRouter(Group *gin.RouterGroup) {
	user := Group.Group("user")
	{
		//user.Use(middleware.Middle)
		user.POST("login", logic.Login)
	}
}

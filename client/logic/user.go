package logic

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"zg5/Homework01/common"
)

func Login(c *gin.Context) {
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	//
	//user, _ := service.QueryTheUser(username)
	//fmt.Println("user", user)
	//if user.Id == 0 {
	//	c.JSONP(http.StatusAccepted, gin.H{
	//		"code": http.StatusAccepted,
	//		"msg":  "请先注册账号",
	//	})
	//	return
	//}
	//
	////p, _ := utils.DecryptThePassword([]byte(user.Password))
	////pwd := string(p)
	//fmt.Println(user.Password, "][][][][")
	//if password != user.Username {
	//	c.JSONP(http.StatusAccepted, gin.H{
	//		"code": http.StatusAccepted,
	//		"msg":  "密码输入错误",
	//	})
	//	return
	//}
	token, _ := common.SetJwtToken(common.GAOWEIMING, time.Now().Unix(), 3600, "1")
	c.JSONP(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "登录成功",
		"data": token,
	})
	return

}

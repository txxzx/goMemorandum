package handlers

/**
    @date: 2024/7/14
**/

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txxzx/goMemorandum/GateWay/pkg/utils"
	"github.com/txxzx/goMemorandum/GateWay/service"
)

// 用户注册
func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)
	ginCtx.JSON(http.StatusOK, gin.H{"data": userResp})
}

// 用户登录
func UserLogin(ginCtx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从gin.Key中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	ginCtx.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"msg":  "成功",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})
}

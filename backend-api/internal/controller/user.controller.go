package controller

import (
	"GOLANG/github.com/HwuuPhuc0904/backend-api/internal/service"
	"GOLANG/github.com/HwuuPhuc0904/backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// Đây là phương thức viết hàm kiểu method with Receiver(Phương thức với bộ nhận) 
// (uc *UserController) là phần receiver - xác định rằng hàm này là một phương thức của struct UserController
// *UserController biểu thị rằng đây là một con trỏ đến kiểu UserController
// uc là tên biến cho receiver được sử dụng bên trong phương thức

//controller -> service -> repo -> model -> db
func (uc * UserController) GetUserByID(c *gin.Context) {
	response.SuccessResponse(c, 2001, uc.userService.GetUserByID())
}
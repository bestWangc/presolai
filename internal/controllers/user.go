package controllers

import (
	"fmt"
	"net/http"

	"presolai/internal/services"
	"presolai/tools/response"

	"github.com/gin-gonic/gin"
)

var userService = new(services.UserService)

// GetUsers 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	users, err := userService.GetAllUsers()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}

	response.Success(c, users)
}

// GetUsers 获取单个用户
// @Summary 获取单个用户
// @Description 获取单个用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /users/:id [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := userService.GetUserByID(id)
	if err != nil {
		response.Error(c, http.StatusNotFound, "User not found")
		return
	}

	response.Success(c, user)
}

// CreateUser 创建一个新用户
// @Summary 创建一个新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param username body string true "用户名"
// @Param addr body string true "用户地址"
// @Success 200 {object} models.User
// @Router /users [post]
func CreateUser(c *gin.Context) {

	var request struct {
		Username string `json:"username" binding:"required"`
		Addr     string `json:"addr" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Println(err)
		response.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if len(request.Addr) != 44 {
		response.Error(c, http.StatusBadRequest, "Addr must be 44 characters long")
		return
	}
	//判断地址是否存在
	exists, err := userService.AddressExists(request.Addr)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to check address")
		return
	}
	if exists {
		response.Error(c, http.StatusConflict, "Address already exists")
		return
	}

	createdUser, err := userService.CreateUser(request.Username, request.Addr)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	response.Success(c, createdUser)
}

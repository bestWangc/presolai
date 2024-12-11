package services

import (
	"errors"
	"presolai/internal/models"
	"presolai/internal/pkg/mysql"
	"presolai/tools"

	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := mysql.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID 根据 ID 获取用户
func (userService *UserService) GetUserByID(id string) (models.User, error) {
	var user models.User
	if err := mysql.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser 创建一个新用户
func (userService *UserService) CreateUser(username string, addr string) (*models.User, error) {
	if username == "" || addr == "" {
		return nil, errors.New("username and address cannot be empty")
	}

	//生成邀请码
	invateCode, _ := tools.GenerateInviteCode(8)
	newUser := models.User{
		Username:   username,
		Addr:       addr,
		InviteCode: invateCode,
	}

	if err := mysql.DB.Create(&newUser).Error; err != nil {
		return &newUser, err
	}
	return &newUser, nil
}

// UpdateUser 更新用户
func (userService *UserService) UpdateUser(id string, username string) (bool, error) {
    if err := mysql.DB.Model(&models.User{}).Where("id = ?", id).Update("username", username).Error; err != nil {
		return false, err // 返回失败和错误
	}
	return true, nil
}

// DeleteUser 删除用户
func (userService *UserService) DeleteUser(id string) error {
	if err := mysql.DB.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (userService *UserService) AddressExists(addr string) (bool,error) {
    var user models.User
    result := mysql.DB.Where("addr = ?",addr).First(&user)
    if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
        return false, result.Error
    }
    return result.RowsAffected > 0, nil
}

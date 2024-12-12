package services

import (
	"errors"
	"fmt"
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
func (userService *UserService) CreateUser(username string, addr string, inviteUserCode string) (*models.User, error) {
	if username == "" || addr == "" {
		return nil, errors.New("username and address cannot be empty")
	}

	db := mysql.DB
	var inviteUID int32 = 0;
	if inviteUserCode != "" {
		//根据邀请码查询邀请人id
		var user models.User
		inviteUser := db.Where("invite_code = ?", inviteUserCode).First(&user)

		//邀请人未查到 给inviteUID 赋值
		if inviteUser.Error == nil {
			inviteUID = user.ID
		}
	}

	//生成邀请码
	inviteCode := tools.GenerateInviteCode(8)
	newUser := models.User{
		Username:   username,
		Addr:       addr,
		InviteCode: inviteCode,
		InviteUID: inviteUID,
	}
	fmt.Println(newUser);

	if err := db.Create(&newUser).Error; err != nil {
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

func (userService *UserService) AddressExists(addr string) (bool, error) {
	var user models.User
	result := mysql.DB.Where("addr = ?", addr).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (userService *UserService) UsernameExists(username string) (bool, error) {
	var user models.User
	result := mysql.DB.Where("username = ?", username).First(&user)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

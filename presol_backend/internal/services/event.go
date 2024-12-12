package services

import (
	"presolai/internal/models"
	"presolai/internal/pkg/mysql"
	"strconv"
	"time"
)

type EventService struct{}

func (eventService *EventService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := mysql.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// GetUserByID 根据 ID 获取用户
func (eventService *EventService) GetUserByID(id string) (models.User, error) {
	var user models.User
	if err := mysql.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser 创建一个新用户
func (eventService *EventService) CreateEvent(uid string, addr string, title string, desc string, deadline int64, img string) (bool, error) {

	db := mysql.DB

	//生成tid
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestampStr := strconv.FormatInt(timestamp, 10)
	parsedUID, _ := strconv.ParseInt(uid, 10, 32)

	newEvent := models.EventList{
		UID:      int32(parsedUID),
		Tid:      timestampStr,
		Title:    title,
		Desc:     desc,
		Deadline: int32(deadline),
		ImgURL: img,
	}

	if err := db.Create(&newEvent).Error; err != nil {
		return false, err
	}
	return true, nil
}

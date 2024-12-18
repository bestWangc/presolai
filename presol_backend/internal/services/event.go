package services

import (
	"fmt"
	"presolai/internal/models"
	"presolai/internal/pkg/mysql"
	"strconv"
	"time"
)

type EventService struct{}

type EventSummary struct {
	ID       int32  `json:"id"`
	Addr     string  `json:"addr"`
	Tid      string `json:"tid"`
	ImgURL   string `json:"img_url"`
	Title    string `json:"title"`
	Deadline int32  `json:"deadline"`
	Desc     string `json:"desc"`
}

func (eventService *EventService) GetAllEvents() ([]EventSummary, error) {
	var eventList []EventSummary
	if err := mysql.DB.Model(&models.EventList{}).Find(&eventList).Error; err != nil {
		return nil, err
	}
	return eventList, nil
}

// GetUserByID 根据 ID 获取用户
func (eventService *EventService) GetEventByTID(tid string) ([]EventSummary, error) {
	var eventList []EventSummary
	if err := mysql.DB.Model(&models.EventList{}).Where("tid = ?",tid).First(&eventList).Error; err != nil {
		return nil, err
	}
	fmt.Println(eventList);
	return eventList, nil
}

// CreateUser 创建一个新事件
func (eventService *EventService) CreateEvent(addr string, title string, desc string, deadline int64, img string) (bool, error) {

	db := mysql.DB

	//生成tid
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	timestampStr := strconv.FormatInt(timestamp, 10)

	newEvent := models.EventList{
		Tid:      timestampStr,
		Addr:     addr,
		Title:    title,
		Desc:     desc,
		Deadline: int32(deadline),
		ImgURL:   img,
	}
	fmt.Println(newEvent);

	if err := db.Create(&newEvent).Error; err != nil {
		return false, err
	}
	return true, nil
}

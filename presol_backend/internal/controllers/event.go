package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"presolai/internal/services"
	"presolai/tools/response"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	confirm "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/types"
	"github.com/gin-gonic/gin"
)

var eventService = new(services.EventService)

// GetUsers 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Router /users [get]
func GetAllEvents(c *gin.Context) {
	users, err := eventService.GetAllEvents()
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
func GetEvent(c *gin.Context) {
	tid := c.Param("tid")
	fmt.Println("tid", tid)
	event, err := eventService.GetEventByTID(tid)
	if err != nil {
		response.Error(c, http.StatusNotFound, "Event not found")
		return
	}

	response.Success(c, event)
}

// CreateUser 创建一个新事件
// @Tags 事件管理
// @Accept json
// @Produce json
// @Param uid body string true "用户uid"
// @Param addr body string true "用户地址"
// @Param title body string true "事件标题"
// @Param desc body string true "事件描述"
// @Param deadline body string true "事件结束时间"
// @Param img body string true "事件图片"
// @Success 200 {object} models.User
// @Router /comment [post]
func CreateEvent(c *gin.Context) {

	var request struct {
		Addr              string `json:"addr" binding:"required"`
		Title             string `json:"title" binding:"required"`
		Desc              string `json:"desc" binding:"required"`
		Deadline          string `json:"deadline" binding:"required"`
		Img               string `json:"img" binding:"required"`
		SignedTransaction string `json:"signedTransaction"`
	}

	//将图片传到aws上 或者存本地服务器
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid input")
		return
	}

	if len(request.Title) > 100 {
		response.Error(c, http.StatusBadRequest, "Title is too long")
		return
	}

	//将用户的签名交易发送到链上
	if _, err := handleBroadcastTransaction(request.SignedTransaction); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to send transaction")
		return
	}

	//记录数据到数据库
	fmt.Println(request.Deadline)
	timeStr := c.DefaultQuery("time", "")
	fmt.Println(timeStr)
	const timeLayout = "2006-01-02 15:04:05"
	parsedTime, _ := time.Parse(timeLayout, request.Deadline)
	deadTimestamp := parsedTime.Unix()

	fmt.Println(deadTimestamp)

	createdEvent, err := eventService.CreateEvent(request.Addr, request.Title, request.Desc, deadTimestamp, request.Img)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create event")
		return
	}
	response.Success(c, createdEvent)
}

func handleBroadcastTransaction(signedTransaction string) (bool, error) {
	//将前端的字节码还原成交易
	tx := &solana.Transaction{}
	if err := tx.UnmarshalBase64(signedTransaction); err != nil {
		fmt.Println("Failed to unmarshal transaction: %v", err)
		return false, err
	}

	// 创建客户端连接
	rpcClient := rpc.New(rpc.DevNet_RPC)
	fmt.Println(rpc.DevNet_RPC)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	if err != nil {
		return false, err
	}
	fmt.Println("ws 连接成功")
	// 创建一个 Solana 交易对象
	sig, err := confirm.SendAndConfirmTransaction(
		context.TODO(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		return false, err
	}
	spew.Dump(sig)
	return true, nil
}

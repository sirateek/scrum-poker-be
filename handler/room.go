package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirateek/poker-be/internal/room"
	"github.com/sirateek/poker-be/model"
	"github.com/sirateek/poker-be/utils"
	"net/http"
)

type Room struct {
	RoomService    room.Service
	ContextManager utils.ContextManager
}

func NewRoom(router *gin.RouterGroup, RoomService room.Service, ContextManager utils.ContextManager) Room {
	roomHandler := Room{
		RoomService:    RoomService,
		ContextManager: ContextManager,
	}

	router.POST("/", roomHandler.CreateRoom)
	router.POST("/join", roomHandler.JoinRoom)
	router.GET("/:id", roomHandler.GetRoom)

	return roomHandler
}

// CreateRoom is the resolver for the createRoom field.
func (r *Room) CreateRoom(c *gin.Context) {
	var request model.CreateRoom
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	result, err := r.RoomService.CreateRoom(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, result)

}

// JoinRoom is the resolver for the joinRoom field.
func (r *Room) JoinRoom(c *gin.Context) {
	userID := r.ContextManager.GetUserID(c.Request.Context())

	var request model.JoinRoom
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	_, err = r.RoomService.JoinRoom(userID, request.ID, request.Passcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success"})
}

// GetRoom is the resolver for the getRoom field.
func (r *Room) GetRoom(c *gin.Context) {
	id := c.Param("id")
	result, err := r.RoomService.GetRoom(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, result)
}

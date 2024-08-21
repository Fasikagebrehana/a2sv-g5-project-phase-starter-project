package gin

import (
	"context"
	"errors"
	"net/http"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/chat"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ChatHandler struct{
	chatUsecase chat.Usecase
}

func NewChatHandler(chatUsecase chat.Usecase) *ChatHandler{
	return &ChatHandler{
		chatUsecase: chatUsecase,
	}
}

func (chatHandler *ChatHandler) CreateChatHandler(c *gin.Context){
	var form chat.CreateChatForm
	if err := c.ShouldBindJSON(&form); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	
	newChat, err := chatHandler.chatUsecase.CreateChat(context.TODO(), form)
	var validationError validator.ValidationErrors
	if errors.As(err, &validationError){
		errs := infrastructure.ReturnErrorResponse(err)
		c.JSON(http.StatusBadRequest, errs)
	}

	c.JSON(http.StatusCreated, newChat)
	
}
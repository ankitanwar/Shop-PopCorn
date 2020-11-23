package http

import (
	"net/http"

	accesstoken "github.com/ankitanwar/OAuth/domain/accessToken"
	"github.com/ankitanwar/OAuth/services"
	"github.com/gin-gonic/gin"
)

//AccessTokenHandler :to handle the accessToken
type AccessTokenHandler interface {
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	UpdateExperationTime(c *gin.Context)
}

type accessTokenhandler struct {
	service services.Service
}

//NewHandler : To handle the request
func NewHandler(service services.Service) AccessTokenHandler {
	return &accessTokenhandler{
		service: service,
	}

}

func (h *accessTokenhandler) GetByID(c *gin.Context) {
	id := c.Param("access_token_id")
	token, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *accessTokenhandler) Create(c *gin.Context) {
	newRequest := &accesstoken.TokenRequest{}
	if err := c.ShouldBind(newRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := h.service.Create(newRequest)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, token)
}

func (h *accessTokenhandler) UpdateExperationTime(c *gin.Context) {

}

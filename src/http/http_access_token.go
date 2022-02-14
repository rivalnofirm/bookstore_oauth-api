package http

import (
	"github.com/gin-gonic/gin"
	"github.com/rivalnofirm/bookstore_oauth-api/src/src/domain/access_token"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accesToken, err := handler.service.GetById(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accesToken)

	c.JSON(http.StatusNotImplemented, "implement me!!!")
}

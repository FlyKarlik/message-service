package synchandle

import (
	"net/http"

	"github.com/FlyKarlik/message-service/internal/domain"
	"github.com/gin-gonic/gin"
)

type addMessageResponse struct {
	Status string          `json:"status"`
	Code   int             `json:"code"`
	Data   *domain.Message `json:"data"`
}

type getMessageResponse struct {
	Status string          `json:"status"`
	Code   int             `json:"code"`
	Data   *domain.Message `json:"data"`
}

type getAllMessageResponse struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   []domain.Message `json:"data"`
}

type getProcessedMsgResponse struct {
	Status string           `json:"status"`
	Code   int              `json:"code"`
	Data   []domain.Message `json:"data"`
}

type getStatsResponse struct {
	Status string        `json:"status"`
	Code   int           `json:"code"`
	Data   *domain.Stats `json:"data"`
}

type errorResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Error  string `json:"error"`
}

func (h *SyncHandler) errResponse(c *gin.Context, statusCode int, msg string, err error) {
	h.log.LogError(c, statusCode, msg, err)
	c.AbortWithStatusJSON(statusCode, errorResponse{Status: "failure", Code: statusCode, Error: err.Error()})
}

func (h *SyncHandler) successResponse(c *gin.Context, msg string, data interface{}) {
	h.log.LogInfo(c, http.StatusOK, msg)
	c.JSON(http.StatusOK, data)
}

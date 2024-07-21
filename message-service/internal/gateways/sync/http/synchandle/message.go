package synchandle

import (
	"encoding/json"
	"net/http"

	"github.com/FlyKarlik/message-service/internal/errs"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

// @Summary AddMessage
// @Description add new message with content
// @Tags message
// @Accept json
// @Produce json
// @Param input body addMessageRequest true "message info"
// @Success 200 {object} addMessageResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/message/ [post]
func (h *SyncHandler) AddMessage(c *gin.Context) {

	span := h.trace.StartSpan("AddMessage")
	defer span.Finish()

	var input addMessageRequest

	if err := c.BindJSON(&input); err != nil {
		h.errResponse(c, http.StatusBadRequest, "[SyncHandler.AddMessage] c.BindJSON failed", err)
		return
	}

	resp, err := h.usecases.MessageService.AddMessage(input.Content)
	if err != nil {
		h.errResponse(c, http.StatusInternalServerError, "[SyncHandler.AddMessage] h.usecases.MessageService.WriteMessage failed", err)
		return
	}

	dataBytes, err := json.Marshal(resp)
	if err != nil {
		h.errResponse(c, http.StatusInternalServerError, "[SyncHandler.AddMessage] json.Marshal failed", err)
		return
	}

	err = h.w.WriteMessages(c.Request.Context(), kafka.Message{
		Key:   []byte{AddMessageRequest},
		Value: dataBytes,
	})

	if err != nil {
		h.errResponse(c, http.StatusInternalServerError, "[SyncHandler.AddMessage] h.p.Produce failed", err)
		return
	}

	h.successResponse(c, "[SyncHandler.AddMessage] AddMessage successfully done", addMessageResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   resp})
}

// @Summary GetAllMessage
// @Description get all messages (processed and unprocessed)
// @Tags message
// @Accept json
// @Produce json
// @Success 200 {object} getAllMessageResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/message/ [get]
func (h *SyncHandler) GetAllMessage(c *gin.Context) {

	span := h.trace.StartSpan("GetAllMessage")
	defer span.Finish()

	response, err := h.usecases.MessageService.GetAllMessage()
	if err != nil {
		h.errResponse(c, http.StatusInternalServerError, "[SyncHandler.GetAllMessage] h.usecases.MessageService.GetAllMessage failed", err)
		return
	}

	h.successResponse(c, "[SyncHandler.GetAllMessage] GetAllMessage successfully done", getAllMessageResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   response,
	})
}

// @Summary GetAllProcessedMessage
// @Description get all processed messages
// @Tags message
// @Accept json
// @Produce json
// @Success 200 {object} getProcessedMsgResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/message/processed/ [get]
func (h *SyncHandler) GetAllProcessedMessage(c *gin.Context) {

	span := h.trace.StartSpan("GetAllProcessedMessage")
	defer span.Finish()

	response, err := h.usecases.MessageService.GetAllProcessedMessage()
	if err != nil {
		h.errResponse(c, http.StatusInternalServerError, "[SyncHandler.GetAllProcessedMessage] h.usecases.MessageService.GetAllProcessedMessage failed", err)
		return
	}

	h.successResponse(c, "[SyncHandler.GetAllProcessedMessage] GetAllProcessedMessage successfully done", getProcessedMsgResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   response,
	})
}

// @Summary AddMessage
// @Description add new message with content
// @Tags message
// @Accept json
// @Produce json
// @Param id path string true "message id"
// @Success 200 {object} getMessageResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/message/{id} [get]
func (h *SyncHandler) GetMessage(c *gin.Context) {

	span := h.trace.StartSpan("GetMessage")
	defer span.Finish()

	id := c.Param("id")
	if len(id) == 0 {
		h.errResponse(c, http.StatusBadRequest, "[SyncHandler.GetMessage] c.Param failed", errs.ErrInvalidMsgId)
		return
	}

	response, err := h.usecases.MessageService.GetMessage(id)
	if err != nil {
		h.errResponse(c, http.StatusInternalServerError, "[SyncHandler.GetMessage] h.usecases.MessageService.GetMessage failed", err)
		return
	}

	h.successResponse(c, "[SyncHandler.GetMessage] GetMessage successfully done", getMessageResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   response,
	})

}

package synchandle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary GetStats
// @Description get statistics by processed messages
// @Tags stats
// @Accept json
// @Produce json
// @Success 200 {object} getStatsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/message/stats/ [get]
func (h *SyncHandler) GetStats(c *gin.Context) {

	span := h.trace.StartSpan("GetStats")
	defer span.Finish()

	response, err := h.usecases.StatsUsecase.GetStats()
	if err != nil {
		h.errResponse(c, http.StatusInternalServerError, "[SyncHandler.GetStats] h.usecases.StatsUsecase.GetStats failed", err)
		return
	}

	h.successResponse(c, "[SyncHandler.GetStats] GetStats successfully done", getStatsResponse{
		Status: "success",
		Code:   http.StatusOK,
		Data:   response,
	})
}

package advertisement

import (
	utilsHandler "advertisement_api/api/handler/utils"
	"advertisement_api/api/request/advertisement"
	"advertisement_api/api/response/failure"
	advertisementService "advertisement_api/internal/service/advertisement"
	"advertisement_api/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerAddAdvertisement struct {
	Service advertisementService.ServiceAddAdvertisement
}

func (h *HandlerAddAdvertisement) Handle(c *gin.Context) {
	advertisementData, ok := utilsHandler.GetFromContext[advertisement.AddAdvertisementRequest](c, "data")
	if !ok {
		e := failure.ServerError{Reason: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	result, err := h.Service.Add(advertisementData)
	if err == nil {
		c.JSON(http.StatusOK, result)
	} else if errors.As(err, &failure.DayLimitError{}) {
		e := failure.ServerError{Reason: "Day limit error, " + err.Error()}
		c.JSON(http.StatusTooManyRequests, e)
	} else if errors.As(err, &failure.DurationLimitError{}) {
		e := failure.ServerError{Reason: "Duration limit error, " + err.Error()}
		c.JSON(http.StatusConflict, e)
	} else {
		go utils.LogError(err.Error())
		e := failure.ServerError{Reason: "Server error, " + err.Error()}
		c.JSON(http.StatusInternalServerError, e)
	}

}

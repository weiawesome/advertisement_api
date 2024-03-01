/*
The handler with service to handle about add advertisement.
At first, it will get value form context then throw the value to the service.
Finally, according to the value rerun from service then return different response to the client.
*/

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

// HandlerAddAdvertisement is the handler for adding advertisement.
type HandlerAddAdvertisement struct {
	Service advertisementService.ServiceAddAdvertisement // service for adding advertisement
}

// Handle to handle the request about adding advertisement
func (h *HandlerAddAdvertisement) Handle(c *gin.Context) {
	// get the request body from context. if failed, it will return error
	advertisementData, ok := utilsHandler.GetFromContext[advertisement.AddAdvertisementRequest](c, "data")
	if !ok {
		e := failure.ServerError{Reason: "Data not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	// use the data to service and get the return status
	result, err := h.Service.Add(advertisementData)

	if err == nil {
		// success to add the advertisement
		c.JSON(http.StatusOK, result)
	} else if errors.As(err, &failure.DayLimitError{}) {
		// fail to add advertisement due to reaching daily limit
		e := failure.ServerError{Reason: "Day limit error, " + err.Error()}
		c.JSON(http.StatusTooManyRequests, e)
	} else if errors.As(err, &failure.DurationLimitError{}) {
		// fail to add advertisement due to reaching duration limit
		e := failure.ServerError{Reason: "Duration limit error, " + err.Error()}
		c.JSON(http.StatusConflict, e)
	} else {
		// fail to add advertisement due to unknown reason
		// log its error reason and return server error
		go utils.LogError(err.Error())
		e := failure.ServerError{Reason: "Server error, " + err.Error()}
		c.JSON(http.StatusInternalServerError, e)
	}

}

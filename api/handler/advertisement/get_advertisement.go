/*
The handler with service to handle about get advertisements.
At first, it will get value form context then throw the value to the service.
Finally, according to the value rerun from service then return different response to the client.
*/

package advertisement

import (
	utilsHandler "advertisement_api/api/handler/utils"
	"advertisement_api/api/response/failure"
	"advertisement_api/internal/service/advertisement"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandlerGetAdvertisement is the handler for getting advertisement.
type HandlerGetAdvertisement struct {
	Service advertisement.ServiceGetAdvertisement // service for getting advertisement
}

// Handle to handle the request about getting advertisement
func (h *HandlerGetAdvertisement) Handle(c *gin.Context) {
	// get the age from context. if failed, it will return error
	age, ageOk := utilsHandler.GetFromContext[int](c, "age")
	if !ageOk {
		e := failure.ServerError{Reason: "Age not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	// get the country from context. if failed, it will return error
	country, countryOk := utilsHandler.GetFromContext[string](c, "country")
	if !countryOk {
		e := failure.ServerError{Reason: "Country not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	// get the gender from context. if failed, it will return error
	gender, genderOk := utilsHandler.GetFromContext[string](c, "gender")
	if !genderOk {
		e := failure.ServerError{Reason: "Gender not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	// get the platform from context. if failed, it will return error
	platform, platformOk := utilsHandler.GetFromContext[string](c, "platform")
	if !platformOk {
		e := failure.ServerError{Reason: "Platform not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	// get the offset from context. if failed, it will return error
	offset, offsetOk := utilsHandler.GetFromContext[int](c, "offset")
	if !offsetOk {
		e := failure.ServerError{Reason: "Offset not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	// get the limit from context. if failed, it will return error
	limit, limitOk := utilsHandler.GetFromContext[int](c, "limit")
	if !limitOk {
		e := failure.ServerError{Reason: "Limit not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	// use the data to service and get the return status
	result, err := h.Service.Get(utilsHandler.GetFullUrl(c), age, country, gender, platform, offset, limit)

	if err == nil {
		// success to get the advertisements
		c.JSON(http.StatusOK, result)
	} else {
		// fail to add advertisement due to unknown reason
		// log its error reason and return server error
		go utils.LogError(err.Error())
		e := failure.ServerError{Reason: "Server error, " + err.Error()}
		c.JSON(http.StatusInternalServerError, e)
	}
}

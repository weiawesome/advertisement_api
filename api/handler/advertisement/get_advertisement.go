package advertisement

import (
	utilsHandler "advertisement_api/api/handler/utils"
	"advertisement_api/api/response/failure"
	"advertisement_api/internal/service/advertisement"
	"advertisement_api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerGetAdvertisement struct {
	Service advertisement.ServiceGetAdvertisement
}

func (h *HandlerGetAdvertisement) Handle(c *gin.Context) {
	age, ageOk := utilsHandler.GetFromContext[int](c, "age")
	if !ageOk {
		e := failure.ServerError{Reason: "Age not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	country, countryOk := utilsHandler.GetFromContext[string](c, "country")
	if !countryOk {
		e := failure.ServerError{Reason: "Country not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	gender, genderOk := utilsHandler.GetFromContext[string](c, "gender")
	if !genderOk {
		e := failure.ServerError{Reason: "Gender not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	platform, platformOk := utilsHandler.GetFromContext[string](c, "platform")
	if !platformOk {
		e := failure.ServerError{Reason: "Platform not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	offset, offsetOk := utilsHandler.GetFromContext[int](c, "offset")
	if !offsetOk {
		e := failure.ServerError{Reason: "Offset not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}
	limit, limitOk := utilsHandler.GetFromContext[int](c, "limit")
	if !limitOk {
		e := failure.ServerError{Reason: "Limit not found in context"}
		c.JSON(http.StatusInternalServerError, e)
		return
	}

	result, err := h.Service.Get(utilsHandler.GetFullUrl(c), age, country, gender, platform, offset, limit)
	if err == nil {
		c.JSON(http.StatusOK, result)
	} else {
		go utils.LogError(err.Error())
		e := failure.ServerError{Reason: "Server error, " + err.Error()}
		c.JSON(http.StatusInternalServerError, e)
	}
}

package core

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"taxpanel/webserver/dao"
	"taxpanel/webserver/util"
	"taxpanel/webserver/validation/errors"
)

func LogById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	result, err := dao.GetLog(idInt)

	if err != nil {
		errors.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Logs(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")

	result, err := dao.GetLogs(from, to)

	if err != nil {
		errors.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateLog(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	log, err := util.SafeBodyToJson(body)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	expensiveId, err := util.SafeString(getJsonValue(log, []string{"expensive_id"}), "The log must have a expensive_id.")
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	tagId, err := util.SafeString(getJsonValue(log, []string{"tag_id"}), "The log must have a tag_id.")
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	mount, err := util.SafeString(getJsonValue(log, []string{"mount"}), "The log must have a mount.")
	if err != nil {
		errors.RespondError(c, err)
		return
	}
	mountInt, err := strconv.Atoi(mount)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	err = dao.SaveLog(expensiveId, tagId, mountInt)
	if err != nil {
		errors.RespondError(c, errors.InternalServerApiError("The log could not be saved.", err))
		return
	}

	c.JSON(http.StatusOK, "OK")
}

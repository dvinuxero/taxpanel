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

func UserById(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	result, err := dao.GetUser(idInt)

	if err != nil {
		errors.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func Users(c *gin.Context) {
	result, err := dao.GetUsers()

	if err != nil {
		errors.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateUser(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	expensive, err := util.SafeBodyToJson(body)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	name, err := util.SafeString(getJsonValue(expensive, []string{"name"}), "The expensive must have a name.")
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	salary, err := util.SafeString(getJsonValue(expensive, []string{"salary"}), "The expensive must have a salary.")
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	salaryInt, err := strconv.Atoi(salary)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	err = dao.SaveUser(name, salaryInt)
	if err != nil {
		errors.RespondError(c, errors.InternalServerApiError("The expensive could not be saved.", err))
		return
	}

	c.JSON(http.StatusOK, "OK")
}

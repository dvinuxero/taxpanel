package core

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
	"taxpanel/webserver/dao"
	"taxpanel/webserver/util"
	"taxpanel/webserver/validation/errors"
)

func getJsonValue(i interface{}, path []string) interface{} {
	collection, _ := i.(map[string]interface{})

	element, path := path[len(path)-1], path[:len(path)-1]
	for _, value := range path {
		newCollection, ok := collection[value]
		if !ok {
			return nil
		}
		collection, ok = newCollection.(map[string]interface{})
		if !ok {
			return nil
		}
	}
	return collection[element]
}

func ExpensiveById(c *gin.Context) {
	id := c.Param("id")

	result, err := dao.GetExpensive(id)

	if err != nil {
		errors.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func ActiveExpensive(c *gin.Context) {
	id := c.Param("id")

	err := dao.ActiveExpensive(id)
	if err != nil {
		errors.RespondError(c, errors.InternalServerApiError("The expensive could not be saved.", err))
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func InactiveExpensive(c *gin.Context) {
	id := c.Param("id")

	err := dao.InactiveExpensive(id)
	if err != nil {
		errors.RespondError(c, errors.InternalServerApiError("The expensive could not be saved.", err))
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func DeleteExpensive(c *gin.Context) {
	id := c.Param("id")

	err := dao.DeleteExpensive(id)
	if err != nil {
		errors.RespondError(c, errors.InternalServerApiError("The expensive could not be saved.", err))
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func Expensives(c *gin.Context) {
	result, err := dao.GetExpensives()

	if err != nil {
		errors.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func CreateExpensive(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)

	expensive, err := util.SafeBodyToJson(body)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	id, err := util.SafeString(getJsonValue(expensive, []string{"id"}), "The expensive must have a id.")
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	description, err := util.SafeString(getJsonValue(expensive, []string{"description"}), "The expensive must have a description.")
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	err = dao.SaveExpensive(strings.ToUpper(id), description)
	if err != nil {
		errors.RespondError(c, errors.InternalServerApiError("The expensive could not be saved.", err))
		return
	}

	c.JSON(http.StatusOK, "OK")
}

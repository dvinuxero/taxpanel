package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"taxpanel/webserver/dao"
	"taxpanel/webserver/validation/errors"
	"time"
)

type UserSalary struct {
	Date        time.Time                `json:"date"`
	Status      string                   `json:"status"`
	Salary      int                      `json:"salary"`
	SalaryUsed  int                      `json:"salary_used"`
	SalaryExtra int                      `json:"salary_extra"`
	Mandatories []map[string]interface{} `json:"mandatories"`
	Extras      []map[string]interface{} `json:"extras"`
}

const GREEN = "GREEN"
const YELLOW = "YELLOW"
const RED = "RED"
const UNDEFINED = "UNDEFINED"

func Salary(c *gin.Context) {
	//future := c.Query("future")
	usr := c.Query("user")
	idInt, err := strconv.Atoi(usr)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	var salaryUsed int = 0
	var status string = UNDEFINED
	var entity UserSalary

	entity.Date = time.Now() //.AddDate(0, future, 0)

	salary, err := dao.GetUserSalary(idInt)
	if err != nil {
		errors.RespondError(c, err)
		return
	}

	entity.Status = status
	entity.Salary = salary

	mandatoryLogs, err := dao.GetLogsMandatory()
	if err != nil {
		errors.RespondError(c, err)
		return
	}
	entity.Mandatories = mandatoryLogs

	extraLogs, err := dao.GetLogsExtra()
	if err != nil {
		errors.RespondError(c, err)
		return
	}
	entity.Extras = extraLogs

	for _, mand := range mandatoryLogs {
		salaryUsed += mand["mount"].(int)
	}
	for _, mand := range extraLogs {
		salaryUsed += mand["mount"].(int)
	}

	if salaryUsed > salary {
		entity.Status = RED
	} else if (salaryUsed*100)/salary > 90 {
		entity.Status = YELLOW
	} else {
		entity.Status = GREEN
	}

	entity.SalaryUsed = salaryUsed
	entity.SalaryExtra = salary - salaryUsed

	c.JSON(http.StatusOK, entity)
}

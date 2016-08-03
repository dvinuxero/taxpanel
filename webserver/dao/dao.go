package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Expensives
func GetExpensive(key string) (interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT id, description from expensive where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var id string
	var description string

	err = stmtOut.QueryRow(key).Scan(&id, &description)
	if err != nil {
		return nil, err
	}

	var expensive = make(map[string]interface{})
	expensive["id"] = id
	expensive["description"] = description

	return expensive, nil
}

func GetExpensives() (interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, description from expensive where status = 'ACTIVE'")

	var expensives = make([]map[string]interface{}, 0)

	var id, description string

	for rows.Next() {
		err = rows.Scan(&id, &description)

		if err != nil {
			panic(err.Error())
		}

		var expensive = make(map[string]interface{})
		expensive["id"] = id
		expensive["description"] = description

		expensives = append(expensives, expensive)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return expensives, nil
}

func SaveExpensive(id string, description string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return err
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO expensive(id, description) VALUES( ?, ? )")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(id, description)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func ActiveExpensive(id string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return err
	}
	defer db.Close()

	stmtIns, err := db.Prepare("UPDATE expensive set status = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec("ACTIVE", id)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func InactiveExpensive(id string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return err
	}
	defer db.Close()

	stmtIns, err := db.Prepare("UPDATE expensive set status = ? where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec("INACTIVE", id)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

func DeleteExpensive(id string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return err
	}
	defer db.Close()

	stmtIns, err := db.Prepare("DELETE FROM expensive where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(id)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

// Users
func GetUser(key int) (interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT id, name, salary from user where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var id int
	var name string
	var salary string

	err = stmtOut.QueryRow(key).Scan(&id, &name, &salary)
	if err != nil {
		return nil, err
	}

	var user = make(map[string]interface{})
	user["id"] = id
	user["name"] = name
	user["salary"] = salary

	return user, nil
}

func GetUserSalary(key int) (int, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return 0, err
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT salary from user where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var salary int

	err = stmtOut.QueryRow(key).Scan(&salary)
	if err != nil {
		return 0, err
	}

	return salary, nil
}

func GetUsers() (interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, salary from user")

	var users = make([]map[string]interface{}, 0)

	var id, name, salary string

	for rows.Next() {
		err = rows.Scan(&id, &name, &salary)

		if err != nil {
			panic(err.Error())
		}

		var user = make(map[string]interface{})
		user["id"] = id
		user["name"] = name
		user["salary"] = salary

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return users, nil
}

func SaveUser(name string, salary int) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return err
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO user(name, salary) VALUES( ?, ? )")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(name, salary)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

// Logs
func GetLog(key int) (interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT id, expensive_id, tag_id, mount, date from log where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmtOut.Close()

	var id int
	var expensiveId string
	var tagId string
	var mount int
	var date string

	err = stmtOut.QueryRow(key).Scan(&id, &expensiveId, &tagId, &mount, &date)
	if err != nil {
		return nil, err
	}

	var log = make(map[string]interface{})
	log["id"] = id
	log["expensive_id"] = expensiveId
	log["tag_id"] = tagId
	log["mount"] = mount
	log["date"] = date

	return log, nil
}

func GetLogs(from string, to string) (interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return nil, err
	}
	defer db.Close()

	whereDate := ""

	if from != "" && to != "" {
		whereDate = "where date >= STR_TO_DATE('" + from + "','%d-%m-%Y')" + " and date < STR_TO_DATE('" + to + "','%d-%m-%Y')"
	}
	rows, err := db.Query("SELECT id, expensive_id, tag_id, mount, date from log " + whereDate)

	var logs = make([]map[string]interface{}, 0)

	var id, mount int
	var expensiveId, tagId, date string

	for rows.Next() {
		err = rows.Scan(&id, &expensiveId, &tagId, &mount, &date)

		if err != nil {
			panic(err.Error())
		}

		var log = make(map[string]interface{})
		log["id"] = id
		log["expensive_id"] = expensiveId
		log["tag_id"] = tagId
		log["mount"] = mount
		log["date"] = date

		logs = append(logs, log)
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return logs, nil
}

func GetLogsMandatory() ([]map[string]interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")
	var expensiveMandatories = make([]map[string]interface{}, 0)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	expensiveActives, err := db.Query("SELECT id from expensive where status = 'ACTIVE'")
	var expensiveId string
	for expensiveActives.Next() {
		err = expensiveActives.Scan(&expensiveId)
		if err != nil {
			panic(err.Error())
		}
		var query string = "select ifnull(mount, -1) as promedio from log where tag_id = 'mandatory' and expensive_id = '" + expensiveId + "'" + " order by date desc limit 1"
		roundMountRows, err := db.Prepare(query)
		var expensiveMandatory = make(map[string]interface{})
		var promedio int
		if err != nil {
			panic(err.Error())
		}
		defer roundMountRows.Close()

		err = roundMountRows.QueryRow().Scan(&promedio)
		if err != nil {
			promedio = -1
		}
		if promedio >= 0 {
			expensiveMandatory["expensive_id"] = expensiveId
			expensiveMandatory["mount"] = promedio

			expensiveMandatories = append(expensiveMandatories, expensiveMandatory)
		}
	}

	return expensiveMandatories, nil
}

func GetLogsExtra() ([]map[string]interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")
	var expensiveExtras = make([]map[string]interface{}, 0)

	if err != nil {
		return nil, err
	}
	defer db.Close()

	expensiveActives, err := db.Query("SELECT id from expensive where status = 'ACTIVE'")
	var expensiveId string
	for expensiveActives.Next() {
		err = expensiveActives.Scan(&expensiveId)
		if err != nil {
			panic(err.Error())
		}
		var query string = "select ifnull(mount, -1) as promedio from log where tag_id = 'extra' and expensive_id = '" + expensiveId + "'" + " order by date desc limit 1"
		roundMountRows, err := db.Prepare(query)
		var expensiveExtra = make(map[string]interface{})
		var promedio int
		if err != nil {
			panic(err.Error())
		}
		defer roundMountRows.Close()

		err = roundMountRows.QueryRow().Scan(&promedio)
		if err != nil {
			promedio = -1
		}
		if promedio >= 0 {
			expensiveExtra["expensive_id"] = expensiveId
			expensiveExtra["mount"] = promedio

			expensiveExtras = append(expensiveExtras, expensiveExtra)
		}
	}

	return expensiveExtras, nil
}

func SaveLog(expensiveId string, tagId string, mount int) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/taxpanel")

	if err != nil {
		return err
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO log(expensive_id, tag_id, mount) VALUES( ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(expensiveId, tagId, mount)
	if err != nil {
		panic(err.Error())
	}

	return nil
}

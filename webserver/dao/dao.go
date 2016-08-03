package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var expensives = map[string]map[string]interface{}{
	"edesur_lanus": {
		"id":          "edesur_lanus",
		"description": "tax for edesur lanus",
	},
}

// Obtain value from DB and convert this info into a json struct
func GetExpensive(key string) (interface{}, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/tablero")

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
	return expensives, nil
}

func SaveExpensive(id string, description string) error {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/tablero")

	if err != nil {
		return err
	}
	defer db.Close()

	stmtIns, err := db.Prepare("INSERT INTO expensive VALUES( ?, ? )")
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

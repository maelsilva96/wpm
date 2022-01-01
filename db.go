package wpm

import (
	"database/sql"
	"github.com/globalsign/mgo/bson"
	_ "github.com/mattn/go-sqlite3"
)

func LoadModelBind() {
	LoadConfig()
	var err error
	var rows *sql.Rows
	dbSqlLite, err = sql.Open("sqlite3", config.DbConnection)
	if err != nil {
		SendError(err)
	}
	var model Temporary
	rows, err = dbSqlLite.Query("SELECT data FROM temporary WHERE key = ?", config.Guid)
	if err != nil {
		SendError(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&model.Data)
		if err != nil {
			SendError(err)
		}
	}
	err = bson.Unmarshal(model.Data, &modelBind)
	if err != nil {
		SendError(err)
	}
}

func PushResponseData(key string, data interface{}) {
	strData, err := bson.Marshal(data)
	if err != nil {
		SendError(err)
	}
	stmt, errSmtp := dbSqlLite.Prepare("UPDATE temporary SET data = ? WHERE key = ?")
	if errSmtp != nil {
		SendError(errSmtp)
	}
	_, err = stmt.Exec(strData, key)
	if err != nil {
		SendError(errSmtp)
	}
	defer closeSqlLiteExec(stmt)
}

func closeSqlLiteExec(stmt *sql.Stmt) {
	err := stmt.Close()
	if err != nil {
		SendError(err)
	}
}

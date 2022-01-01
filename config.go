package wpm

import (
	"database/sql"
	"fmt"
	"os"
)

type ArgName string

const (
	ArgDbConnection ArgName = "-db"
	ArgTypeRequest  ArgName = "-typeRequest"
	ArgPath         ArgName = "-path"
	ArgGuid         ArgName = "-guid"
)

var config Config
var modelBind ModelBind
var dbSqlLite *sql.DB

func LoadConfig() {
	for i, arg := range os.Args {
		switch arg {
		case string(ArgDbConnection):
			config.DbConnection = os.Args[i+1]
		case string(ArgGuid):
			config.Guid = os.Args[i+1]
		case string(ArgTypeRequest):
			config.TypeRequest = os.Args[i+1]
		case string(ArgPath):
			config.Path = os.Args[i+1]
		}
	}
}

func GetConfig() Config {
	return config
}

func SendData(statusCode int, headers map[string]string, data []byte) {
	model := ModelBind{}
	model.Body = data
	model.Headers = headers
	model.StatusCode = statusCode
	PushResponseData(config.Guid, model)
	os.Exit(0)
}

func SendError(err error) {
	fmt.Print(err)
	os.Exit(1)
}

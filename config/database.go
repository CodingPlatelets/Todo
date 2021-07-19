package config

var dbConfig map[string]interface{}

func init() {
	// init db_server config
	dbConfig = make(map[string]interface{})

	//dbConfig["hostname"] = os.Getenv("database_host")
	//dbConfig["port"] = os.Getenv("database_port")
	//dbConfig["database"] = os.Getenv("database_name")
	//dbConfig["username"] = os.Getenv("database_user")
	//dbConfig["password"] = os.Getenv("database_passwd")
	dbConfig["hostname"] = "127.0.0.1"
	dbConfig["port"] = "3306"
	dbConfig["database"] = "todo"
	dbConfig["username"] = "root"
	dbConfig["password"] = "12345678"
	dbConfig["charset"] = "utf8"
	dbConfig["parseTime"] = "True"

	dbConfig["maxIdleConns"] = 20
	dbConfig["maxOpenConns"] = 100
	/*dbConfig["connMaxIdleTime"] = 10000
	dbConfig["connMaxLifetime"] = 600000*/

}

func GetDbConfig() map[string]interface{} {
	return dbConfig
}

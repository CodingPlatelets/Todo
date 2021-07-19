package config

var mailConfig map[string]interface{}

func init() {
	mailConfig = make(map[string]interface{})

	mailConfig["charset"] = "utf-8"
	mailConfig["smtp_debug"] = 0
	mailConfig["host"] = "smtp.163.com"
	mailConfig["smtp_secure"] = "ssl"
	mailConfig["port"] = 25
	mailConfig["username"] = "ez4zzw@163.com"
	mailConfig["password"] = "TSVITNPANSTOHEMM"
	mailConfig["from"] = "ez4zzw@163.com"
	mailConfig["from_name"] = "夜莺科技"
	mailConfig["address"] = "127.0.0.1:8800"
}

func GetMailConfig() map[string]interface{} {
	return mailConfig
}
